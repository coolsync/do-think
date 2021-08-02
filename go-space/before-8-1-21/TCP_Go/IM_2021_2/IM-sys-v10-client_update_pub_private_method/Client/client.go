package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
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

// query all users
func (client *Client) QueryAll() {
	msg := "who\n"

	_, err := client.conn.Write([]byte(msg))
	if err != nil {
		log.Println("QueryAll conn write err:", err)
		return
	}
}

// private chat msg handle
func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	client.QueryAll()
	fmt.Println("Please input private user, exit chat end:")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		if len(remoteName) != 0 {

			fmt.Println("please input private msg, exit chat end.")
			fmt.Scanln(&chatMsg)

			for chatMsg != "exit" {
				if len(chatMsg) != 0 {

					sendMsg := "to|" + remoteName + "|" + chatMsg + "\n"

					_, err := client.conn.Write([]byte(sendMsg))
					if err != nil {
						log.Println("PrivateChat conn write err:", err)
						break
					}
				}
				chatMsg := ""
				fmt.Println("please input private msg, exit chat end.")
				fmt.Scanln(&chatMsg)
			}

		}
		remoteName = ""
		fmt.Println("Please input private user, exit chat end:")
		fmt.Scanln(&remoteName)
	}
}

// public chat msg handle
func (client *Client) PublicChat() {
	// store input msg
	var chatMsg string

	// prompt user input msg
	fmt.Println("Please input msg, exit退出.")

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
		}
		chatMsg = ""
		fmt.Println("please input msg, exit退出.")
		fmt.Scanln(&chatMsg)
	}
}

// read server send msg, input msg to stdout
func (client *Client) dealResponse() {
	
	io.Copy(os.Stdout, client.conn) // once has msg, send stdout, 永久阻塞listen
	
	// buf := make([]byte, 4096)
	// for {
	// 	n, err := client.conn.Read(buf)
	// 	if n == 0 {
	// 		fmt.Println("server aready close...")
	// 		return
	// 	}

	// 	if err != nil {
	// 		log.Println("dealResponse err:", err)
	// 		return
	// 	}

	// 	fmt.Fprintf(os.Stdout, "%s\n", string(buf[:n]))	
	// }
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
			client.PrivateChat()
			break
		case 3:
			client.updateName()
			break
		}
	}
}

// menu method
func (client *Client) menu() bool {
	var fgStr string // flag

	fmt.Println("1 Public chat mode")
	fmt.Println("2 Private chat mode")
	fmt.Println("3 Update user name")
	fmt.Println("0 退出")

	fmt.Scanln(&fgStr)

	fg, err := strconv.Atoi(fgStr)
	if err != nil {
		fmt.Println(">>>>> please choice valid number <<<<<")
		return false
	}

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
