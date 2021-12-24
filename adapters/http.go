package adapters

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type HttpSender struct {
	sessionKey string
	VerifyKey  string
	QQ         int64
	URL        string
}

func (h *HttpSender) SetSessionKey(sessionKey string) {
	h.sessionKey = sessionKey
}

func (h HttpSender) Connect(ws chan Sender) error {
	response := make(map[string]interface{})
	h.Send("POST", "verify", map[string]string{
		"verifyKey": h.VerifyKey,
	}, &response)
	if response["code"].(float64) != 0 {
		return fmt.Errorf("code is not 0: %v", response["code"].(float64))
	}
	h.SetSessionKey(response["sessionKey"].(string))
	return nil
}

func (h HttpSender) Send(method string, uri string, data interface{}, result interface{}) error {
	client := http.DefaultClient
	var r *http.Response
	var err error
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
			return err
		}
		defer r.Body.Close()
	} else {
		bytesData, err := json.Marshal(data)
		if err != nil {
			return err
		}
		r, err = client.Post(uri, "application/json", bytes.NewReader(bytesData))
		if err != nil {
			return err
		}
	}
	if r.StatusCode != http.StatusOK {
		return errors.New(r.Status)
	}
	resp, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(resp, &result)
	body := result.(map[string]interface{})
	if body["code"] != 0 {
		return errors.New(body["msg"].(string))
	}
	result, ok := body["data"]
	if !ok {
		result, ok = body["messageId"]
		if !ok {
			return nil
		}
	}

	return err
}

func NewHttpAdapter(URL string, verifyKey string, QQ int64) *HttpAdapter {
	sender := HttpSender{
		QQ:        QQ,
		URL:       URL,
		VerifyKey: verifyKey,
	}
	sender.Connect(nil)

	httpServer := &HttpAdapter{
		Sender: sender,
	}

	return httpServer
}
