package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A: %d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B: %d\n", i)
	}
}

func c(i int) {
	defer wg.Done()
	fmt.Printf("C: %d\n", i)
}

func d(i int) {
	defer wg.Done()
	fmt.Printf("D: %d\n", i)
}
func main() {
	// // GOMAXPROCS sets the maximum number of CPUs that can be executing
	// simultaneously and returns the previous setting.
	runtime.GOMAXPROCS(2)
	fmt.Printf("cpus: %d\n", runtime.NumCPU())
	// wg.Add(2)
	// go a()
	// go b()
	// wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go c(i)
		go d(i)
	}

	wg.Wait()
}
