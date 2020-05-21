package main

import (
	"fmt"
	"time"
)

const max = 100000000

func main() {
	bitCheck()
	arrayIndex()
}

func bitCheck() {
	x := 16
	n := time.Now()
	for i := 0; i < max; i++ {
		_ = x & 255
	}
	fmt.Println("bit: ", time.Now().Sub(n))
}

func arrayIndex() {
	ary := []int{0, 0, 0, 0, 1, 0, 0, 0}
	n := time.Now()
	for i := 0; i < max; i++ {
		_ = ary[4]
	}
	fmt.Println("array: ", time.Now().Sub(n))
}
