# nya

基于websocket的即时聊天系统

后端使用Golang，前端使用vue+element-ui

### 编译

#### 后端

首先设置`GOPATH`环境变量为系统中go文件夹路径

引入websocket解析器：
```shell
go get github.com/gorilla/websocket
```
编译运行：
```shell
cd go/bin
go build -o nya ../src/*.go
./nya
```
此时开始监听服务器的2333端口

### 前端

前端位于`pub`目录，使用vue-cli构建

配置`src/config.vue`中的goweb服务器和websocket服务器url。还有一个表情包的配置数组，可以是任何能访问到的图片地址集合

```
npm run build
```

将dist目录下编译结果置于web服务器访问

### 演示地址

[im.moonprism.cc](http://im.moonprism.cc)

[develop - go](http://git.kicoe.com/kicoe/nya)
[develop - pub](http://git.kicoe.com/kicoe/nya-pub)
