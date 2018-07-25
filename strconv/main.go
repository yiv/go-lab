package main

import (
	"fmt"
	"strconv"
)

func main() {
	//v32 := "-354634382"
	//if s, err := strconv.ParseInt(v32, 10, 32); err == nil {
	//	fmt.Printf("%T, %v\n", s, s)
	//}
	//if s, err := strconv.ParseInt(v32, 16, 32); err == nil {
	//	fmt.Printf("%T, %v\n", s, s)
	//} else {
	//	fmt.Println(err)
	//}
	//
	//v64 := "257682991242215425"
	//if s, err := strconv.ParseInt(v64, 10, 64); err == nil {
	//	fmt.Printf("%T, %v\n", s, s)
	//} else {
	//	fmt.Println(err)
	//}
	//if s, err := strconv.ParseInt(v64, 16, 64); err == nil {
	//	fmt.Printf("%T, %v\n", s, s)
	//} else {
	//	fmt.Println(err)
	//}

	int64n := int64(257682991242215425)
	ss := strconv.FormatInt(int64n, 10)
	fmt.Printf("%T, %v\n", ss, ss)

	xx := strconv.ParseInt(ss,10,64)

}
