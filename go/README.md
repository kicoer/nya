## Go语言版本

采用[github.com/gorilla/websocket](https://github.com/gorilla/websocket)作为websocket处理器。

### 实现

服务器开启时首先注册全局处理主goroutine `route.run()`，监听所有websocket连接登录，注销和消息事件

为每个websocket连接建立两个goroutine分别监听读和写事件

写事件`nya.Writing`监听的是主goroutine向对应channel中发送的事件，需要该消息写入websocket客户端。读事件`nya.Reading`监听websocket连接对服务器的发送消息事件，并将写入的消息再通过channel触发主goroutine监听的事件

读事件goroutine在浏览器关闭时可以从读取信息中关闭，写事件则需要使用定时器配合关闭。

### 开发

前端模拟接口：
```js
var users = [];
/* event */
var eve = {
    login: function(data){
        // 处理所有用户列表
        // {u: uuid, a: [uuid_name_img, uuid_name_img ...]}
        console.log(data)
    },
    online: function(data){
        // 处理用户上线
        // {u: uuid, a: [name, img]}
        console.log(data)
    },
    offline: function(data){
        // 处理用户下线
        // {u: uuid, a: []}
        console.log(data)
    },
    message: function(data){
        // 处理某用户发送消息
        // {u: uuid, a: [massage]}
        console.log(data)
    },
    broadcast: function(data){
        // 处理某用户广播消息
        // {u: uuid, a: [massage]}
        console.log(data)
    }
}
// 登录 name:用户名 img:头像
var ws = new WebSocket("ws://127.0.0.1:2333/ws?name=kicoe&img=nya")
var all = 0;
ws.onopen = function(){

};
ws.onmessage = function(e){
    console.log(e.data);
    var data = JSON.parse(e.data)
    eve[data.E](data.A)
};
// 向某用户发送消息
function send(id, message){
    ws.send('{"c":"send","a":{"u":'+id+',"a":["'+message+'"]}}')
}
// 推送广播消息
function broadcast(message){
    ws.send('{"c":"broadcast","a":{"u":"","a":["'+message+'"]}}')
}
```

