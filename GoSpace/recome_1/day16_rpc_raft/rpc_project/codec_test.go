package rpcimpl

import (
	"encoding/gob"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func TestCodec(t *testing.T) {
	gob.Register(Person{}) // struct 作为 interface pass, not need impl interface

	p := Person{Name: "hello", Age: 18}
	args := []interface{}{"hello", 1, 2, true, p}
	data := RPCData{Name: "func name", Args: args}

	b, err := encode(&data)
	if err != nil {
		t.Fatal(err)
	}

	// fmt.Println(string(b))

	// rpc_data, err := decode(b)
	_, err = decode(b)

	if err != nil {
		t.Fatal(err)
	}

	// fmt.Println(rpc_data)
}
