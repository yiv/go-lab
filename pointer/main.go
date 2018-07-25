package main

import "fmt"

func main() {
	a := 5
	b := 6
	test1(a)
	test2(&b)
	fmt.Println(a)
	fmt.Println(b)

}
func test1(a int) {
	a++
}
func test2(b *int) {
	*b++
}
