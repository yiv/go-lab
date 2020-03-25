package main

import "fmt"

func main() {
	test1()
}
func test2() {
	xx := [20][]int{
		0:  {5},
		10: {5, 55},
	}
	fmt.Println(xx[0:10])
	fmt.Println(xx[10:20])
}

func test1() {
	xx := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} //, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}
	fmt.Println(xx[0:10])
	//fmt.Println(xx[10:20])
}
