package main

import (
	"context"
	"fmt"
)

func main() {
	type contextKey string

	f := func(ctx context.Context, k contextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found key: ", v)
			return
		}
		fmt.Println("not found key")
	}

	k := contextKey("lang")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, contextKey("color"))
}

// type TraceCode string

// func worker(ctx context.Context) {
// 	key := TraceCode("hello")
// 	val, ok := ctx.Value(key).(string)
// 	if !ok {
// 		fmt.Println("invalid key")
// 	}

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println(ctx.Err())
// 			return
// 		case <-time.After(time.Millisecond * 70):
// 			fmt.Println("value: ", val)
// 			return
// 		}
// 	}
// }

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
// 	ctx = context.WithValue(ctx, TraceCode("hello"), "123123")
// 	go worker(ctx)
// 	time.Sleep(time.Second * 2)
// 	cancel()
// }
