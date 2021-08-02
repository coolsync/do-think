package main

import (
	"fmt"
	"net"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			// address := fmt.Sprintf("scanme.nmap.org:%d", j)
			address := fmt.Sprintf("www.tlyz.net:%d", j)

			// fmt.Println(address)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}

	wg.Wait()
	// time.Sleep(time.Second*30)
}
