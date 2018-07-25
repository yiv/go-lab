package main

import (
	"fmt"
	"time"
)

func main() {
	closeClosedChan()
}

func closeClosedChan() {
	var doneChan chan struct{}
	close(doneChan)
}

func ch1(share chan struct{}) {
	fmt.Println("ch1 start")
	timer := time.NewTimer(3 * time.Second)
	<-timer.C
	close(share)
	fmt.Println("ch1 exit")
}

func ch2(share chan struct{}) {
	fmt.Println("ch2 start")
	select {
	case <-share:
		fmt.Println("ch2 read share....")
	}
	fmt.Println("ch2 exit")
}
