package main

import (
	"context"
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"sync"
	"time"
)

var addr = flag.String("addr", ":10050", "http service address")

func main() {
	flag.Parse()
	wg := sync.WaitGroup{}
	// Mechanical domain.
	now := time.Now()
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client()
		}()
	}

	wg.Wait()
	log.Println("took: ", time.Now().Sub(now))

}

func client() {
	wg := sync.WaitGroup{}
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				err = c.WriteMessage(websocket.TextMessage, []byte("hello"))
				if err != nil {
					log.Println("write err:", err)
					return
				}
			}

		}
	}()

	wg.Add(1)
	go func() {
		defer func() {
			c.Close()
			wg.Done()
		}()
		count := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				_, _, err := c.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					return
				}
				count++
				if count > 10000 {
					cancel()
					log.Println("read: ", count)
					return
				}
			}
		}
	}()
	wg.Wait()
}
