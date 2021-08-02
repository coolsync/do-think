package main

import (
	"fmt"
	"log"
	"net"
)

// Client struct
type Client struct {
	ServerIP   string
	ServerPort uint
	Name       string
	conn       net.Conn
}

// create Client api
func NewClient(serverIP string, serverPort uint) (*Client, error) {
	// create client obj
	client := &Client{
		ServerIP:   serverIP,
		ServerPort: serverPort,
	}
	// dial server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", client.ServerIP, client.ServerPort))
	if err != nil {
		return nil, fmt.Errorf("net dial err: %v", err)
	}

	client.conn = conn

	// return obj
	return client, nil
}

func main() {
	cli, err := NewClient("localhost", 8081)
	if cli == nil {
		log.Printf(">>>>> connect server faild:%v\n", err)
		return
	}

	log.Println(">>>>> connect server ok...")

	// cli.conn.Write([]byte("conn ok!"))

	select {}
}
