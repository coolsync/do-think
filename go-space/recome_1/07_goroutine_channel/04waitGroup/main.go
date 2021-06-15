package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f1() {
	// Seed uses the provided seed value to initialize the default Source to a
	// deterministic state.
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10) // Intn returns, as an int, a non-negative pseudo-random number in [0,n)

		fmt.Println(r1, r2)
	}
}

var wg sync.WaitGroup

func f2(i int) {
	defer wg.Done() // Done decrements the WaitGroup counter by one.
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(5)))
	fmt.Println(i)
}

func main() {
	for i := 0; i < 10; i++ {
		go f2(i)
		wg.Add(1) // Add adds delta, which may be negative, to the WaitGroup counter.
	}

	wg.Wait() // Wait blocks until the WaitGroup counter is zero.
}
