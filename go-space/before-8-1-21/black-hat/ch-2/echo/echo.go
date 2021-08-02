package main

import (
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 512)

	for {
		size, err := conn.Read(buf[:])
		if size == 0 {
			return
		}

		if err != nil {
			log.Fatal(err)
		}
		log.Printf("received %d bytes: %s\n", size, string(buf[:size]))

		// back write to client
		log.Println("writing data")
		if _, err = conn.Write(buf[:size]); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	liser, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("net port unable!")
	}
	log.Println("listen tcp :20080 ok")

	for {
		conn, err := liser.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go echo(conn)
	}
}
