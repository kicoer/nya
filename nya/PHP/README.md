## PHP版本

采用用Libevent扩展做事件监听。

考虑到效率的话还可以用pcntl_fork开启多进程，Redis队列共享消息。

用PHP写过很多很多次聊天室了...

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

