package main

import (
	"fmt"
	"sync"
)

var ch chan int
var wg sync.WaitGroup

func unbufChan() {
	fmt.Println(ch) // nil

	ch := make(chan int) // must 初始化 才能使用

	fmt.Println(ch) // 0xc000112000

	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-ch
		fmt.Printf("recv data from channel %v\n", x)
	}()

	ch <- 8 // sent data from channel
	fmt.Println("sent 8 to ch")

	wg.Wait()
}

func bufChan() {
	ch := make(chan int, 2)

	ch <- 8
	fmt.Println("sent 8 to ch")

	ch <- 10
	fmt.Println("sent 10 to ch")

	close(ch) // sent done, close channel

	for {
		// x := <-ch
		if x, ok := <-ch; ok {
			fmt.Printf("recv from ch %d\n", x)
		} else {
			break
		}
	}
}
func main() {
	bufChan()
}
