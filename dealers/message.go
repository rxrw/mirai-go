package dealers

import "rxrw/mirai-go/dos"

type MessageDealer interface {
	MessageDeal(message dos.Message)
}
