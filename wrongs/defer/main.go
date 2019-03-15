package main

import "fmt"

func main()  {
	TestDefer1()
}

func TestDefer1()  {
	for i := 0; i < 5; i++{
		defer InvokeInDefer(i)
	}
	defer InvokeInDefer(9999)
}

func InvokeInDefer(msg int)  {
	fmt.Println("InvokeInDefer.....", msg)
}
