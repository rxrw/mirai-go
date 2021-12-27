---
lastmod: '2021-12-27T08:56:06.723Z'
---
# mirai go http sdk

## 安装

```bash
go get github.com/rxrw/mirai-go
```

## 说明

懒得写readme...

MIRAI-HTTP-API 的 GOLANG-SDK

目前支持:

1. http / websocket 两种通信方式

2. 消息处理器

3. HTTP 定时轮询消息，websocket接收推送消息

4. 完整的结构体定义

## 用法

websocket 服务端会主动推送消息，用 dealer 处理
http 如果注册了dealer，会每分钟获取一次消息，上限1000条。

发送消息的结构都定义好了，推送消息会解析为 rxrw/mirai-go/dos/Message 对象

好累

用例： exmaple.go
