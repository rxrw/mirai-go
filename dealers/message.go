package dealers

import "reprover/mirai-go/dos"

type MessageDealer interface {
	MessageDeal(message dos.Message)
}
