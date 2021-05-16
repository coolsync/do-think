package main

import (
	"fmt"
)

// func hello(i int) {
// 	fmt.Printf("hello， %d\n", i)
// }

func main() {
	fmt.Println("main goroutine!")

	for i := 0; i < 10000; i++ {
		// go hello(i)
		go func(n int) {
			fmt.Printf("hello， %d\n", n)
		}(i)
	}
}
