package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

// create Server struct
type Server struct {
	IP   string
	Port uint

	// online user list
	// OnlineMap map[string]*User
	// mapLock   sync.Mutex
	OnlineMap sync.Map

	// global Message channel
	MessageChan chan string
}

// create Server api
func NewServer(ip string, port uint) (*Server, error) {
	srv := &Server{
		IP:   ip,
		Port: port,
		// OnlineMap:   make(map[string]*User),
		MessageChan: make(chan string),
	}

	return srv, nil
}

// BroadCast user on line method
func (srv *Server) BroadCast(user *User, msg string) {
	sendMsg := fmt.Sprintf("[%s]%s:%s\n", user.Name, user.Addr, msg)

	srv.MessageChan <- sendMsg
}

// HandleConn handle single clinet connection
func (srv *Server) HandleConn(conn net.Conn) {
	// fmt.Println("ok conn")
	user := NewUser(conn)

	// store user to list
	// srv.mapLock.Lock()
	// srv.OnlineMap[user.Name] = user
	// srv.mapLock.Unlock()
	srv.OnlineMap.Store(user.Name, user)

	// broadcast user on line
	srv.BroadCast(user, "aready online")

	// receive user msg
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)

			if n == 0 {
				fmt.Printf("user %s abort connection", user.Addr)
				srv.BroadCast(user, "Off Line")
				return
			}

			if err != nil && err != io.EOF {
				log.Println("receive user msg err:", err)
			}

			msg := buf[:n-1] // trim nc '\n'

			// broadcast user msg
			srv.BroadCast(user, string(msg))

		}

	}()

	// prevent client goroutine auto exit
	select {}
}

// Listen global Message channel, has msg, send msg to user channle
func (srv *Server) ListenMessage() {
	for {
		msg := <-srv.MessageChan

		// srv.mapLock.Lock()
		// for _, cli := range srv.OnlineMap {
		// 	cli.C <- msg
		// }
		// srv.mapLock.Unlock()
		srv.OnlineMap.Range(func(key, value interface{}) bool {
			value.(*User).C <- msg
			return true
		})
	}
}

// start Server
func (srv *Server) Start() {
	// listen tcp
	liser, err := net.Listen("tcp", fmt.Sprintf("%s:%d", srv.IP, srv.Port))
	if err != nil {
		log.Println("net listen err:", err)
	}
	// close listen
	defer liser.Close()

	go srv.ListenMessage()

	for {
		// accept
		conn, err := liser.Accept()
		if err != nil {
			log.Println("accept err:", err)
			continue
		}

		// handler
		go srv.HandleConn(conn)
	}
}
