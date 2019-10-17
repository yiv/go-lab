package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Stu struct {
	Age                                                                                                                int
	P1, P2, P3, P4, P5, P6, P7, P8, P9, P10, P11, P12, P13, P14, P15, P16, P17, P18, P19, P20, P21, P22, P23, P24, P25 People
}
type People struct {
	Name, Name1, Name2, Name3, Name4, Name5, Name6, Name7, Name8, Name9, Name10, Name11, Name12                                                                                                                                                            string
	Name13, Name14, Name15, Name16, Name17, Name18, Name19, Name20, Name21, Name22, Name23, Name24, Name25, Name26, Name27, Name28, Name29, Name30, Name31, Name32, Name33, Name34, Name35, Name36, Name37, Name38, Name39, Name40, Name41, Name42, Name43 string
}

var StuPool = sync.Pool{New: func() interface{} {
	return &Stu{}
}}

func main() {
	var sumAge int
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	
	for x := 0; x < 10000000000; x++ {
		sumAge += UsePool()
	}
	
	fmt.Println(sumAge)
	fmt.Println(time.Now().Sub(start))
}

func UsePool() int {
	stu := StuPool.Get().(*Stu)
	stu.Age = int(rand.Int31n(100))
	age := stu.Age
	StuPool.Put(stu)
	return age
}

func NotUsePool() int {
	stu := &Stu{}
	stu.Age = int(rand.Int31n(100))
	return stu.Age
}
