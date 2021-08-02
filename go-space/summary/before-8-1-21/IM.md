# IM



# Server

# Start Server



# User Connection Handler  



# Read User Send Msg



#  Encapsulation User  Business



# Query All OnLine User

```go

// user send msg action
func (user *User) UserMsgProcess(msg string) {

	// query all users
	if msg == "who" {
		var allUsers string
		user.server.mapLock.Lock()
		for _, cli := range user.server.OnlineMap {
			allUsers += cli.Name + "\n"
		}
		user.server.mapLock.Unlock()

		user.SendMessage(allUsers)
	} else {
		user.server.BroadCast(user, msg)
	}
}
```



```go
// send msg to correpoding client
func (user *User) SendMessage(msg string) {
	user.Conn.Write([]byte(msg))
}
```



# Rename Present User





# Private Chat



```go
} else if len(msg) >= 4 && msg[:3] == "to|" { // Private Chat
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
		if remoteName == ""{
			user.SendMessage("name must exists ...\n")
			return
		}

		remoteUser, ok := user.server.OnlineMap[remoteName]
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

		remoteUser.SendMessage(fmt.Sprintf("%s 对您说： %s\n", user.Name, content))

	} else {
		user.server.BroadCast(user, msg)
	}
```





# Client



# Client Create

```go
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
```



```go
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
```



# Parse command flag



```go
var srvIP string
var srvPort uint

func init() {
	flag.StringVar(&srvIP, "IP", "localhost", "Set Server IP Addresss (Defualt: localhost)")
	flag.UintVar(&srvPort, "Port", 8081, "Set Server Port (Defualt: 8081)")
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

	select {}
}
```



# Update user name



```go
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
```



```go
// start client
func (client *Client) Run() {
	for client.fg != 0 {
		for client.menu() != true {
		}

		switch client.fg {
		case 1:
			client.publicChat()
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
```





# Public chat mode



```go
// public chat msg handle
func (client *Client) publicChat() {
	// store input msg
	var chatMsg string

	// prompt user input msg
	fmt.Println("Please input msg, exit退出.")

	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		sendMsg := chatMsg + "\n"
		// send msg to server
		_, err := client.conn.Write([]byte(sendMsg))
		if err != nil {
			log.Println("client conn write err:", err)
			break
		}

		chatMsg = ""
		fmt.Println("please input msg, exit退出.")
		fmt.Scanln(&chatMsg)
	}
}
```



```go
// start client
func (client *Client) Run() {
	for client.fg != 0 {
		for client.menu() != true {
		}

		switch client.fg {
		case 1:
			client.publicChat()
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
```





# Private chat mode



```go

```













