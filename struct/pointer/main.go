package main

import "fmt"

type Data struct {
	Name string
}
func main()  {
	pointer := &Data{Name:"hahah"}
	value := *pointer
	fmt.Println(value)
}
