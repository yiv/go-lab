package main

import (
	"fmt"
	"sync"
)

type Test struct {
	Total int64
}

func (t Test) Incr() {
	fmt.Printf("edwin #12 %p\n", &t)
	defer func() {
		t.Total++
		fmt.Printf("edwin #15 %p\n", &t)
	}()
}

func main() {
	test := Test{}
	test.Incr()
	fmt.Printf("edwin #25 %p\n", &test)
	wg := &sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for x := 0; x < 1; x++ {
				fmt.Printf("edwin #35 %p\n", &test)
				test.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println("edwin 31", test.Total)
}
