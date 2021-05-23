package rpcimpl

import (
	"log"
	"net"
	"reflect"
)

// def client
type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

func (c *Client) callRPC(rpcName string, fPtr interface{}) {
	fn := reflect.ValueOf(fPtr).Elem()

	var f = func(agrs []reflect.Value) []reflect.Value {
		inArgs := make([]interface{}, 0, len(agrs))

		for _, arg := range agrs {
			inArgs = append(inArgs, arg.Interface())
		}

		// encode rpcData
		reqRpc := &RPCData{Name: rpcName, Agrs: inArgs}
		reqData, err := encode(reqRpc)
		if err != nil {
			log.Fatal(err)
		}

		cliSession := NewSession(c.conn)

		err = cliSession.Write(reqData)
		if err != nil {
			log.Fatal(err)
		}

		// dec serve send data
		respData, err := cliSession.Read()
		if err != nil {
			log.Fatal(err)
		}

		respRPC, err := decode(respData)
		if err != nil {
			log.Fatal(err)
		}

		outArgs := make([]reflect.Value, 0, len(respRPC.Agrs))

		for i, arg := range respRPC.Agrs {
			if arg == nil {
				outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
				continue
			}
			outArgs = append(outArgs, reflect.ValueOf(arg))
		}

		return outArgs
	}

	v := reflect.MakeFunc(fn.Type(), f) // fn <---> f
	fn.Set(v)
}

// 	fn := reflect.ValueOf(fPtr).Elem()
// outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
// 	v := reflect.MakeFunc(fn.Type(), f) // fn <---> f
// fn.Set(v)
