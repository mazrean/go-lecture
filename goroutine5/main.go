package main

import (
	"context"
	"fmt"
	"time"
)

func despair(ctx context.Context) {
	for {
		fmt.Println("Help me!")
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go despair(ctx)

	time.Sleep(time.Second * 1)
	cancel()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		time.Sleep(time.Millisecond * 100)
	}
}
