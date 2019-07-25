package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	Modify()
}

func NilTest() {
	type User struct {
		Name  string
		Age   int32
		Label []string
	}

	var param interface{} = User{Name: "nick"}

	fmt.Println(reflect.TypeOf(param))
	//typ := reflect.TypeOf(param)
	value := reflect.ValueOf(param)
	for i := 0; i < value.NumField(); i++ {
		fmt.Println(value.Field(i).Kind())
	}
}

func Modify() {
	type Stu struct {
		Name   string
		Reward []int
	}
	s := Stu{Name: "aa"}
	var iter interface{} = &s
	ot := reflect.ValueOf(iter)
	ot = ot.Elem()
	fmt.Println(ot.CanAddr())
	count := ot.NumField()
	for x := 0; x < count; x++ {
		f := ot.Field(x)
		//ft := f.Type()
		fmt.Println(f.CanAddr())
		fk := f.Kind()
		switch fk {
		case reflect.String:
			f.SetString("wwwwwwww")
		case reflect.Slice:
			s := makeslice(f.Interface())
			f.Set(reflect.ValueOf(s))
		}
	}
	res, _ := json.Marshal(s)
	fmt.Println("final ", string(res))

}

func makeslice(slice interface{}) interface{} {
	newsliceval := reflect.MakeSlice(reflect.TypeOf(slice), 0, 0)
	newslice := reflect.New(newsliceval.Type()).Elem()
	newslice.Set(newsliceval)
	return newslice.Interface()
}
