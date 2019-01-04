package main

import (
	"context"
	"flag"
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/websocket"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/go-kit/kit/log"
)

var addr = flag.String("addr", "192.168.1.12:10050", "http service address")
var logger log.Logger

func main() {
	flag.Parse()

	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestamp)
	logger = log.With(logger, "caller", log.DefaultCaller)

	wg := sync.WaitGroup{}
	// Mechanical domain.
	now := time.Now()
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client()
		}()
	}

	wg.Wait()
	level.Info(logger).Log("took ", time.Now().Sub(now))
}

func client() {
	wg := sync.WaitGroup{}
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		level.Error(logger).Log("err ", err.Error())
		return
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
				time.Sleep(time.Second)
				err = c.WriteMessage(websocket.TextMessage, []byte("hello"))
				if err != nil {
					level.Error(logger).Log("err ", err.Error())
					return
				}
				//time.Sleep(time.Second)
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
					level.Error(logger).Log("err ", err.Error())
					return
				}
				count++
				if count > 10000 {
					cancel()
					level.Info(logger).Log("done ", count)
					return
				}
			}
		}
	}()
	wg.Wait()
}
