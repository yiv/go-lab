package main

import "fmt"

//map是引用类型，类似指针，slice也是，将map作为参数传入，在函数内
func main() {
	mm := map[int]int{
		5: 5,
		4: 2,
	}
	fmt.Println("#1 ", mm)
	modifyMap(mm)
	fmt.Println("#2 ", mm)
}
func modifyMap(m map[int]int) {
	m[6] = 5
}
