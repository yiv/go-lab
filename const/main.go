package main

import "fmt"

func main() {
	const (
		a int = iota + 1
		b
		c
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
