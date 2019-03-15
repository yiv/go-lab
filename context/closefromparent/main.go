package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	father()
}

func father() {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	son(ctx, wg)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
}

func son(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			fmt.Println("son close")
		}()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Second)
				fmt.Println("son")
			}
		}

	}()
	grandSon(ctx, wg )
}

func grandSon(ctx context.Context, wg *sync.WaitGroup) {
	go func() {
		defer func() {
			wg.Done()
			fmt.Println("grand son close")
		}()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Second)
				fmt.Println("grand son")
			}
		}

	}()
}
