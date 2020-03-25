package main

import (
	"fmt"
	"os"
)

func main() {
	err := TestErr()
	fmt.Println(err)
}

func TestErr() (err error) {
	x, err := NewErr()
	if err != nil {
		fmt.Println("16", err.Error())
		os.Exit(1)
	}
	fmt.Println(x)
	return err
}

func NewErr() (int32, error) {
	return 100, fmt.Errorf("a err")
}
