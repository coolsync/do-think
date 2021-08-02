package main

import (
	"fmt"
	"time"
)

func main() {
	go sppider(time.Millisecond * 500)
	const n = 45
	total := fib(n)
	fmt.Printf("\rfib number result: %d\n", total)
}

func sppider(delay time.Duration) {
	for {

		for _, v := range `-\|/` {
			fmt.Printf("\r%c", v)
			time.Sleep(delay)
		}

	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
