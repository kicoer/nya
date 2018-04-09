// 自定义用户类

package main

type user struct {
	users map[int]*Nya
}

var uu *user = &user{
	users:	make(map[int]*Nya),
}

// 获取所有用户列表
func (u *user) GetList() map[int]*Nya {
	return u.users
}
// 添加单个用户
func (u *user) Add(n *Nya) {
	u.users[n.uuid] = n 
}
// 是否存在某个用户
func (u *user) Has(uuid int) bool {
	_, ok := u.users[uuid]
	if ok {
		return true
	} else {
		return false
	}
}
// 下线用户
func (u *user) Rem(uuid int) bool {
	_, ok := u.users[uuid]
	if ok {
		delete(u.users, uuid)
		return true
	} else {
		return false
	}
}

func (u *user) Push(uuid int, msg []byte) {
	u.users[uuid].recv <- msg
}

func (u *user) PushAll(msg []byte) {
	for _, i := range(u.users) {
		i.recv <- msg
	}
}
