package adapters

import (
	"github.com/rxrw/mirai-go/dealers"
)

type Sender interface {
	Send(method string, uri string, data interface{}) (interface{}, error)
	Connect(ws chan Sender) error
	GetDealer() dealers.MessageDealer
}
