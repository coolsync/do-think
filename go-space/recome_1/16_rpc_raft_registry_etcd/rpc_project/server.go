package rpcimpl

import (
	"errors"
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

// manager a map list, store func name <--> func reflect val
func (s *Server) Register(rpc_name string, f interface{}) error {
	if _, ok := s.funcs[rpc_name]; ok {
		return errors.New("no register")
	}
	fnVal := reflect.ValueOf(f)
	s.funcs[rpc_name] = fnVal
	return nil
}

// server run method
func (s *Server) Run() {
	// listen ip
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("server net.Listen err:%v", err)
	}
	defer lis.Close()
	// wait cli connect
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("server Accept() err:%v", err)
		}
		// defer conn.Close()
		serSession := NewSession(conn)

		// read cli send msg
		b, err := serSession.Read()
		if err != nil {
			log.Fatalf("server serSession.Read() err:%v", err)
		}

		// decode bs
		rpc_data, err := decode(b)
		if err != nil {
			log.Fatalf("server decode err:%v", err)
		}

		// query map
		f, ok := s.funcs[rpc_data.Name]
		if !ok {
			log.Fatalf("server query map err:%v", err)
		}

		// traveser rpc args
		inArgs := make([]reflect.Value, 0, len(rpc_data.Args))

		for _, arg := range rpc_data.Args {
			inArgs = append(inArgs, reflect.ValueOf(arg))
		}

		// call args, out: return cli msg
		out := f.Call(inArgs)

		outArgs := make([]interface{}, 0, len(out))

		for _, o := range out {
			outArgs = append(outArgs, o.Interface())
		}

		// Encapsulating data send to cli
		respRPC := RPCData{rpc_data.Name, outArgs}

		respData, err := encode(&respRPC)
		if err != nil {
			log.Fatalf("server encode(&respRPC) err:%v", err)
		}

		err = serSession.Write(respData)
		if err != nil {
			log.Fatalf("server serSession.Write err:%v", err)
		}
	}
}
