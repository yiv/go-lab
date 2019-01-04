package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	WithCancel()
}

func WithTimeOut() {

	work := func(ctx context.Context, id int, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("work #", id, " done ", ctx.Err())
				return
			default:
				fmt.Println("work #", id)
				time.Sleep(time.Second)
			}
		}
	}

	wg := &sync.WaitGroup{}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	wg.Add(1)
	go work(ctx, 0, wg)
	wg.Wait()

}

func WithCancel() {

	work := func(ctx context.Context, id int, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("work #", id, " done ", ctx.Err())
				return
			default:
				fmt.Println("work #", id)
				time.Sleep(time.Second)
			}
		}
	}

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go work(ctx, i, wg)
	}
	<-time.After(time.Second * 5)
	cancel()
	wg.Wait()

}
