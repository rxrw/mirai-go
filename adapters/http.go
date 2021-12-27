package adapters

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rxrw/mirai-go/dealers"
	"github.com/rxrw/mirai-go/dos"

	"github.com/goinggo/mapstructure"
)

type HttpSender struct {
	sessionKey    string
	VerifyKey     string
	QQ            int64
	URL           string
	MessageDealer dealers.MessageDealer
}

func (h HttpSender) GetDealer() dealers.MessageDealer {
	return h.MessageDealer
}

func (h *HttpSender) SetSessionKey(sessionKey string) {
	h.sessionKey = sessionKey
}

func (h HttpSender) Connect(ws chan Sender) error {
	responseBody, err := h.Send("POST", "verify", map[string]string{
		"verifyKey": h.VerifyKey,
	})
	if err != nil {
		return err
	}
	response := responseBody.(map[string]interface{})
	if response["code"].(float64) != 0 {
		return fmt.Errorf("code is not 0: %v", response["code"].(float64))
	}

	h.SetSessionKey(response["session"].(string))

	h.Send("POST", "bind", map[string]interface{}{
		"sessionKey": h.sessionKey,
		"qq":         h.QQ,
	})

	if response["code"].(float64) != 0 {
		return fmt.Errorf("bind code is not 0: %v", response["code"].(float64))
	}

	return nil
}

func (h HttpSender) Close() {
	h.Send("POST", "release", map[string]interface{}{
		"sessionKey": h.sessionKey,
		"qq":         h.QQ,
	})
}

func (h HttpSender) Send(method string, uri string, data interface{}) (interface{}, error) {
	client := http.DefaultClient
	var r *http.Response
	var err error
	uri = fmt.Sprintf("%s/%s", h.URL, uri)
	if method == "GET" {
		values := "sessionKey=" + h.sessionKey
		if data != nil {
			switch data.(type) {
			case map[string]interface{}:
			case map[string]int:
			case map[string]int64:
			case map[string]string:
				for k, v := range data.(map[string]interface{}) {
					values = fmt.Sprintf("%s&%s=%v", values, k, v)
				}
			default:
				jjsonData, _ := json.Marshal(data)
				req := make(map[string]interface{})
				json.Unmarshal(jjsonData, &req)
				for k, v := range req {
					values = fmt.Sprintf("%s&%s=%v", values, k, v)
				}
			}
		}

		uri := fmt.Sprintf("%s?%s", uri, values)

		r, err = client.Get(uri)
		if err != nil {
			return nil, err
		}
	} else {
		bytesData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		r, err = client.Post(uri, "application/json", bytes.NewReader(bytesData))
		if err != nil {
			return nil, err
		}
	}

	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	resp, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	body := make(map[string]interface{})
	json.Unmarshal(resp, &body)

	if body["code"].(float64) != 0 {
		return body, errors.New(body["msg"].(string))
	}
	result, ok := body["data"].(map[string]interface{})
	if !ok {
		return body, err
	}

	return result, err
}

type HttpAdapter struct {
	GeneralAdapter
	// 特有方法
}

// 起定时队列
func (h HttpAdapter) RangeMessage() {
	dealer := h.Sender.GetDealer()
	if dealer != nil {
		go func() {
			for range time.Tick(time.Minute) {
				count, _ := h.CountMessage()
				if count > 0 {
					messages, _ := h.FetchMessage(count + 10)
					for _, message := range messages {
						dealer.MessageDeal(message)
					}
				}
			}
		}()
	}
}

func (h HttpAdapter) CountMessage() (int, error) {
	uri := "countMessage"
	result, err := h.Send("GET", uri, nil)
	if err != nil {
		return 0, err
	}
	return int(result.(map[string]interface{})["data"].(float64)), nil
}

func (h HttpAdapter) FetchMessage(count int) ([]dos.Message, error) {
	uri := "fetchMessage"
	params := make(map[string]int)
	params["count"] = count
	result, err := h.Send("GET", uri, params)
	if err != nil {
		return nil, err
	}
	result = result.(map[string]interface{})["data"]
	var res []dos.Message
	mapstructure.Decode(result, &res)
	return res, err
}

func (h HttpAdapter) FetchLatestMessage(count int) ([]dos.Message, error) {
	uri := "fetchMessage"
	params := make(map[string]int)
	params["count"] = count
	result, err := h.Send("GET", uri, params)
	if err != nil {
		return nil, err
	}
	result = result.(map[string]interface{})["data"]
	var res []dos.Message
	mapstructure.Decode(result, &res)
	return res, err
}

func (h HttpAdapter) PeekMessage(count int) ([]dos.Message, error) {
	uri := "fetchMessage"
	params := make(map[string]int)
	params["count"] = count
	result, err := h.Send("GET", uri, params)
	if err != nil {
		return nil, err
	}
	var res []dos.Message
	result = result.(map[string]interface{})["data"]
	mapstructure.Decode(result, &res)
	return res, err
}

func (h HttpAdapter) PeekLatestMessage(count int) ([]dos.Message, error) {
	uri := "fetchMessage"
	params := make(map[string]int)
	params["count"] = count
	result, err := h.Send("GET", uri, params)
	if err != nil {
		return nil, err
	}
	result = result.(map[string]interface{})["data"]
	var res []dos.Message
	mapstructure.Decode(result, &res)
	return res, err
}

func NewHttpAdapter(URL string, verifyKey string, QQ int64, messageDealer dealers.MessageDealer) *HttpAdapter {
	sender := HttpSender{
		QQ:            QQ,
		URL:           URL,
		VerifyKey:     verifyKey,
		MessageDealer: messageDealer,
	}
	sender.Connect(nil)

	httpAdapter := &HttpAdapter{
		GeneralAdapter{
			Sender: sender,
		},
	}

	if messageDealer != nil {
		httpAdapter.RangeMessage()
	}

	return httpAdapter
}
