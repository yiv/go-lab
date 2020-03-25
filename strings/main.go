package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	Unicode()
}

func TrimRight() {
	fmt.Println(strings.TrimRight("`ha, ha`ha, ", ","))
}
func Split() {
	str := "1,22222"
	newstr := strings.Split(str, ",")
	fmt.Println(newstr[0])
}

func Contain() {
	str := "18200**3239"
	fmt.Println(strings.Contains(str, "*"))
}

func Unicode() {
	type RawUnicodeString string
	type Message struct {
		Code RawUnicodeString
	}
	data := `{"code":"\u5728\u4e30\u5fb7\u5c14Berro\u8212\u9002\u76841\u623f\u5355\u4f4d"}`
	var r Message
	json.Unmarshal([]byte(data), &r)
	fmt.Println(r.Code)
}
