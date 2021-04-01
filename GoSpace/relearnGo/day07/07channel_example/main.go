package main

import (
	"fmt"
)

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squraer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func main() {
	// 1. generate 100 个数 传入到 ch1
	// 2. 从ch1 read, 并将其平方
	// 3. main printer

	ch1 := make(chan int)
	ch2 := make(chan int)

	go counter(ch1)
	go squraer(ch2, ch1)

	printer(ch2)
}
