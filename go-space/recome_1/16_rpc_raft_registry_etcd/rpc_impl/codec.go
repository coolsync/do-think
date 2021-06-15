package rpcimpl

import (
	"bytes"
	"encoding/gob"
)

// def rpc data struct
type RPCData struct {
	Name string
	Agrs []interface{}
}

// rpc data struct ecoder
func encode(rpc_data *RPCData) ([]byte, error) {
	// bytes buffer
	var buf bytes.Buffer

	// rpc data ecoder
	bufEnc := gob.NewEncoder(&buf)

	if err := bufEnc.Encode(&rpc_data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// bytes decoder
func decode(b []byte) (*RPCData, error) {
	buf := bytes.NewBuffer(b) //  retrun pointer buf

	var rpc_data *RPCData

	// bytes decoder
	bufDec := gob.NewDecoder(buf)

	// decode
	if err := bufDec.Decode(&rpc_data); err != nil {
		return nil, err
	}

	return rpc_data, nil
}
