package main

import (
	"fmt"
	"github.com/sony/sonyflake"
	"log"
	"time"
)

func main() {
	t := time.Now()
	var st sonyflake.Settings
	sf := sonyflake.NewSonyflake(st)
	for i := 0; i < 1000000; i++ {
		if _, err := sf.NextID(); err != nil {
			log.Fatal(err.Error())
			return
		}
	}
	fmt.Println(time.Now().Sub(t))

}
