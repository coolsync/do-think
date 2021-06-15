package main

import (
	"log"
	"net"
	"net/rpc"
)

// HelloServer obj
type HelloServer struct {
}

// remote call method
func (s *HelloServer) Hello(request string, reply *string) error {
	*reply = "srv hello:" + request
	return nil
}

func main() {
	// rpc register name
	rpc.RegisterName("HelloServer", new(HelloServer))

	// listen ip and port
	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("ListenTcp error:", err)
	}
	// accept
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	// tcp binding rpc
	rpc.ServeConn(conn)
}
