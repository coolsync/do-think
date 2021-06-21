package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func main() {
	conn, err := net.Dial("tcp", ":20080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	
	cmd := exec.Command("ls", "-l /")
	rp, wp := io.Pipe()

	cmd.Stdin = 
	cmd.Stdout = conn
	go io.Copy(conn, wp)

	// cmd.Stdout = conn
}