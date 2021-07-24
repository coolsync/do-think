package rpcimpl

import (
	"bytes"
	"encoding/gob"
)

// def rpc data struct
type RPCData struct {
	Name string        // visit method
	Args []interface{} // params data
}

// encode rpc data
func encode(rpc_ata *RPCData) ([]byte, error) {
	var buf bytes.Buffer // Simple byte buffer for marshaling data.

	// 创建 []byte 编码器
	bufEnc := gob.NewEncoder(&buf)

	// struct rpc_ata 编码到 buf
	if err := bufEnc.Encode(&rpc_ata); err != nil {
		return nil, err
	}
	// 获取 buf 上 字节slice
	return buf.Bytes(), nil
}

// decode rpc data
func decode(b []byte) (*RPCData, error) {
	buf := bytes.NewBuffer(b)	// NewBuffer creates and initializes a new Buffer using b as its initial contents

	bufDec := gob.NewDecoder(buf)

	var rpc_data *RPCData

	if err := bufDec.Decode(&rpc_data); err != nil {
		return rpc_data, err
	}

	return rpc_data, nil
}

