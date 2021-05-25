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
	hello := "hello"

	var wg sync.WaitGroup
	wg.Add(2)
	
	// srv read routine
	go func() {
		defer wg.Done()
		// listen ip
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("Listen err:%v\n", err)
		}

		// wait cli send data
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// server session
		serSession := NewSession(conn)

		b, err := serSession.Read()
		if err != nil {
			log.Fatal(err)
		}

		if hello != string(b) {
			log.Fatal(err)
		}

		fmt.Println(b)
	}()
	// cli write routine
	go func() {
		defer wg.Done()
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			log.Fatal("dial: ", err)
		}

		cliSession := NewSession(conn)

		err = cliSession.Write([]byte(hello))
		if err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
}
