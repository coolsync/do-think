package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

// create Server struct
type Server struct {
	IP   string
	Port uint

	// user on line list
	// OnlineMap map[string]*User
	// mapLock   sync.Mutex
	OnlineMap sync.Map

	// global channel
	MessageChan chan string
}

// create Server api
func NewServer(ip string, port uint) *Server {
	server := &Server{
		IP:   ip,
		Port: port,
		// OnlineMap:   make(map[string]*User),
		MessageChan: make(chan string),
	}

	return server
}

// Listen global channel msg
func (server *Server) ListenMessager() {
	for {
		msg := <-server.MessageChan

		// send msg to more client
		// server.mapLock.Lock()
		// for _, user := range server.OnlineMap {
		// 	user.C <- msg
		// }
		// server.mapLock.Unlock()
		server.OnlineMap.Range(func(key, value interface{}) bool {
			value.(*User).C <- msg
			return true
		})
	}
}

// BroadCast user msg
func (srv *Server) BroadCast(user *User, msg string) {
	sendMsg := fmt.Sprintf("[%s]%s:%s\n", user.Addr, user.Name, msg)

	// send to global message channel
	srv.MessageChan <- sendMsg
}

// handle single client bz
func (srv *Server) HandleConn(conn net.Conn) {

	user := NewUser(conn, srv)

	user.UserOnLine()

	// Judge user if not live
	isLive := make(chan bool)

	// handle user send msg
	go func() {
		buf := make([]byte, 4096)

		for {
			n, err := conn.Read(buf)
			if n == 0 {
				log.Println(user.Name + " off line")
				// srv.BroadCast(user, "Off Line")
				user.UserOffLine()
				return
			}

			if err != nil && err != io.EOF {
				log.Println("conn read user msg err:", err)
				return
			}

			// broadcast user msg to other
			msg := string(buf[:n-1]) // trim nc '\n'

			user.UserMsgProcess(msg)

			// this user has msg, send bool val to isLive
			isLive <- true
		}
	}()

	// prevent conn auto exit
	// listen chnneal data stream
	for {
		select {
		// judge user if not live
		case <-isLive:
			// 不作为， 证明user还live

		case <-time.After(time.Second * 300):
			// delete user
			user.SendMessage("你被剔出！\n")

			// user.UserOffLine()

			// destory this user used resources
			close(user.C)

			// close user conn
			conn.Close()

			return // exit
		}
	}
}

// Start Server
func (s *Server) Start() {
	// listen ip and port
	liser, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		log.Println("net listen err:", err)
		return
	}

	// close listen
	defer liser.Close()

	// start global channel goroutine
	go s.ListenMessager()

	for {
		// accept
		conn, err := liser.Accept()
		if err != nil {
			log.Println("accept err:", err)
			return
		}

		// do handler conn
		go s.HandleConn(conn)
	}
}
