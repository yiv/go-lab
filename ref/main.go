package main

import "fmt"

func main() {
	dint := int32(50)
	slice := []int32{50}
	dmap := map[int32]int32{5: 50}
	modify(dint, slice, dmap)
	fmt.Println(dint, slice, dmap)
}

func modify(dint int32, slice []int32, dmap map[int32]int32) {
	dint = 5
	slice = append(slice, 5)
	dmap[5] = 5
}
