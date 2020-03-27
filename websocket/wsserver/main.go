package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"sync"
	"time"

	//"io/ioutil"
	//"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	//"time"

	"github.com/gorilla/websocket"
	//"zycase.cn/shaoyong/wechat/agent/websocket"
)

var (
	webSocketAddr = flag.String("websocket.addr", ":4314", "game agent webSocket address")
	sum           = 0
	logger        log.Logger
	mtx           sync.Mutex
)

type Message struct {
	MType   int
	Content []byte
}

func main() {
	flag.Parse()

	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestamp)
	logger = log.With(logger, "caller", log.DefaultCaller)

	// Mechanical domain.
	errc := make(chan error)
	// Interrupt handler.
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()
	go func() {
		for {
			level.Info(logger).Log("sum", sum)
			time.Sleep(time.Second * 3)
		}

	}()
	go func() {
		m := http.NewServeMux()
		m.HandleFunc("/ws", webSocketWaitToRead)
		errc <- http.ListenAndServe(*webSocketAddr, m)
	}()
	level.Info(logger).Log("end", <-errc)
}

func updateSum(i int) {
	mtx.Lock()
	sum += i
	mtx.Unlock()
}

func webSocketWaitToRead(w http.ResponseWriter, r *http.Request) {
	defer updateSum(-1)
	var upgrader = websocket.Upgrader{} // use default options
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		level.Error(logger).Log("err", err.Error())
		return
	}
	updateSum(1)
	for {
		c.SetReadDeadline(time.Now().Add(time.Hour * 10))
		_, _, err := c.ReadMessage()
		if err != nil {
			level.Error(logger).Log("err", err.Error())
			return
		}
	}

}

func webSocketServerStd(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{} // use default options
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		level.Error(logger).Log("err", err.Error())
		return
	}
	sum++

	fmt.Println("webSocketServerStd new client, sum ", sum)
	ctx, cancel := context.WithCancel(context.Background())
	buf := make(chan Message, 10000)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				c.SetReadDeadline(time.Now().Add(time.Second * 10))
				mt, message, err := c.ReadMessage()
				if err != nil {
					level.Error(logger).Log("err", err.Error())
					cancel()
					return
				}
				buf <- Message{MType: mt, Content: message}
			}
		}
	}()

	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-buf:
				c.SetWriteDeadline(time.Now().Add(time.Second * 10))
				err = c.WriteMessage(msg.MType, msg.Content)
				if err != nil {
					level.Error(logger).Log("err", err.Error())
					return
				}
			}
		}
	}()

	wg.Wait()
	c.SetWriteDeadline(time.Now().Add(time.Second * 10))
	c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	time.Sleep(time.Second * 10)
	c.Close()
}
