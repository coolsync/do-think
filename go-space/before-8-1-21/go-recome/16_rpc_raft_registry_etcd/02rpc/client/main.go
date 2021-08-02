package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// Define Params Struct
type ArithRequest struct {
	A, B int
}

// 返回给客户端结构体
type ArithResponse struct {
	Pro int // 乘法
	Quo int // 商
	Rem int // 余数， 取模
}

func main() {
	// 1. dail rpc service
	cli, err := rpc.DialHTTP("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// 2. define Params struct, recv server send msg
	req := ArithRequest{3, 0}

	var res *ArithResponse

	// 3. call remote method
	err = cli.Call("Arith.Multiply", req, &res)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)

	err = cli.Call("Arith.Divide", req, &res)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d / %d = Quo %d, Rem %d\n", req.A, req.B, res.Quo, res.Rem)
}
