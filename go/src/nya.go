// nya == client
// 每一个ws连接生成一个结构体，其中包含了发送给自己消息的chan，通过goroutine监听并写给用户
// write goroutine监听用户ReadMessage信息，将其写入route.send或广播

package main

import (
	"time"
	ws "github.com/gorilla/websocket"
	"lo"
)

type Nya struct {
	conn *ws.Conn
	recv chan []byte
	send []byte
	uuid int
	name string
	img  string
}

const (
	// 设置过期时间检测为了定时消除占用的goroutine
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = 50 * time.Second
)

func (n *Nya) Writing() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		n.Close()
	}()
	for {
		select {
		case message, ok := <-n.recv:
			n.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				n.conn.WriteMessage(ws.CloseMessage, []byte{})
				return
			}

			w, err := n.conn.NextWriter(ws.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			n.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := n.conn.WriteMessage(ws.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (n *Nya) Reading() {
	defer func() {
		n.Close()
	}()
	for {
		_, message, err := n.conn.ReadMessage()
		if err != nil {
			// 错误处理++
			lo.Instance.Info("Error", err)
			break
		}
		n.send = message
		route.send <- n
	}
}

func (n *Nya) Close() {
	n.conn.Close()
	SignOut(n.uuid)
}
