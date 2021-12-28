package adapters

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/rxrw/mirai-go/dealers"
	"github.com/rxrw/mirai-go/dos"

	"github.com/goinggo/mapstructure"
	"github.com/gorilla/websocket"
)

var (
	message     chan WebsocketResponse
	syncMessage chan WebsocketResponse
	ws          *websocket.Conn
)

type WebsocketSender struct {
	sessionKey    string
	URL           string
	VerifyKey     string
	QQ            int64
	MessageDealer dealers.MessageDealer
}

func (w WebsocketSender) GetDealer() dealers.MessageDealer {
	return w.MessageDealer
}

func (w WebsocketSender) Connect(senderChan chan Sender) error {
	var err error
	u := url.URL{Scheme: "ws", Host: w.URL, RawQuery: fmt.Sprintf("verifyKey=%s&qq=%d", w.VerifyKey, w.QQ), Path: "/all"}
	ws, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	return nil
}

func (w WebsocketSender) Close() error {
	return ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseAbnormalClosure, ""))
}

//以 HTTP 为例
func (w WebsocketSender) Send(method string, uri string, data interface{}) (interface{}, error) {
	syncID := time.Now().String()

	var err error
	req := &WebsocketRequest{
		SyncID:     syncID,
		Command:    uri,
		SubCommand: strings.ToLower(method),
		Content:    data,
	}

	err = ws.WriteJSON(req)
	if err != nil {
		return nil, fmt.Errorf("send request error: %v", err)
	}

	message := <-syncMessage
	if message.SyncID != syncID {
		syncMessage <- message
		return nil, fmt.Errorf("non same syncId")
	}

	body := message.Data.(map[string]interface{})

	return body, err

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

type WebsocketAdapter struct {
	GeneralAdapter
}

//WaitingMessage 的生产队列
func (w WebsocketAdapter) WaitingMessage() {
	messageBody := WebsocketResponse{}
	for {
		err := ws.ReadJSON(&messageBody)
		if err != nil {
			continue
		}
		message <- messageBody
	}
}

func (w WebsocketAdapter) ConsumeMessage() {
	for message := range message {
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

// UnmarshalMessage unmarshal a websocket message into struct
func (w WebsocketAdapter) UnmarshalMessage(message WebsocketResponse) error {
	syncID := message.SyncID
	if syncID != "-1" {
		// 说明是之前请求留下的，想想怎么办
		syncMessage <- message
		return nil
	}
	// 这里接事件处理器
	if w.Sender.GetDealer() == nil {
		return errors.New("no dealer registered")
	}

	data := message.Data.(map[string]interface{})

	newMessage := dos.Message{}

	mapstructure.Decode(data, &newMessage)
	// 事件推送

	result := w.Sender.GetDealer().MessageDeal(newMessage)

	switch v := result.(type) {
	case string:
		// 直接回复消息
		w.ReplyMessage(newMessage, true, []interface{}{dos.NewPlainMessageChain(v)})
	}
	return nil
}
func NewWebsocketAdapter(URL string, verifyKey string, QQ int64, messageDealer dealers.MessageDealer) *WebsocketAdapter {
	sender := WebsocketSender{
		QQ:            QQ,
		URL:           URL,
		VerifyKey:     verifyKey,
		MessageDealer: messageDealer,
	}

	message = make(chan WebsocketResponse)
	syncMessage = make(chan WebsocketResponse)

	sender.Connect(nil)

	websocketServer := &WebsocketAdapter{
		GeneralAdapter: GeneralAdapter{
			Sender: sender,
		},
	}

	go websocketServer.WaitingMessage()

	go websocketServer.ConsumeMessage()

	return websocketServer
}
