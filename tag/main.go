package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/fatih/structtag"
)

type Student struct {
	Name   string `json:"name"    jsona:"name"`
	Age    int32  `json:"age"     jsona:"-"`
	Avatar string `json:"avatar"  jsona:"avatar"`
}

func main() {
	s := Student{"nick", 12, "haha"}
	if res, err := ExtractByTag(s, "jsona"); err != nil {
		panic(err.Error())
	} else {
		j, _ := json.Marshal(res)

		fmt.Printf("edwin #21 %v \n", string(j))
	}

}

func ExtractByTag(obj interface{}, tagName string) (res map[string]interface{}, err error) {
	res = make(map[string]interface{})
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	var (
		tags *structtag.Tags
		tag  *structtag.Tag
	)
	for i := 0; i < objT.NumField(); i++ {
		field := objT.Field(i)
		if tags, err = structtag.Parse(string(field.Tag)); err != nil {
			return
		}
		if tag, err = tags.Get(tagName); err != nil {
			return
		}
		if tag.Name == "-" {
			continue
		} else {
			res[tag.Name] = objV.Field(i).Interface()
		}
	}
	return
}

func IterateStructField() {
	type User struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	user := User{
		Id:    1,
		Name:  "John Doe",
		Email: "john@example",
	}

	// TypeOf returns the reflection Type that represents the dynamic type of variable.
	// If variable is a nil interface value, TypeOf returns nil.
	t := reflect.TypeOf(user)

	// Get the type and kind of our user variable
	fmt.Println("Type:", t.Name())
	fmt.Println("Kind:", t.Kind())

	// Iterate over all available fields and read the tag value
	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)

		// Get the field tag value
		tagString := field.Tag
		tags, err := structtag.Parse(string(tagString))
		if err != nil {
			fmt.Println(err.Error())
		}
		tagEle, _ := tags.Get("json")
		fmt.Printf("%#v \n", tagEle.Name)
	}
}
