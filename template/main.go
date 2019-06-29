package main

import (
	"os"
	"text/template"
)

func main() {
	testRangeMap()
}

func ifelse() {
	//ary := []string{"nick", "edwin", "padme"}
	type Data struct {
		X int
		Y int
	}
	data := Data{X: 1, Y: 2}
	tmpl, _ := template.New("test").Parse("{{if eq .X .Y}} x  {{else}} y {{end}}")
	tmpl.Execute(os.Stdout, data)
}
func testRangeSlice() {
	ary := []string{"nick", "edwin", "padme"}
	tmpl, _ := template.New("test").Parse("{{range $v := .}} value = {{ $v }} {{end}}")
	tmpl.Execute(os.Stdout, ary)
}
func testRangeMap() {
	ary := map[string]int{"nick": 5, "edwin": 10, "padme": 3}
	tmpl, _ := template.New("test").Parse("{{range $k,$v := .}} {{ $k }} = {{ $v }} {{end}}")
	tmpl.Execute(os.Stdout, ary)
}

func test() {
	type Inventory struct {
		Material string
		Count    uint
	}
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}
