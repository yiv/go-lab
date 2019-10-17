package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Stu struct {
	Name string
	Age  int
}

var StuPool = sync.Pool{New: func() interface{} {
	return &Stu{}
}}

func main() {
	for x := 0; x < 100; x++ {
		fmt.Println(UsePool())
	}
	
}

func UsePool() int {
	stu := StuPool.Get().(*Stu)
	stu.Age = int(rand.Int31n(100))
	age := stu.Age
	StuPool.Put(stu)
	return age
}


