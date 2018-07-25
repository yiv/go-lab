package main

import (
	"fmt"
	"time"
)

func main() {
	newArray()
}

func newArray() {
	var xx [5]int64
	for k := range xx {
		xx[k] = time.Now().Unix()
	}

	fmt.Println(xx)

}
func testArrayInt() (slice []int, p *[]int) {
	slice = make([]int, 100)
	slice[0] = 500
	p = &slice
	return slice, p
}
