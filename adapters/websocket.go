package adapters

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"reprover/mirai-go/dealers"
	"reprover/mirai-go/dos"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type WebsocketSender struct {
	sessionKey    string
	URL           string
	VerifyKey     string
	QQ            int64
	ws            *websocket.Conn
	message       chan WebsocketResponse
	syncMessage   chan WebsocketResponse
	MessageDealer dealers.MessageDealer
}

func (w WebsocketSender) Connect(ws chan Sender) error {
	var err error
	u := url.URL{Scheme: "ws", Host: w.URL, RawQuery: fmt.Sprintf("verifyKey=%s&qq=%d", w.VerifyKey, w.QQ), Path: "/all"}
	w.ws, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	w.message = make(chan WebsocketResponse)
	w.syncMessage = make(chan WebsocketResponse)

	go w.WaitingMessage()

	go w.ConsumeMessage()

	ws <- w

	for range interrupt {
		w.ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseAbnormalClosure, ""))
	}

	return nil
}

//message 的生产队列
func (w WebsocketSender) WaitingMessage() {
	message := WebsocketResponse{}
	for {
		err := w.ws.ReadJSON(&message)
		if err != nil {
			continue
		}
		w.message <- message
	}
}

func (w WebsocketSender) ConsumeMessage() {
	for message := range w.message {
		if w.sessionKey != "" {
			w.UnmarshalMessage(message)
			continue
		}
		k := message.Data.(map[string]interface{})
		code, exists := k["code"]
		if !exists {
			w.UnmarshalMessage(message)
			continue
		}
		if code.(float64) != 0 {
			log.Printf("code is not 0: %d", k["code"])
			continue
		}
		session, ok := k["session"]
		if ok {
			w.sessionKey = session.(string)
		}
	}
}

// UnmarshalMessage unmarshals a websocket message into struct
func (w WebsocketSender) UnmarshalMessage(message WebsocketResponse) error {
	log.Printf("receive message from unmarshal:%v\n", message)
	syncID := message.SyncID
	if syncID != "-1" {
		// 说明是之前请求留下的，想想怎么办
		w.syncMessage <- message
		return nil
	}
	data := message.Data.(map[string]interface{})

	d, _ := json.Marshal(data)
	// 事件推送
	fmt.Printf("its a -1 message non session: %s", string(d))

	// 这里接事件处理器
	if w.MessageDealer == nil {
		return errors.New("no dealer registered")
	}

	newMessage := dos.Message{}

	json.Unmarshal(d, &newMessage)

	w.MessageDealer.MessageDeal(newMessage)
	return nil
}

//以 HTTP 为例
func (w WebsocketSender) Send(method string, uri string, data interface{}, result interface{}) error {
	syncID := time.Now().String()

	var err error
	req := &WebsocketRequest{
		SyncID:     syncID,
		Command:    uri,
		SubCommand: strings.ToLower(method),
		Content:    data,
	}

	err = w.ws.WriteJSON(req)
	if err != nil {
		return fmt.Errorf(" send request error: %v", err)
	}

	message := <-w.syncMessage
	if message.SyncID != syncID {
		w.syncMessage <- message
		return fmt.Errorf("non same syncId")
	}

	jsonData, _ := json.Marshal(message.Data)

	json.Unmarshal(jsonData, &result)

	return err

}

type WebsocketRequest struct {
	SyncID     string      `json:"syncId"`
	Command    string      `json:"command"`
	SubCommand string      `json:"subCommand"`
	Content    interface{} `json:"content"`
}

type WebsocketResponse struct {
	SyncID string      `json:"syncId"`
	Data   interface{} `json:"data"`
}

func NewWebsocketAdapter(URL string, verifyKey string, QQ int64, messageDealer dealers.MessageDealer) *GeneralAdapter {
	ws := make(chan Sender)
	sender := WebsocketSender{
		QQ:            QQ,
		URL:           URL,
		VerifyKey:     verifyKey,
		MessageDealer: messageDealer,
	}
	go sender.Connect(ws)

	res := <-ws

	websocketServer := &GeneralAdapter{
		Sender: res,
	}

	return websocketServer
}
