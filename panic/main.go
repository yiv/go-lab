package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func main() {
	PanicInGoroutine()
}

func PanicInGoroutine() {
	errChan := make(chan error)
	defer recoveryTablePanic()
	go func() {
		defer recoveryTablePanic()
		time.Sleep(time.Second * 5)
		panic("go panic")
	}()
	<-errChan
}

func TestPanic() {
	defer recoveryTablePanic()
	var err error
	fmt.Println(err.Error())
}

func recoveryTablePanic() {
	if err := recover(); err != nil {
		fmt.Println("edwin 15", err)
		fmt.Println("edwin #11", string(debug.Stack()))
	}
	return
}
