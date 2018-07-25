package main

import (
	"net/http"
	"log"
	"fmt"
	"time"
	"crypto/md5"
	"io"
	"os"
	"strconv"
	"html/template"
)

func main() {
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":6060", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)


		fmt.Printf("other form : %v ", r.Form["file_path"])
	}
}
