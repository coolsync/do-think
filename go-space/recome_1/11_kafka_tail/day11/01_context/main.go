package main

import (
	"context"
	"fmt"
	"time"
)

// main --> worker1 --> worker2
// var wg sync.WaitGroup

func worker2(ctx context.Context) {
	for {
		fmt.Println("mark")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			// break
			return
		default: // 什么都不做， for 继续执行
		}
	}
}

func worker(ctx context.Context) {
	// defer wg.Done()
	go worker2(ctx)
	for {
		fmt.Println("bob")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			// break
			return
		default: // 什么都不做， for 继续执行
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel() // notify 子 goroutine end
	// wg.Wait()
}
