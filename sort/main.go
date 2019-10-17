package main

import (
	"fmt"
	"sort"
)

func main() {
	sort_by_method()
}

type Student struct {
	Name string
	Age  int
}

type Students []Student

func (a Students) Len() int           { return len(a) }
func (a Students) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Students) Less(i, j int) bool { return a[i].Age > a[j].Age }

func sort_by_method() {
	s1 := []Student{{Name: "edwin", Age: 5}, {Name: "watt", Age: 1}, {Name: "padme", Age: 4}}
	s2 := Students(s1)
	sort.Sort(s2)
	fmt.Println(s2)
}

func sort_by_closure()  {
	s1 := []Student{{Name: "edwin", Age: 5}, {Name: "watt", Age: 1}, {Name: "padme", Age: 4}}
	byAge := func(p1, p2 *Student) bool {
		return p1.Age < p2.Age
	}
	sort
}