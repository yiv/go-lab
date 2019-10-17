package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

const (
	Host = "127.0.0.1"
	Port = 80
)

func main() {
	http.Handle("/", accessControl(http.FileServer(http.Dir("./"))))
	log.Println("Listening on ", Port)
	go func() {
		time.Sleep(time.Second)
		openbrowser(fmt.Sprintf("http://%v:%v", Host, Port))
	}()
	err := http.ListenAndServe(fmt.Sprintf(":%v", Port), nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			res := `{"code":200}`
			w.Write([]byte(res))
			return
		}
		fmt.Println(r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

func openbrowser(url string) {
	var err error
	
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Println(err)
	}
	
}
