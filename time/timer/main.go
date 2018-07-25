package main

import (
	"fmt"
	"time"
)

func main() {
	errc := make(chan error)

	go timer()
	go stop(errc)
	fmt.Println(<-errc)
}
func timer() {
	count := 0
	timer1 := time.NewTimer(time.Second * 3)
	for {
		select {
		case <-timer1.C:
			timer1.Reset(time.Second * 3)
			fmt.Println("timer1 time out")
			if count >= 10 {
				timer1.Stop()
			}
			count++
		}
	}
}
func stop(err chan error) {
	timerstop := time.NewTimer(time.Minute * 3)
	<-timerstop.C
	err <- fmt.Errorf("haha")
}
