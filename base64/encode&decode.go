package main

import (
	//"os"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	fileBytes, err := ioutil.ReadFile("fb.jpg")
	errCheck(err)
	base64CodeStr := base64.StdEncoding.EncodeToString(fileBytes)
	fmt.Println(base64CodeStr)
	fmt.Println("len of str", len(base64CodeStr))
	fileBytes2, err := base64.StdEncoding.DecodeString(base64CodeStr)
	errCheck(err)
	err = ioutil.WriteFile("icon2.jpg", fileBytes2, 0664)
	errCheck(err)
}

func errCheck(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
