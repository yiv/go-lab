package main

import "fmt"

func main() {
	InterfaceWithHiddenVariable()
}

func InterfaceWithHiddenVariable()  {
	var data interface{} = "great"
	if data, ok := data.(int); ok {
		fmt.Println("[is an int] value =>",data)
	} else {
		fmt.Println("[not an int] value =>",data)
		//prints: [not an int] value => 0 (not "great")
	}

}

func NilInterface()  {
	var data *byte
	var in interface{}
	fmt.Println(data, data == nil) //prints: <nil> true
	fmt.Println(in, in == nil)     //prints: <nil> true
	in = data
	fmt.Println(in, in == nil) //prints: <nil> false
	//'data' is 'nil', but 'in' is not 'nil'
}