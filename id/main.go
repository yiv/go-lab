package main

import (
	"fmt"
	"log"
	"github.com/sony/sonyflake"
)

func main() {
	var st sonyflake.Settings
	sf := sonyflake.NewSonyflake(st)
	for i := 0; i < 1000; i ++ {
		if id, err := sf.NextID(); err != nil {
			log.Fatal(err.Error())
			return
		} else {
			fmt.Println(id)
		}
	}

}
