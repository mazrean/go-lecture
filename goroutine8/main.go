package main

import (
	"context"
	"fmt"
	"time"
)

func despair(ctx context.Context) {
	innerCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	for {
		fmt.Println(innerCtx.Value("nyan").(string))

        select {
        case <- innerCtx.Done():
            fmt.Println("...")
            return
        }
	}
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "nyan", "pasu~")
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	go despair(ctx)

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}