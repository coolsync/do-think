package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// Client struct
type Client struct {
	ServerIP   string
	ServerPort uint
	Name       string
	conn       net.Conn
	fg         int // current client mode
}

// create Client api
func NewClient(serverIP string, serverPort uint) (*Client, error) {
	// create client obj
	client := &Client{
		ServerIP:   serverIP,
		ServerPort: serverPort,
		fg:         999,
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

var srvIP string
var srvPort uint

func init() {
	flag.StringVar(&srvIP, "ip", "localhost", "Set Server IP Addresss (Defualt: localhost)")
	flag.UintVar(&srvPort, "port", 8081, "Set Server Port (Defualt: 8081)")
}

// public chat msg handle
func (client *Client) PublicChat() {
	// store input msg
	var chatMsg string

	// prompt user input msg
	fmt.Println("please input msg, exit退出.")

	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			// send msg to server
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				log.Println("client conn write err:", err)
				break
			}

			chatMsg = ""
			// prompt user input msg
			fmt.Println("please input msg, exit退出.")
			fmt.Scanln(&chatMsg)
		}
	}
}

// read server send msg, input msg to stdout
func (client *Client) dealResponse() {
	io.Copy(os.Stdout, client.conn) // once has msg, send stdout, 永久阻塞listen
}

// client update user name method
func (client *Client) updateName() bool {
	// msg := fmt.Sprintf("re")
	// input user name
	fmt.Println(">>>>> please input user name: ")
	fmt.Scanln(&client.Name)

	// 封装 msg content
	msg := "rename|" + client.Name + "\n"

	// send msg to server
	_, err := client.conn.Write([]byte(msg))
	if err != nil {
		log.Println("client conn write err:", err)
		return false
	}

	return true
}

// start client
func (client *Client) Run() {
	for client.fg != 0 {
		for client.menu() != true {
		}

		switch client.fg {
		case 1:
			client.PublicChat()
			break
		case 2:
			fmt.Println("private chat mode")
			break
		case 3:
			client.updateName()
			break
		}
	}
}

func (client *Client) menu() bool {
	var fg int // flag

	fmt.Println("1 Public chat mode")
	fmt.Println("2 Private chat mode")
	fmt.Println("3 Update user name")
	fmt.Println("0 退出")

	fmt.Scanln(&fg)
	// fmt.Println(fg)

	if fg >= 0 && fg <= 3 {
		client.fg = fg
		return true
	}
	fmt.Println(">>>>> please choice valid number <<<<<")
	return false

}

func main() {
	// parse command flag
	flag.Parse()

	cli, err := NewClient(srvIP, srvPort)
	if cli == nil {
		log.Printf(">>>>> connect server faild:%v\n", err)
		return
	}

	log.Println(">>>>> connect server ok...")

	// cli.conn.Write([]byte("conn ok!"))
	// start goroutine, listen server send msg
	go cli.dealResponse()

	// select {}
	cli.Run()
}
