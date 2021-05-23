package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = time.Millisecond * 5

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(time.Millisecond * 10):
		fmt.Println("done")
	}

}
