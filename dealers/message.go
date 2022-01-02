package dealers

import "github.com/rxrw/mirai-go/dos"

type MessageDealer interface {
	MessageDeal(message dos.Message) interface{}
	EventDeal(event map[string]interface{}) interface{}
}
