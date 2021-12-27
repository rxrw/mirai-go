package main

import (
	"fmt"
	"reprover/mirai-go/adapters"
	"reprover/mirai-go/dos"
)

func main() {

	dealer := Dealer{}

	websocketServer := adapters.NewWebsocketAdapter("crti.cn:18081", "CNMCNMQQ1", 2337935952, dealer)
	// httpServer := adapters.NewHttpAdapter("http://crti.cn:18080", "CNMCNMQQ1", 2337935952, dealer)
	fmt.Println(websocketServer.FriendInfo(460514723))
	// fmt.Println(httpServer.FriendMessage(dos.FriendMessageRequest{
	// 	GeneralMessage: dos.GeneralMessage{
	// 		Target: 460514723,
	// 		MessageChain: []interface{}{
	// 			dos.NewPlainMessageChain("ces"),
	// 		},
	// 	},
	// }))
	// fmt.Println(httpServer.CountMessage())
	// fmt.Printf("http server %v\n", httpServer.Sender)
	select {}

}

type Dealer struct {
}

func (d Dealer) MessageDeal(message dos.Message) {
	fmt.Println("it's a message:", message)
}
