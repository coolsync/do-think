package rpcimpl

import (
	"fmt"
	"log"
	"net"
	"reflect"
)

// def server
type Server struct {
	addr  string
	funcs map[string]reflect.Value
}

// server api
func NewServer(addr string) *Server {
	return &Server{addr: addr, funcs: make(map[string]reflect.Value)}
}

// register func
func (s *Server) RegisterName(rpcName string, f interface{}) {

	if _, ok := s.funcs[rpcName]; ok {
		return
	}

	fVal := reflect.ValueOf(f)
	s.funcs[rpcName] = fVal
}

// srv run
func (s *Server) Run() {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
		}
		serSession := NewSession(conn)

		// decode rpcData
		reqData, err := serSession.Read()
		if err != nil {
			log.Fatal(err)
		}

		reqRpc, err := decode(reqData)
		if err != nil {
			log.Fatal(err)
		}

		// judge is not exists in map
		f, ok := s.funcs[reqRpc.Name]
		if !ok {
			fmt.Println("req rpcData is not exist ...")
			return
		}

		// get Args
		outArgs := make([]reflect.Value, 0, len(reqRpc.Agrs))

		for _, arg := range reqRpc.Agrs {
			outArgs = append(outArgs, reflect.ValueOf(arg))
		}

		inArgs := f.Call(outArgs)

		out := make([]interface{}, 0, len(inArgs))

		for _, arg := range inArgs {
			out = append(out, arg.Interface())
		}

		// encode data send to cli
		respRpc := &RPCData{Name: reqRpc.Name, Agrs: out}
		b, err := encode(respRpc)
		if err != nil {
			log.Fatal(err)
		}

		err = serSession.Write(b)
		if err != nil {
			log.Fatal(err)
		}
	}
}
