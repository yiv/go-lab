package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"./ss"
)

func Write(s ss.SS) {
	rand.Seed(time.Now().Unix())
	k := rand.Int31n(8)
	s.M.Store(k, time.Now().Unix())
}
func Read(s ss.SS) {
	rand.Seed(time.Now().Unix())
	k := rand.Int31n(8)
	if v, ok := s.M.Load(k); ok {
		fmt.Printf("k = %d, v = %d\n", k, v)
	}
}
func New() ss.SS {
	s := ss.SS{
		M: &sync.Map{},
	}
	return s
}
func main() {
	ss := New()
	wg := sync.WaitGroup{}
	wg.Add(10000000)
	start := time.Now()
	for i := 0; i < 5000000; i++ {
		go func() {
			defer wg.Done()
			Write(ss)
		}()
	}
	for i := 0; i < 5000000; i++ {
		go func() {
			defer wg.Done()
			Read(ss)
		}()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start).Seconds())
}
