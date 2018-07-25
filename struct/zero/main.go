package main

import (
	"fmt"
)

type Name struct {
	Age     int32
	Balance int32
}

func main() {
	x := &Name{Age: 40, Balance: 50}
	y := &Name{Age: 40, Balance: int32(0)}
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
}
