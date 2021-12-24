package main

import (
	"fmt"
	"reprover/mirai-go/adapters"
	"reprover/mirai-go/dos"
)

func main() {

	dealer := Dealer{}

	websocketServer := adapters.NewWebsocketAdapter("crti.cn:18081", "CNMCNMQQ1", 2337935952, dealer)

	fmt.Printf("server %v\n", websocketServer.Sender)

	select {}

}

type Dealer struct {
}

func (d Dealer) MessageDeal(message dos.Message) {
	fmt.Println(message)
}
