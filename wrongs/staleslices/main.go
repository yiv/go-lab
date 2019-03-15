package main

import "fmt"

func main() {
	ary := []int{1, 2, 3}
	fmt.Println(len(ary), cap(ary), ary) //prints 3 3 [1 2 3]
	s2 := ary[1:]
	fmt.Println(len(s2), cap(s2), s2) //prints 2 2 [2 3]

	for i := range s2 {
		s2[i] += 20
	}
	//still referencing the same array
	fmt.Println(len(ary), cap(ary), ary) //prints 3 3 [1 22 23]
	fmt.Println(len(s2), cap(s2), s2) //prints 2 2 [22 23]


	//在一个slice中添加新的数据，在原有数组无法保持更多新的数据时，将导致分配一个新的数组。而现在其他的slice还指向老的数组（和老的数据）
	s2 = append(s2,4)
	for i := range s2 { s2[i] += 10 }
	fmt.Println(len(ary), cap(ary), ary) //prints 3 3 [1 22 23]
	fmt.Println(len(s2), cap(s2), s2) //prints 3 4 [32 33 14]
}
