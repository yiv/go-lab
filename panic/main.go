package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func main() {
	for {
		TestPanic()
		time.Sleep(time.Second * 3)
	}
}

func TestPanic() {
	defer  recoveryTablePanic()
	var err error
	fmt.Println(err.Error())
}

func recoveryTablePanic() {
	if err := recover(); err != nil {
		fmt.Println(err)
		fmt.Println("edwin #11", string(debug.Stack()))
	}
	return
}
