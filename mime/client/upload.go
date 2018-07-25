package main

import (
	"mime/multipart"
	"io"
	"fmt"
	"bytes"
	"os"
	"net/http"
	"io/ioutil"
)


func main()  {
	err := postFile("./icon.jpg", "http://127.0.0.1:6060/upload")
	fmt.Println(err)
}

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}
	// open file handle
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	fieldWriter, err := bodyWriter.CreateFormField("file_path")
	if err != nil {
		fmt.Println("CreateFormField failed")
		return err
	}
	_, err = fieldWriter.Write([]byte("hahahahah"))
	if err != nil {
		fmt.Println("CreateFormField  Write failed")
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}