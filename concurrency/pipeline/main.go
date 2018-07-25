package main

import (
	"fmt"
	"time"
)

func gen(nums ...int) <-chan int {
	out := make(chan int, 2)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
		fmt.Println("gen close", out)

	}()
	return out
}

func sq(in <-chan int) <-chan int {
	time.Sleep(time.Second * 5)

	out := make(chan int)
	fmt.Println("sq begin", out)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
		fmt.Println("sq close")
	}()
	return out
}

func main() {
	// Set up the pipeline.
	c := gen(2, 3)
	out := sq(c)

	// Consume the output.
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9

	fmt.Println("main close")
}
