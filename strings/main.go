package main

import (
	"fmt"
	"strings"
)

func main() {
	TrimRight()
}

func TrimRight() {
	fmt.Println(strings.TrimRight("`ha, ha`ha, ", ","))
}
func Split() {
	str := "1,22222"
	newstr := strings.Split(str, ",")
	fmt.Println(newstr[0])
}
