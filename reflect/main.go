package main

import (
	"fmt"
	"golang.org/x/net/context"
	"time"
)

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Millisecond)
	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}
