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
	webSocketAddr = flag.String("websocket.addr", ":999", "game agent webSocket address")
	sum           = 0
)
var logger log.Logger

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
		m := http.NewServeMux()
		m.HandleFunc("/ws", webSocketServerStd)
		errc <- http.ListenAndServe(*webSocketAddr, m)
	}()
	level.Info(logger).Log("end", <-errc)
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

//func webSocketServer(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("webSocketServer new client")
//	conn, err := websocket.NewWebsocketServerConn(w, r)
//	if err != nil {
//		return
//	}
//	//conn.SetReadDeadline(time.Now().Add(15 * time.Second))
//	for {
//		//bytes, err := ioutil.ReadAll(conn)
//
//		buf := make([]byte, 20)
//
//		n, err := conn.Read(buf)
//		if err != nil {
//			fmt.Println("webSocket read err ", err)
//			return
//		}
//		fmt.Printf("webSocket read  %d byte, buf = %v", n, buf)
//		conn.Write([]byte("hello "))
//	}
//
//}
