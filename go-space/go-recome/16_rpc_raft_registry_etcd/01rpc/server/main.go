package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Params struct {
	Width  int
	Height int
}

type Rect struct{}

// *Params: inComing param
// res: outGoing param
func (r *Rect) Area(p *Params, res *int) error {
	*res = p.Width * p.Height
	return nil
}

func (r *Rect) Perimeter(p *Params, res *int) error {
	*res = p.Width + p.Height
	return nil
}

func main() {
	// 1. new Rect
	r := new(Rect)

	// 2. service register, bind struct data and method
	err := rpc.Register(r)
	if err != nil {
		log.Fatalf("rpc register err: %v\n", err)
	}

	// 3. service bind http protocal
	rpc.HandleHTTP()

	// 4. listen ip and port
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("http listen err: %v\n", err)
	}
}
