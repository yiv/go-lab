package main

import (
	_ "edwin/test/package/pack1"
	_ "edwin/test/package/pack2"
	_ "edwin/test/package/pack3"
	"fmt"
)

func main() {
	i := 50
	f1()
	fmt.Println("main Test = ", i)
}

func f1() {
	x := 80
	fmt.Println(x)
}
