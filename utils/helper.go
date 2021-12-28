package utils

import (
	"github.com/rxrw/mirai-go/adapters"
	"github.com/rxrw/mirai-go/dos"
)

// SendSimpleMessage 发送一个简单的文本消息
func SendSimpleMessage(server adapters.GeneralAdapter, messageType string, content string, target ...int64) (int64, error) {
	switch messageType {
	case dos.FRIEND:
		return server.FriendMessage(target[0], 0, []interface{}{dos.NewPlainMessageChain(content)})
	case dos.GROUP:
		return server.GroupMessage(target[0], 0, []interface{}{dos.NewPlainMessageChain(content)})
	case dos.TEMP:
		return server.TempMessage(target[0], target[1], 0, []interface{}{dos.NewPlainMessageChain(content)})
	}
	return 0, nil
}
