package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	// write
	v := uint32(500)
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, v)

	fmt.Println(buf)
	// read
	x := binary.BigEndian.Uint32(buf)

	fmt.Println(x)
}
