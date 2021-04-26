package main

import (
	"fmt"
	"time"
)

// var wg sync.WaitGroup

// var done bool
var exitCh = make(chan struct{})

func f1() {
	// wg.Add(1)
	var n int
	for {
		n++
		fmt.Printf("%d, hello\n", n)
		time.Sleep(time.Millisecond * 500)
		// if done {
		// 	wg.Done()
		// }
		select {
		case <-exitCh:
			return
			// wg.Done()
		default:
		}
	}
}

func main() {
	// create a goroutine
	go f1()
	time.Sleep(time.Second * 5)
	// done = true
	exitCh <- struct{}{}
	// wg.Wait()
	fmt.Println("main ok")
}
