package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	pp := "abc/ddd/f.txt"
	r, _ := filepath.Abs(pp)
	fmt.Println(r)

	r = filepath.Base(pp)
	fmt.Println(r)

	r = filepath.Dir(pp)
	fmt.Println(r)
}
