# nya

基于websocket的即时聊天系统

后端使用Golang，前端使用vue+element-ui

### 编译

#### 后端

首先设置`GOPATH`环境变量为系统下项目中`go`目录路径

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
此时开始监听服务器的2333端口，守护进程可以使用Supervisord

### 前端

前端位于`pub`目录，使用vue-cli构建

配置`src/config.vue`中的goweb服务器和websocket服务器url。还有一个默认表情包的配置数组，可以是任何能访问到的图片地址集合

```
npm install
npm run build
```

将dist目录下编译结果置于web服务器访问

web服务器还需要配合vue-route配置重定向(nginx):
```
location / {
     try_files $uri $uri/ @router;
     index index.html;
 }
location @router {
    rewrite ^.*$ /index.html last;
}
```

### 演示地址

> 左上角的link图标复制的URL还将保存自定义的表情包信息哦，CSRF?嗯，这个是什么我不知道

[im.moonprism.cc](http://im.moonprism.cc)

[~](http://im.moonprism.cc/chat/github~/https%3A%2F%2Fassets-cdn.github.com%2Ffavicon.ico/%2Fstatic%2Fimg%2Fi2.png%2C%2Fstatic%2Fimg%2Fi3.png%2Chttp%3A%2F%2F7xk6io.com1.z0.glb.clouddn.com%2Fem%2FXgEasheU.png%2C%2Fstatic%2Fimg%2Fi5.png)