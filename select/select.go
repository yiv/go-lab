package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int)
	select {
	case x := <- waitChan(intChan):
		fmt.Println(x)
	}

}

func waitChan(intChan chan int) chan int {
	x, _ := wait()
	intChan <- x
	return intChan
}
func wait() (int, error) {
	time.Sleep(time.Second * 5)
	return 5, nil
}
