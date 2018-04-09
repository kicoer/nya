## Go语言版本

[未完成版本]

采用[github.com/gorilla/websocket](https://github.com/gorilla/websocket)作为websocket处理器

参照其中例子 examples/chat

### 实现

服务器开启时首先注册全局处理主goroutine `route.run`，监听所有客户端读写事件

然后为每个websocket连接建立两个客户端goroutine分别监听读和写事件

读事件`nya.Writing`监听的是服务器主goroutine向对应channel中发送的事件，需要该消息写入websocket客户端。写事件是监听客户端对服务器的写入消息，并将写入的消息再通过channel触发主goroutine的事件

读事件在浏览器关闭时可以从读取信息中关闭，写事件则需要使用定时器配合关闭goroutine

### 使用

```shell
cd $GOPATH/src
go run *.go
```

前端参照`src/demo.html`中的功能

