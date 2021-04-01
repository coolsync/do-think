package main

import "net"

// create User struct
type User struct {
	Name string
	Addr string
	C    chan string
	Conn net.Conn

	server *Server
}

// create User api
func NewUser(conn net.Conn, srv *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		Conn:   conn,
		server: srv,
	}

	// use channle connection
	go user.ListenMessage()

	return user
}

// user on line action
func (user *User) UserOnLine() {
	// store user to online list
	user.server.mapLock.Lock()
	user.server.OnlineMap[user.Name] = user
	user.server.mapLock.Unlock()

	// broadcast user on line
	user.server.BroadCast(user, "on line")
}

// user off line action
func (user *User) UserOffLine() {
	// delete user from onlien list
	user.server.mapLock.Lock()
	delete(user.server.OnlineMap, user.Name)
	user.server.mapLock.Unlock()

	// broadcast user off line
	user.server.BroadCast(user, "off line")
}

// user send msg action
func (user *User) UserMsgProcess(msg string) {
	user.server.BroadCast(user, msg)
}

// listen user channel, has, sent to correponding client
func (user *User) ListenMessage() {
	for {
		msg := <-user.C

		user.Conn.Write([]byte(msg))
	}
}
