package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	log.Println("Listening on 80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
