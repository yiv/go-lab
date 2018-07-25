package main

import "fmt"

type bodyIf interface {
	myhead()
	myleg()
}

type head struct {
	a int
	b int
}

func (h *head) myhead() {
	fmt.Println(h.a)
}

type leg struct {
	x int
	y int
}

func (l *leg) myleg() {
	fmt.Println(l.x)
}

type body struct {
	head
	leg
}

func main() {
	ed := &body{head: head{5, 20}}
	fmt.Println(ed)
	ed.y = 50
	ed.x = 500
	fmt.Println(ed)
	ed.myhead()
	var edi bodyIf
	edi = ed
	edi.myleg()
}
