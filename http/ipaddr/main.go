package main

import (
	"log"
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", viewHandler)
	err := http.ListenAndServe(":777", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	//xRealIP := r.Header.RemoteAddr
	//xForwardedFor := r.Header.Get("X-Forwarded-For")
	fmt.Println(r.RemoteAddr)
	//fmt.Println(xForwardedFor)
}