package main

import (
	"fmt"

	"github.com/rxrw/mirai-go/adapters"

	"github.com/rxrw/mirai-go/dos"
)

func main() {

	dealer := Dealer{}

	websocketServer := adapters.NewWebsocketAdapter("127.0.0.1:18081", "VERIFY_KEY", 123456789, dealer)
	//httpServer := adapters.NewHttpAdapter("127.0.0.1:18081", "verifyKey", 1233456, dealer)
	fmt.Println(websocketServer.FriendInfo(7788990))
	//fmt.Println(httpServer.FriendMessage(dos.FriendMessageRequest{
	//	GeneralMessage: dos.GeneralMessage{
	//		Target: 789999888,
	//		MessageChain: []interface{}{
	//			dos.NewPlainMessageChain("测试"),
	//		},
	//	},
	//}))
	// fmt.Println(httpServer.CountMessage())
	// fmt.Printf("http server %v\n", httpServer.Sender)
	select {}

}

// Dealer 需要实现 rxrw/dealers/MessageDealer 接口
type Dealer struct {
}

func (d Dealer) MessageDeal(message dos.Message) interface{} {
	fmt.Println("it's a message:", message)
	return "回复"
}

func (d Dealer) EventDeal(event map[string]interface{}) interface{} {
	return nil
}
