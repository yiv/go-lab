package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func main()  {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("error opening file")
		panic(err.Error())
	}
	defer file.Close()
	
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", "test.txt")
	if err != nil {
		fmt.Println("error writing to buffer")
		panic(err.Error())
	}
	
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		panic(err.Error())
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
}
