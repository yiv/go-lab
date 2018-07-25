package test

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//if one goroutine is writing to a map, no other goroutine should be reading (which includes iterating) or writing the map concurrently

func main() {
	wg := sync.WaitGroup{}

	oneMap := &sync.Map{}
	oneMap.Store(5, "go")
	oneMap.Store(4, "java")
	oneMap.Store(3, "php")

	for i := 0; i < 50; i++ {
		wg.Add(2)
		go readMap(oneMap, &wg)
		go writeMap(oneMap, &wg)
	}
	wg.Wait()
}

func readMap(oneMap *sync.Map, wg *sync.WaitGroup) {
	oneMap.Range(func(key, value interface{}) bool {
		fmt.Println("i = ", key, "s = ", value)
		return true
	})
	wg.Done()
	return
}
func writeMap(oneMap *sync.Map, wg *sync.WaitGroup) {
	rand.Seed(time.Now().UnixNano())
	test := []string{
		"c",
		"c++",
		"rust",
	}
	for _, s := range test {
		oneMap.Store(int32(rand.Intn(10000)), s)
	}
	wg.Done()
	return
}
