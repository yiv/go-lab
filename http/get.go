package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	for i := 97; i <= 122; i++ {
		wg.Add(1)
		go get(i, i)
	}

	wg.Wait()

}

func get(start, end int) {
	defer func() {
		wg.Done()
	}()
	httpCli := http.Client{Timeout: time.Duration(10 * time.Second)}

	for i := start; i <= end; i++ {
		for j := 97; j <= 122; j++ {
			for x := 97; x < 122; x++ {
				name := string(i) + string(j) + string(x)
				url := fmt.Sprintf("https://github.com/%s", name)
				resp, err := httpCli.Get(url)
				if err == nil {
					if resp.Status != "200 OK" {
						fmt.Println(name + ": " + resp.Status)
					} else {
						// fmt.Println(name + ": " + resp.Status)
					}
				} else {
					// fmt.Println("error :", name)
				}
			}
		}
	}

}
