// 路由

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// 路由事件
type RE struct {
	login chan *Nya
	signOut chan *Nya
	send chan *Nya
}

var route *RE = &RE{
	login:		make(chan *Nya),
	signOut:	make(chan *Nya),
	send:		make(chan *Nya),
}

// 参数结构
type A struct {
	U int
	A []string
}
// 接收结构
type R struct {
	// 控制器
	C string
	// 参数
	A A
}
// 响应结构
type S struct {
	// 事件
	E string
	// 参数
	A A
}

// 用户唯一标识
var ufd int = 1

func (r *RE) run() {
	for {
		select {
		case nya := <-r.login:
			// 添加到uu
			// 登录事件 uuid 向所有用户发送消息 for recv<-{e:online, a:[uid,name,img]}
			r.Lo(nya)
		case nya := <-r.signOut:
			// 登出事件 uuid 向所有其他用户发送offline事件 for recv<-{e:offline, a:uid}
			r.So(nya.uuid)
		case nya := <-r.send:
			// 消息事件 解析找出发送uuid 写入nya send "msg{uid: , data:}"
			r.Se(nya.uuid, nya.send)
		}
	}
}

// 发送消息
func (r *RE) Se(uuid int, msg []byte) {
	// 解析websocket传来的信息
	var p R
	var s S
	err := json.Unmarshal(msg, &p)
	if err != nil {
		fmt.Println("json err:", err)
	}
	if p.C == "send" {
		// 发送消息
		if len(p.A.A)>0 && uu.Has(p.A.U) {
			// 构造消息
			s.E = "message"
			s.A.U = uuid
			s.A.A = []string{p.A.A[0]}
			b, err := json.Marshal(s)
			if err != nil {
				fmt.Println("json err:", err)
			}
			uu.Push(p.A.U, b)
		}
	}
}
// 登出
func (r *RE) So(uuid int) {
	var s S
	if uu.Rem(uuid) {
		s.E = "offline"
		s.A.U = uuid
		b, err := json.Marshal(s)
		if err != nil {
			fmt.Println("json err:", err)
		}
		uu.PushAll(b)
	}
}
// 登录
func (r *RE) Lo(nya *Nya) {
	uu.Add(nya)
	var s S
	s.E = "online"
	s.A.U = nya.uuid
	s.A.A = []string{nya.name, nya.img}
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	uu.PushAll(b)
	var s2 S
	s2.E = "login"
	s2.A.U = nya.uuid
	s2.A.A = []string{}
	user_list := uu.GetList();
	if user_list != nil {
		for _, u := range(user_list) {
			s2.A.A = append(s2.A.A, strconv.Itoa(u.uuid)+"_"+u.name+"_"+u.img)
		}
	}
	c, err := json.Marshal(s2)
	if err != nil {
		fmt.Println("json err:", err)
	}
	uu.Push(nya.uuid, c)
}