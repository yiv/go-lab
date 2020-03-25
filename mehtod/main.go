package main

import "fmt"

type Test struct {
	Total  int32
	Total2 *int32
	slice  []int32
}

func (t *Test) Incr() {
	t.Total = 500
	*t.Total2 = 5000
	t.slice = append(t.slice, 500)
}
func (t Test) Incr2() {
	t.Total = 600
	*t.Total2 = 6000
	t.slice = append(t.slice, 600)

}

func main() {
	var zero int32 = 0
	t := Test{Total2: &zero}
	fmt.Println("edwin 18", t.Total, *t.Total2, t.slice)
	t.Incr()
	fmt.Println("edwin 19", t.Total, *t.Total2, t.slice)
	t.Incr2()
	fmt.Println("edwin 20", t.Total, *t.Total2, t.slice)
}
