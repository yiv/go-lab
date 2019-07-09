package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Println("err", err.Error())
		}
		fmt.Println()
		var fields []string
		for k, v := range r.Form {
			if len(v) > 0 && v[0] != "" {
				fields = append(fields, k)
				fmt.Println(k, v)
			}
		}
		fmt.Println("Hello from a HandleFunc #1!")
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
