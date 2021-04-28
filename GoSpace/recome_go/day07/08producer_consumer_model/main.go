package main

import (
	"fmt"
	"math/rand"
	"time"
)

//  只是为 条件变量 铺垫
func producer(out chan<- int, i int) {
	for x := 0; x < 50; x++ {
		num := rand.Intn(800)
		fmt.Printf("生产者 %dth, 生产 %d\n", i, num)
		out <- num
	}
	// close(out)
}

// send on closed channel
func consumer(in <-chan int, i int) {
	for {
		x, ok := <-in
		if !ok {
			break
		}
		fmt.Printf("---- 消费者  %dth, 消费 %d\n", i, x)
	}
	// for x := range in {
	// 	fmt.Printf("---- 消费者  %dth, 消费 %d\n", i, x)
	// }
}

func main() {
	product := make(chan int)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 3; i++ {
		go producer(product, i+1)
	}

	for i := 0; i < 5; i++ {
		go consumer(product, i+1)
	}

	for {
		;
	}
}
