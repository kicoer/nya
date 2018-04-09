package main

import (
	"net/http"
	"net/url"
	_"net/http/pprof"
	"log"
	ws "github.com/gorilla/websocket"
)

// 初始化uuid参数
var nya_uu int = 0

var upgrader = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func indexH(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "test.html")
}

// 启动ws处理
func wsH(name string, img string, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	nya_uu = nya_uu + 1
	nya := &Nya{
		conn:	conn,
		recv:	make(chan []byte),
		uuid:	nya_uu,
		name:	name,
		img:	img,
	}
	route.login <- nya
	// 监听读
	go nya.Writing()
	// 监听写
	go nya.Reading()
}

func main() {
	// 初始化主服务goroutine
	go route.run()
	// 页面服务
	http.HandleFunc("/", indexH)
	// websocket服务
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// 解析参数
		r.ParseForm()
		queryForm, err := url.ParseQuery(r.URL.RawQuery)
		if err == nil && len(queryForm["name"]) > 0 && len(queryForm["img"]) > 0{
			wsH(queryForm["name"][0], queryForm["img"][0], w, r)
		}
	})
	err := http.ListenAndServe("127.0.0.1:2333", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
