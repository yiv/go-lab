//try to use "=" to copy value of different struct type
//fail
package main

import "fmt"

type big struct {
	a string
	b string
}
type small struct {
	a string
}

func main() {
	s := small{"aaa"}
	b := big{}
	b = big(s)
	fmt.Println(b)
}
