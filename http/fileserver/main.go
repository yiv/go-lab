package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

var (
	port = flag.String("port", "80", "port")
)

func main() {
	flag.Parse()
	http.Handle("/", accessControl(http.FileServer(http.Dir("./"))))
	log.Println("Listening on ", *port)
	go func() {
		time.Sleep(time.Second)
		openbrowser(fmt.Sprintf("http://127.0.0.1:%v", *port))
	}()
	err := http.ListenAndServe(fmt.Sprintf(":%v", *port), nil)
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
