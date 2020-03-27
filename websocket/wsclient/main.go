package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

var (
	addr  = flag.String("addr", "127.0.0.1:4314", "http service address")
	count = flag.Int("count", 1, "count of concurrency")
	delay = flag.Int("delay", 10, "delay of new connection(Microsecond)")
	freq  = flag.Int("freq", 10, "frequency of data send(Second)")
	round = flag.Int("round", 1000000, "round of data send/receive")
)

var (
	logger log.Logger
	sum    int
	mtx    sync.Mutex
)

func main() {
	flag.Parse()

	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestamp)
	logger = log.With(logger, "caller", log.DefaultCaller)

	wg := sync.WaitGroup{}
	// Mechanical domain.
	now := time.Now()

	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			_ = level.Info(logger).Log("sum", sum)
			time.Sleep(time.Second * 3)
		}
	}()

	for i := 0; i < *count; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			open(id)
		}(i)
		time.Sleep(time.Microsecond * time.Duration(*delay))
	}

	wg.Wait()
	_ = level.Info(logger).Log("sum", sum, "took ", time.Now().Sub(now))
}

func updateSum(i int) {
	mtx.Lock()
	sum += i
	mtx.Unlock()
}

func open(id int) {

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		_ = level.Error(logger).Log("id", id, "err ", err.Error())
		return
	}

	updateSum(1)

	defer func() {
		updateSum(-1)
	}()

	var (
		bytes int
		st    = time.Now()
	)

	level.Debug(logger).Log("open", id, "sum", sum)
	for i := 1; i < *round; i++ {
		msg := []byte(fmt.Sprintf("id= %v, %v", id, time.Now()))
		err = c.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			_ = level.Error(logger).Log("id", id, "err ", err.Error())
			return
		}
		_, msg, err = c.ReadMessage()
		if err != nil {
			_ = level.Error(logger).Log("id", id, "err ", err.Error())
			return
		}
		bytes += len(msg)
		//_ = level.Debug(logger).Log("id", id, "msg ", string(msg))
		if *freq > 0 {
			time.Sleep(time.Second * time.Duration(*freq))
		}

	}

	_ = level.Info(logger).Log("id", id, "round", *round, "bytes ", bytes, "time", time.Now().Sub(st))

}

func client(id int) {
	wg := sync.WaitGroup{}
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		_ = level.Error(logger).Log("id", id, "err ", err.Error())
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
				_ = c.SetWriteDeadline(time.Now().Add(time.Second * 10))
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
				_ = c.SetReadDeadline(time.Now().Add(time.Second * 10))
				_, _, err := c.ReadMessage()
				if err != nil {
					_ = level.Error(logger).Log("id", id, "err ", err.Error())
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

	_ = c.SetWriteDeadline(time.Now().Add(time.Second * 10))
	_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	time.Sleep(time.Second * 10)
	_ = c.Close()
	_ = level.Info(logger).Log("id", id, "done ", count)
}
