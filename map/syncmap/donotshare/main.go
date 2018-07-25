package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SS struct {
	m *sync.Map
}

func (s SS) Write() {
	rand.Seed(time.Now().Unix())
	k := rand.Int31n(8)
	s.m.Store(k, time.Now().Unix())
}
func (s SS) Read() {
	rand.Seed(time.Now().Unix())
	k := rand.Int31n(8)
	if v, ok := s.m.Load(k); ok {
		fmt.Printf("k = %d, v = %d\n", k, v)
	}
}
func New() SS {
	s := SS{
		m: &sync.Map{},
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
			ss.Write()
		}()
	}
	for i := 0; i < 5000000; i++ {
		go func() {
			defer wg.Done()
			ss.Read()
		}()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start).Seconds())
}
