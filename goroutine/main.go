package main

import (
	"fmt"
	"time"
)

func main() {

	start()
	time.Sleep(time.Second * 10)

}

func start() {
	go g1()
	go g2()
	fmt.Println("start end")
}

func g1() {
	for true {
		time.Sleep(time.Second * 2)
		fmt.Println("g1 is alive")
	}
}

func g2() {
	for true {
		time.Sleep(time.Second * 2)
		fmt.Println("g2 is alive")
	}
}
