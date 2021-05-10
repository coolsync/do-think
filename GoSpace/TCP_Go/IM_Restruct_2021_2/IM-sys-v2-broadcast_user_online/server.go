package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

// Server
type Server struct {
	IP   string
	Port uint

	// online user list
	// OnlineMap map[string]*User
	// mapLock   sync.RWMutex
	OnlineMap sync.Map

	// global msg channel
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

// Listen global msg chan
func (s *Server) ListenMessager() {

	for {
		msg := <-s.MessageChan

		// s.mapLock.Lock()
		// for _, cli := range s.OnlineMap {
		// 	cli.C <- msg
		// }
		// s.mapLock.Unlock()
		s.OnlineMap.Range(func(_, value interface{}) bool {
			// cli := value.(*User)
			value.(*User).C <- msg
			return true
		})
	}

}

// BroadCast user online
func (server *Server) BroadCast(user *User, msg string) {
	sendMsg := fmt.Sprintf("[%s]%s:%s", user.Addr, user.Name, msg)

	server.MessageChan <- sendMsg
}

// Hander connect single client operation
func (s *Server) HandleConn(conn net.Conn) {
	// fmt.Println("connect ok!")
	user := NewUser(conn)

	// s.mapLock.Lock()
	// s.OnlineMap[user.Name] = user
	// s.mapLock.Unlock()
	s.OnlineMap.Store(user.Name, user)
	s.BroadCast(user, "已上线")

	// prevent goroutine exit
	select {}
}

// start Server
func (s *Server) Start() {
	liser, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		log.Println("net.Listen err:", err)
	}

	defer liser.Close()

	// Start listen Message channnel goroutine
	go s.ListenMessager()

	for {
		// accept
		conn, err := liser.Accept()
		if err != nil {
			log.Println("Accept err:", err)
			continue
		}
		// do handler
		// start net Conn gourtine, listen more client connet
		go s.HandleConn(conn)
	}
}
