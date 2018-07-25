package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	go func() {
		for {
			select {
			case t := <-tick:
				fmt.Println(t.String())
			}
		}

	}()
	select {}
}
