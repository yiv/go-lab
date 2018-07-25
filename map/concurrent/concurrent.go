package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//if one goroutine is writing to a map, no other goroutine should be reading (which includes iterating) or writing the map concurrently

func main() {
	wg := sync.WaitGroup{}

	oneMap := map[int32]string{
		5: "go",
		4: "java",
		3: "php",
	}

	for i := 0; i < 50; i++ {
		wg.Add(2)
		go readMap(oneMap, &wg)
		go writeMap(oneMap, &wg)
	}
	wg.Wait()
}

func readMap(oneMap map[int32]string, wg *sync.WaitGroup) {
	for i, s := range oneMap {
		fmt.Println("i = ", i, "s = ", s)
	}
	wg.Done()
	return
}
func writeMap(oneMap map[int32]string, wg *sync.WaitGroup) {
	rand.Seed(time.Now().UnixNano())
	test := []string{
		"c",
		"c++",
		"rust",
	}
	for _, s := range test {
		oneMap[int32(rand.Intn(10000))] = s
	}
	wg.Done()
	return
}
