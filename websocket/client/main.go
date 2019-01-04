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

var addr = flag.String("addr", "192.168.1.12:999", "http service address")
var logger log.Logger
var sum int

func main() {
	flag.Parse()

	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestamp)
	logger = log.With(logger, "caller", log.DefaultCaller)

	wg := sync.WaitGroup{}
	// Mechanical domain.
	now := time.Now()
	for i := 0; i < 20000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client(i)
		}()
	}

	wg.Wait()
	level.Info(logger).Log("sum", sum, "took ", time.Now().Sub(now))
}

func client(id int) {
	wg := sync.WaitGroup{}
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		level.Error(logger).Log("id", id, "err ", err.Error())
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	count := 0

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Second)
				c.SetWriteDeadline(time.Now().Add(time.Second * 10))
				err = c.WriteMessage(websocket.TextMessage, []byte("hello"))
				if err != nil {
					//level.Error(logger).Log("id", id, "err ", err.Error())
					return
				}
				//time.Sleep(time.Second)
			}

		}
	}()

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				c.SetReadDeadline(time.Now().Add(time.Second * 10))
				_, _, err := c.ReadMessage()
				if err != nil {
					level.Error(logger).Log("id", id, "err ", err.Error())
					return
				}
				if count > 10 {
					cancel()
					return
				}
				count++
				sum++
			}
		}
	}()

	wg.Wait()

	c.SetWriteDeadline(time.Now().Add(time.Second * 10))
	c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	time.Sleep(time.Second * 10)
	c.Close()
	level.Info(logger).Log("id", id, "done ", count)
}
