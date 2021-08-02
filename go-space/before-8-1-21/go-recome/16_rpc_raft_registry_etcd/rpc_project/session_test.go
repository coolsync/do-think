package rpcimpl

import (
	"fmt"
	"log"
	"net"
	"sync"
	"testing"
)

func TestSession(t *testing.T) {
	addr := "localhost:8081"
	my_data := "hello"

	var wg sync.WaitGroup
	wg.Add(2)
	// write routine
	go func() {
		defer wg.Done()
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatal("net.Listen", err)
		}
		defer lis.Close()
		conn, _ := lis.Accept()
		s := NewSession(conn)
		err = s.Write([]byte(my_data))
		if err != nil {
			log.Fatal("s.Write: ", err)
		}
	}()

	// read routine
	go func() {
		defer wg.Done()
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			log.Fatal("net.Dial: ", err)
		}

		s := NewSession(conn)
		data, err := s.Read()
		if err != nil {

			log.Fatal("s.Read(): ", err)
		}

		if my_data != string(data) {
			log.Fatal("my_data valid: ", err)
		}
		fmt.Println(data)
	}()
	wg.Wait()
}
