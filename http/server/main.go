package main

import (
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	errChan := make(chan struct{})
	go GetServer()
	time.Sleep(time.Second)
	go PostClient()
	go GetClient()
	<-errChan
}

func GetServer() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Fatal("17", err.Error())
		}
		log.Println("19", r.Form)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetClient() {
	resp, err := http.Get("http://127.0.0.1:8080?x=y")
	if err != nil {
		log.Fatal("36", err.Error())
	}
	log.Println("34", resp.Body)
}

func PostClient() {
	resp, err := http.PostForm("http://127.0.0.1:8080", url.Values{"key": {"Value"}, "id": {"123"}})
	if err != nil {
		log.Fatal("41", err.Error())
	}
	log.Println("43", resp.Body)
}
