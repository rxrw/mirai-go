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
}

func (h HttpSender) Send(method string, uri string, data interface{}) (interface{}, error) {
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
			return nil, err
		}
		defer r.Body.Close()
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
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	resp, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	body := make(map[string]interface{})
	json.Unmarshal(resp, &body)
	if body["code"] != 0 {
		return nil, errors.New(body["msg"].(string))
	}
	val, ok := body["data"]
	if !ok {
		val, ok = body["messageId"]
		if !ok {
			return nil, nil
		}
	}
	return val, nil
}
