package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	for p := range ports {
		// address := fmt.Sprintf("scanme.nmap.org:%d", j)
		address := fmt.Sprintf("www.tlyz.net:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int) 

	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	// results
	for i := 1; i <= 1024; i++ {
		port := <-results
		// fmt.Println(port)
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)

	for _, p := range openports {
		fmt.Printf("%d open\n", p)
	}
}
