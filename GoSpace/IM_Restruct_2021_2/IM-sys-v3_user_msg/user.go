package main

import "net"

// create User
type User struct {
	Name string
	Addr string
	C    chan string
	Conn net.Conn
}

// create User api
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		Conn: conn,
	}

	// start more go
	go user.ListenMessage()

	return user
}

// listen user self channel
func (user *User) ListenMessage() {
	for {
		msg := <-user.C

		// write to again client
		user.Conn.Write([]byte(msg))
	}
}
