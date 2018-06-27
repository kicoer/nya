package main

import (
    "net/http"
    "net/url"
    _"net/http/pprof"
    ws "github.com/gorilla/websocket"
    "crypto/md5"
    "encoding/hex"
    "fmt"
    //"strconv"
    "lo"
)

// 初始化uuid参数
var nya_uu int = 0

// 初始化websocket解析
var upgrader = ws.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func indexH(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "test.html")
}

func wsH(name string, img string, w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        lo.Instance.Info("Error", err)
        return
    }
    nya_uu = nya_uu + 1
    nya := &Nya{
        conn:   conn,
        recv:   make(chan []byte),
        uuid:   nya_uu,
        name:   name,
        img:    img,
    }
    // 暂时不记录登录日志
    // lo.Instance.Info("Info", "["+name+"]("+img+") "+strconv.Itoa(nya_uu)+" login")
    route.login <- nya
    // 监听读
    go nya.Writing()
    // 监听写
    go nya.Reading()
}

func main() {
    // 定义可访问权限
    upgrader.CheckOrigin = func(r *http.Request) bool {
        return true
    }
    // 初始化主服务goroutine
    go route.run()
    // 主页面
    http.HandleFunc("/", indexH)
    // websocket服务
    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        // 解析参数
        r.ParseForm()
        queryForm, err := url.ParseQuery(r.URL.RawQuery)
        if err == nil && len(queryForm["name"]) > 0 && len(queryForm["img"]) > 0{
            var imgs string = queryForm["img"][0]
            if len(queryForm["nk"]) > 0 && len(queryForm["s"]) > 0 {
                imgs = imgs+"&nk="+queryForm["nk"][0]+"&s="+queryForm["s"][0]
            }
            wsH(queryForm["name"][0], imgs, w, r)
        }
    })
    // 头像解析服务
    http.HandleFunc("/nya", func(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
        queryForm, err := url.ParseQuery(r.URL.RawQuery)
        var iURL string = "nya.jpg"
        // m:[g,q] u[email, qq]
        if err == nil && len(queryForm["m"]) > 0 && len(queryForm["u"]) > 0{
            if(queryForm["m"][0] == "g") {
                // 获取gravatar头像链接
                hasher := md5.New()
                hasher.Write([]byte(queryForm["u"][0]))
                iURL = "https://secure.gravatar.com/avatar/"+hex.EncodeToString(hasher.Sum(nil))
            } else if (queryForm["m"][0] == "q") {
                // 获取qq头像链接
                iURL = "https://q.qlogo.cn/g?b=qq&nk="+queryForm["u"][0]+"&s=100"
            }
        }
        w.Header().Set("Access-Control-Allow-Origin", "*")
        fmt.Fprint(w, iURL)
    })
    err := http.ListenAndServe("0.0.0.0:2333", nil)
    if err != nil {
        lo.Instance.Info("Error", err)
    }
}
