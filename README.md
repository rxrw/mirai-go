# mirai go http sdk

懒得写readme...

目前支持 http / websocket 两种方式

用法:

websocket 服务端会主动推送消息，用 dealer 处理
http 如果注册了dealer，会每分钟获取一次消息，上限1000条。

发送消息的结构都定义好了，推送消息会解析为 reprover/dos/Message 对象

好累
