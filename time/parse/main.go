package main

import (
	"fmt"
	"time"
)

func main() {
	date := "2017-11-22"
	start, e := time.Parse("2006-01-02", date)
	if e != nil {
		fmt.Println("err", e.Error())
		return
	}
	fmt.Println("err", start)
}
