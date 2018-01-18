# Nya

一个非常简单的WebSocket聊天服务器，使用Libevent扩展做事件监听。

考虑到效率的话还可以用pcntl_fork开启多进程Redis队列共享消息。

很久以前就非常想写聊天室来着，前后写了四次，分别用轮询，websocket+socket_select()，swoole，还有这次的libevent扩展实现。关于网络连接，最近接触的一些东西确实让自己学到挺多的感觉。。。

## 使用

依赖：

* PHP5.6
* libevent扩展

运行：

`php ws.php`

## 扩展

客户端发送格式：
```
{
    c:'controller',
    a:[arg1, arg2 ...]
}
```

指向app目录下的Route.php，可通过这个类扩展后台功能

服务器推送格式：
```
{
    e:'event',
    a:[arg1, arg2 ...]
}
```

前台处理请参照public中的demo

> 目前还没有前端的 v0.1

