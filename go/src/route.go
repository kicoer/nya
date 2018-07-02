// 路由

package main

import (
    "encoding/json"
    "strconv"
    "lo"
)

// 路由事件
type RouteEvent struct {
    login chan *Nya
    signOut chan *Nya
    send chan *Nya
}

var route *RouteEvent = &RouteEvent{
    login:      make(chan *Nya),
    signOut:    make(chan *Nya),
    send:       make(chan *Nya),
}

// 附加信息结构
type Attach struct {
    User int `json:"u"`
    Arg []string `json:"a"`
}
// 接收结构
type Receive struct {
    // 控制器
    Controller string `json:"c"`
    // 参数
    Att Attach `json:"a"`
}
// 响应结构
type Response struct {
    // 事件
    Event string `json:"e"`
    // 参数
    Att Attach `json:"a"`
}

// 用户唯一标识
var ufd int = 1

func (r *RouteEvent) run() {
    for {
        select {
        case nya := <-r.login:
            // 添加到uu
            // 登录事件 uuid 向所有用户发送消息 for recv<-{e:online, a:[uid,name,img]}
            Login(nya)
        case nya := <-r.signOut:
            // 登出事件 uuid 向所有其他用户发送offline事件 for recv<-{e:offline, a:uid}
            SignOut(nya.uuid)
        case nya := <-r.send:
            // 消息事件 解析找出发送uuid 写入nya send "msg{uid: , data:}"
            Send(nya.uuid, nya.send)
        }
    }
}

// 发送消息
func Send(uuid int, msg []byte) {
    // 解析websocket传来的信息
    var rc Receive //p
    var rs Response //s
    err := json.Unmarshal(msg, &rc)
    if err != nil {
        lo.Instance.Info("Error", err)
        return
    }
    // 构造消息
    rs.Att.User = uuid
    rs.Att.Arg = []string{rc.Att.Arg[0]}

    if rc.Controller == "send" {
        // 发送消息
        if len(rc.Att.Arg)>0 && uu.Has(rc.Att.User) {
            rs.Event = "message"
            b, err := json.Marshal(rs)
            if err != nil {
                lo.Instance.Info("Error", err)
                return
            }
            uu.Push(rc.Att.User, b)
        }
    } else if rc.Controller == "broadcast" {
        // 广播消息
        if len(rs.Att.Arg)>0 {
            rs.Event = "broadcast"
            b, err := json.Marshal(rs)
            if err != nil {
                lo.Instance.Info("Error", err)
                return
            }
            uu.PushAllNos(b, uuid)
        }
    }
}
// 登出
func SignOut(uuid int) {
    var rs Response
    if uu.Rem(uuid) {
        rs.Event = "offline"
        rs.Att.User = uuid
        b, err := json.Marshal(rs)
        if err != nil {
            lo.Instance.Info("Error", err)
        }
        uu.PushAll(b)
    }
}
// 登录
func Login(nya *Nya) {
    uu.Add(nya)
    var rs Response
    rs.Event = "online"
    rs.Att.User = nya.uuid
    rs.Att.Arg = []string{nya.name, nya.img}
    b, err := json.Marshal(rs)
    if err != nil {
        lo.Instance.Info("Error", err)
        return
    }
    // 除该用户外全部推送
    uu.PushAllNos(b, nya.uuid)
    var rs2 Response
    rs2.Event = "login"
    rs2.Att.User = nya.uuid
    rs2.Att.Arg = []string{}
    user_list := uu.GetList();
    if user_list != nil {
        for _, u := range(user_list) {
            rs2.Att.Arg = append(rs2.Att.Arg, strconv.Itoa(u.uuid)+"_"+u.name+"_"+u.img)
        }
    }
    c, err := json.Marshal(rs2)
    if err != nil {
        lo.Instance.Info("Error", err)
        return
    }
    uu.Push(nya.uuid, c)
}