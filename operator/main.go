package main

import "fmt"

func main() {
	b := 23975
	fmt.Println((b / 10000000) * 10000000)
	fmt.Println((b % 10000000) / 100000)
	fmt.Println((b % 100000) / 10)
	fmt.Println(b / 1000)
	fmt.Println(b % 1000)
}
