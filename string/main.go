package main

import (
	"strings"
	"fmt"
)

func main() {
	str1 := "xGPA.3308-2660-0538-10888"
	b := strings.HasPrefix(str1, "GPA")
	fmt.Println(b)
}
