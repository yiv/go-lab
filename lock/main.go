package main

import (
	"fmt"
	"sync"
	"time"
)

type CC struct {
	mtx sync.RWMutex
	wg  sync.WaitGroup
}

func (c *CC) read() {
	c.mtx.Lock()
	defer func() {
		c.mtx.Unlock()
		c.wg.Done()
	}()
	time.Sleep(time.Second * 5)
	fmt.Println("xx")

}
func main() {
	i := 50
	c := &CC{}

	for x := 0; x < i; x++ {
		c.wg.Add(1)
		go c.read()
	}
	c.wg.Wait()
}
