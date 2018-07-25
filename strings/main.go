package main

import (
	"fmt"
	"strings"
)

func main()  {
	str := "1,22222"
	newstr := strings.Split(str,",")
	fmt.Println( newstr[0])
}
