package rpcimpl

import (
	"encoding/gob"
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func TestCodec(t *testing.T) {
	gob.Register(Person{}) // not impl interface, codec is not err

	p := &Person{Name: "mark", Age: 32}

	rpc_data := &RPCData{"f1", []interface{}{1, true, p}}

	// enc
	bs, err := encode(rpc_data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(bs)

	// dec
	rpcObj, err := decode(bs)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(rpcObj)
}
