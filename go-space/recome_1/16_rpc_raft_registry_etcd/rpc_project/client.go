package rpcimpl

import (
	"log"
	"net"
	"reflect"
)

// def cli
type Client struct {
	conn net.Conn
}

// client api
func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

func (c *Client) callRPC(rpcName string, fPtr interface{}) {
	fn := reflect.ValueOf(fPtr).Elem() // pass fn varaible

	f := func(args []reflect.Value) []reflect.Value {
		// handle inner args
		inArgs := make([]interface{}, 0, len(args))
		for _, arg := range args {
			inArgs = append(inArgs, arg.Interface())
		}
		// create cli session
		// conn, err := net.Dial("tcp", ":8081")
		// if err != nil {
		// 	log.Fatalf("net dial err:%v\n", err)
		// }
		cliSession := NewSession(c.conn)

		// encode data
		reqRpc := RPCData{rpcName, inArgs}
		b, err := encode(&reqRpc)
		if err != nil {
			log.Fatalf("encode(&reqRpc) err:%v\n", err)
		}
		// send to server
		err = cliSession.Write(b)
		if err != nil {
			log.Fatalf("cliSession.Write(b) err:%v\n", err)
		}

		// handle server return data
		respData, err := cliSession.Read()
		if err != nil {
			log.Fatalf("cliSession.Read() err:%v\n", err)
		}

		rpcData, err := decode(respData)
		if err != nil {
			log.Fatalf("cli decode(respData) err:%v\n", err)
		}
		// parse args
		outAgrs := make([]reflect.Value, 0, len(rpcData.Args))

		// return out args
		for i, arg := range rpcData.Args {
			if arg == nil {
				outAgrs = append(outAgrs, reflect.Zero(fn.Type().Out(i)))
				continue
			}
			outAgrs = append(outAgrs, reflect.ValueOf(arg))
		}
		return outAgrs
	}

	// MakeFunc returns a new function of the given Type that wraps the function fn.
	v := reflect.MakeFunc(fn.Type(), f)
	fn.Set(v)
}
