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
// 推送单用户
func (u *user) Push(uuid int, msg []byte) {
	u.users[uuid].recv <- msg
}
// 推送所有用户
func (u *user) PushAll(msg []byte) {
	for _, n := range(u.users) {
		n.recv <- msg
	}
}
// 推送处某某之外的所有用户
func (u *user) PushAllNos(msg []byte, uid int) {
	for i, n := range(u.users) {
		if i!=uid {
			n.recv <- msg
		}
	}
}
