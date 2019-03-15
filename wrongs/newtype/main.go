package main

import "fmt"

type User struct {
	Nick string
}

func (u *User) Name() {
	fmt.Println(u.Nick)
}

type Student User

func main() {
	s := Student{Nick: "edwin"}
	//把一个现有（非interface）的类型定义为一个新的类型时，新的类型不会继承现有类型的方法
	//s.Name()
}
