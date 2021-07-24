package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
)

// 乘法 Pro, 商 Quo, 余数 Rem, 运算服务

// Arith 用于注册
type Arith struct{}

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

// 乘法运算
func (a *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

// 除法运算
func (a *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("除数不能为 0")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B

	return nil
}
func main() {
	// 1. 注册服务
	arith := new(Arith)
	// 注册一个Arith服务
	rpc.Register(arith)

	// 2. 服务处理 绑定http协议
	rpc.HandleHTTP()

	// 3. Listen addr
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("listen addr failed, err: %v\n", err)
	}
}
