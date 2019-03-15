package main

import (
	"fmt"
	"time"
)

var Card2Value = map[byte]byte{
	0x01: 1,
	0x02: 2,
	0x03: 3,
	0x04: 4,
	0x05: 5,
}

func main() {
	ShiftVSMap()
}

func ShiftVSMap() {
	var card byte = 0x15
	start := time.Now()
	for i := 0; i < 10000000; i ++ {
		GetValueByShift(card)
	}
	fmt.Println("time used ", time.Now().Sub(start))

	start = time.Now()
	for i := 0; i < 10000000; i ++ {
		GetValueByMap(card)
	}
	fmt.Println("time used ", time.Now().Sub(start))
}

func GetValueByShift(card byte) (value byte) {
	value = card & 0xf
	return
}
func GetValueByMap(card byte) (value byte) {
	value = Card2Value[card]
	return
}

