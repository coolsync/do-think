package main

import (
	"fmt"
	"net"
)

func main() {
	// conn, err := net.Dial("tcp", "scanme.nmap.org:80")
	// conn, err := net.Dial("tcp", "www.baidu.com:80")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("connect ok")
	// defer conn.Close()

	for i := 1; i <= 1024; i++ {
		// address := fmt.Sprintf("scanme.nmap.org:%d", i)
		address := fmt.Sprintf("www.tlyz.net:%d", i)
		// fmt.Println(address)
		conn, err := net.Dial("tcp", address)

		if err != nil {
			// fmt.Printf("%d close and filter\n", i)
			continue
		}
		conn.Close()

		fmt.Printf("%d open\n", i)
	}
}
