package main

import (
	"fmt"
	"net/http"

	"github.com/Masterminds/go-fileserver"
)

func main() {
	port := ":80"
	// Specity a NotFoundHandler to use when no file is found.
	fileserver.NotFoundHandler = func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, "That page could not be found.")
	}

	// Serve a directory of files.
	dir := http.Dir("./")
	fmt.Println("serve files on http ", port)
	http.ListenAndServe(port, fileserver.FileServer(dir))

}
