package main

import (
	"fmt"
	"sync"
)

var mtx sync.RWMutex

func main() {
	f1()


}

func f1() {
	if true {
		f2()
		return
	}

	mtx.Lock()
	defer func() {
		mtx.Unlock()
		fmt.Println("defer print")
	}()

}

func f2() {
	//mtx.Lock()
	defer mtx.Unlock()
	fmt.Println("f2")
}
