package main

import (
	"fmt"
	"time"
)

func main() {
	unixNow := time.Now().Unix()
	fmt.Println(unixNow)
}
