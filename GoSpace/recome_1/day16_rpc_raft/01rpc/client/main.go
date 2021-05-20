package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Width  int
	Height int
}

func main() {
	// connect rpc server
	cli, err := rpc.DialHTTP("tcp", ":8081")
	if err != nil {
		log.Fatalf("rpc dial server err: %v\n", err)
	}

	var reply int // recv server send date

	// call remote method

	// &Params{10, 20}: inComing param
	// &reply: outGoing param
	err = cli.Call("Rect.Area", &Params{10, 20}, &reply)
	if err != nil {
		log.Fatalf("rpc client call remote Area method err: %v\n", err)
	}

	fmt.Println("area:", reply)

	err = cli.Call("Rect.Perimeter", &Params{10, 20}, &reply)
	if err != nil {
		log.Fatalf("rpc client call remote Perimeter method err: %v\n", err)
	}

	fmt.Println("perimeter:", reply)
}
