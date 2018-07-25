package main

import (
	"fmt"
	"time"
)

func main() {
	go input()
	go output()
	select {}
}

func input() {
	var cmdCode uint32
	for {
		fmt.Scanln(&cmdCode)
		fmt.Println("cmd code : ", cmdCode)
	}
}

func output() {
	x := time.Tick(1 * time.Second)
	for {
		<-x
		fmt.Println("tick.....")
	}
}
