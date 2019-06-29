package main

import (
	"fmt"
	"reflect"
)

func main() {
	type User struct {
		Name  string
		Age   int32
		Label []string
	}

	NilTest(User{Name: "nick"})

}
func NilTest(param interface{}) {
	fmt.Println(reflect.TypeOf(param))
	//typ := reflect.TypeOf(param)
	value := reflect.ValueOf(param)
	for i := 0; i < value.NumField(); i++ {
		fmt.Println(value.Field(i).Kind())
	}
}
