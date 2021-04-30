package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

func main() {
	// rpc 连接 服务器
	// conn, err := rpc.Dial("tcp", "localhost:8081")
	conn, err := jsonrpc.Dial("tcp", "localhost:8081")

	if err != nil {
		log.Fatal("rpc.Dial: ", err)
	}
	defer conn.Close()

	// rpc 调用 远程方法
	var reply string // 传出值

	err = conn.Call("hello.HelloWorld", "李白", &reply)
	if err != nil {
		log.Fatal("conn.Call: ", err)
	}

	fmt.Println(reply)
}
