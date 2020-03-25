package main

import (
	"fmt"
	"time"
)

func main() {
	date := "2017-11-22"
	start, e := time.ParseInLocation("2006-01-02", date, time.FixedZone("CST", 60*60*8))
	if e != nil {
		fmt.Println("err", e.Error())
		return
	}
	fmt.Println("err", start)
	fmt.Println("err", start.Unix())
}
