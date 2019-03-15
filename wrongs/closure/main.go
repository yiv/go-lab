package main


import (
	"fmt"
	"runtime"
	"time"
)
func main() {
	MethodWithClosure()
}

type User struct {
	Nick string
}

func (u *User)Println()  {
	fmt.Println(u.Nick)
}

func MethodWithClosure()  {
	runtime.GOMAXPROCS(1)
	data := []User{ {"one"},{"two"},{"three"} }
	for _,v := range data {
		//v := v
		go v.Println()
	}
	time.Sleep(3 * time.Second)
}

func RangeWithClosure()  {
	//data := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21}
	runtime.GOMAXPROCS(1)
	data := []int{1,2,3,4,5,6,7,8,9,10,11,12,13}
	for _, v := range data {
		fmt.Println("for ", v)
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(3 * time.Second)
}

func ForWithClosure()  {
	runtime.GOMAXPROCS(1)
	for v := 0; v < 50; v++ {
		fmt.Println("for ", v)
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(3 * time.Second)
}