package main

import "fmt"

type User struct {
	Nick string
}

func main()  {
	AsignToStuct()
}

func AsignToStuct()  {
	list := map[string]User{
		"edwin":{Nick:"edwin"},
	}
	fmt.Println(list)
	//list["edwin"].Nick = "haha" // cannot assign to struct field list["edwin"].Nick in map
}