package main

import (
	"fmt"
	"net"
)

// Server create Server struct
type Server struct {
	IP   string
	Port int
}

// NewServer instance Server
func NewServer(ip string, port int) *Server {
	server := &Server{
		IP:   ip,
		Port: port,
	}

	return server
}

// Handler handle client connect
func (s *Server) Handler(conn net.Conn) {
	fmt.Println("connnet ok!")
	conn.Write([]byte("ok!"))
}

// Start start server, finish a few client connect
func (s *Server) Start() {
	// listen port
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}
	// close listen port
	defer listener.Close()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		// do handler
		go s.Handler(conn)
	}
}
