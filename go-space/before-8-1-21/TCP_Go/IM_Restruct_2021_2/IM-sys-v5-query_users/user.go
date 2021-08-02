package main

import (
	"fmt"
	"net"
)

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
	user.server.OnlineMap.Store(user.Name, user)

	// broadcast user on line
	user.server.BroadCast(user, "on line")
}

// user off line action
func (user *User) UserOffLine() {
	// delete user from onlien list
	user.server.OnlineMap.Delete(user.Name)

	// broadcast user off line
	user.server.BroadCast(user, "off line")
}

// send msg to correpoding client
func (user *User) SendMessage(msg string) {
	user.Conn.Write([]byte(msg))
}

// user send msg action
func (user *User) UserMsgProcess(msg string) {

	// query all users
	if msg == "who" {
		var allUsers string
		user.server.OnlineMap.Range(func(key, value interface{}) bool {
			// allUsers += value.(*User).Name + " on line...\n"
			allUsers += fmt.Sprintf("[%s]:%s on line...\n", value.(*User).Addr, value.(*User).Name)
			return true
		})

		user.SendMessage(allUsers)

	} else {
		user.server.BroadCast(user, msg)
	}
}

// listen user channel, has, sent to correponding client
func (user *User) ListenMessage() {
	for {
		msg := <-user.C

		user.Conn.Write([]byte(msg))
	}
}
