package twochan

import (
	"fmt"
	"time"
)

func main() {
	f1()
	select {}
}

func f1() {
	ch := make(chan int)
	go write(ch)
	go read(ch)

	fmt.Println("f1 end")

}
func write(ch chan int) {
	for i := 0; i < 10000; i++ {
		ch <- i
		time.Sleep(time.Second * 1)
	}
}
func read(ch chan int) {
	for {
		fmt.Println("#", <-ch)
	}

}
