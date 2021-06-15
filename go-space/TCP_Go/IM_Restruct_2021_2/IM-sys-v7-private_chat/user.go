package main

import (
	"fmt"
	"net"
	"strings"
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
	// user.server.mapLock.Lock()
	// user.server.OnlineMap[user.Name] = user
	// user.server.mapLock.Unlock()
	user.server.OnlineMap.Store(user.Name, user)
	// broadcast user on line
	user.server.BroadCast(user, "on line")
}

// user off line action
func (user *User) UserOffLine() {
	// delete user from onlien list
	// user.server.mapLock.Lock()
	// delete(user.server.OnlineMap, user.Name)
	// user.server.mapLock.Unlock()
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

		// user.server.mapLock.Lock()
		// for _, cli := range user.server.OnlineMap {
		// 	allUsers += fmt.Sprintf("[%s]%s\n", cli.Addr, cli.Name)
		// }
		// user.server.mapLock.Unlock()
		user.server.OnlineMap.Range(func(key, value interface{}) bool {
			allUsers += fmt.Sprintf("[%s]%s\n", value.(*User).Addr, value.(*User).Name)
			return true
		})

		user.SendMessage(allUsers)

	} else if len(msg) >= 8 && msg[:7] == "rename|" {
		// rename current user
		newName := strings.Split(msg, "|")[1]

		// thinking user.server.OnlineMap.LoadOrStore(newName, user) ?

		// if _, ok := user.server.OnlineMap[newName]; ok {
		if _, ok := user.server.OnlineMap.Load(newName); ok {
			user.SendMessage("Name 已经存在！\n")

		} else {

			// modify name operation
			// user.server.mapLock.Lock()
			// delete(user.server.OnlineMap, user.Name)
			// user.server.OnlineMap[newName] = user
			// user.server.mapLock.Unlock()
			user.server.OnlineMap.Delete(user.Name)
			user.server.OnlineMap.Store(newName, user)

			user.Name = newName
			user.SendMessage("你的name已经更新:" + user.Name + "\n")
		}

	} else if len(msg) >= 3 && msg[:3] == "to|" { // Private Chat
		// msg format: to|name|msg content
		msgSli := strings.Split(msg, "|")
		if len(msgSli) != 3 {
			user.SendMessage("msg format: to|name|msg content...\n")
			return
		}
		// get other user name
		// remoteName := strings.Split(msg, "|")[1]
		remoteName := msgSli[1]
		// name judge
		if remoteName == "" {
			user.SendMessage("name must exists ...\n")
			return
		}

		// remoteUser, ok := user.server.OnlineMap[remoteName]
		remoteUser, ok := user.server.OnlineMap.Load(remoteName)

		if !ok {
			user.SendMessage("remote name not exists...\n")
			return
		}

		// get and judge content
		// content := strings.Split(msg, "|")[2]
		content := msgSli[2]
		if content == "" {
			user.SendMessage("conntent is nil...\n")
			return
		}

		remoteUser.(*User).SendMessage(fmt.Sprintf("%s 对您说： %s\n", user.Name, content))

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
