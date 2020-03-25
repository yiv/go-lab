package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	TestHour()
}

func TestZone() {
	if loc, err := time.LoadLocation("Asia/Shanghai"); err != nil {
		log.Fatal(err.Error())
	} else {
		t := time.Now().In(loc)
		fmt.Println(t)
	}
}

func TestFixZone() {
	fmt.Println(time.Now().In(time.FixedZone("Asia/Shanghai", 0)))
	fmt.Println(time.Now().In(time.FixedZone("CST", 60*60*8)))
}

func TestHour() {
	fmt.Println(time.Now().Hour())
}
