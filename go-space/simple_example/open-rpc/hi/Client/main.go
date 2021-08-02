package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dial err:", err)
	}

	// create receive value
	var reply string

	err = client.Call("HelloServer.Hello", "clt:hello", &reply)
	if err != nil {
		log.Fatal("Call err:", err)
	}

	fmt.Println(reply)
}
