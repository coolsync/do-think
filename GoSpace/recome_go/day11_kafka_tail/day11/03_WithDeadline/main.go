package main

import (
	"context"
	"fmt"
	"time"
)

// const shortDuration = time.Millisecond * 100
const shortDuration = time.Second * 5

func main() {
	d := time.Now().Add(shortDuration)

	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(time.Second * 2):
		fmt.Println("sleep so long")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // context deadline exceeded
	}
}
