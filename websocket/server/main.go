package main

import (
	"context"
	"flag"
	"fmt"
	//"io/ioutil"
	//"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	//"time"

	ws "github.com/gorilla/websocket"
	//"zycase.cn/shaoyong/wechat/agent/websocket"
)

var (
	webSocketAddr = flag.String("websocket.addr", ":10050", "game agent webSocket address")
)

type Message struct {
	MType   int
	Content []byte
}

func main() {
	flag.Parse()
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
	fmt.Println("terminated", <-errc)
}

func webSocketServerStd(w http.ResponseWriter, r *http.Request) {
	fmt.Println("webSocketServerStd new client")
	var upgrader = ws.Upgrader{} // use default options
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print("upgrade:", err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	buf := make(chan Message, 10000)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				mt, message, err := c.ReadMessage()
				if err != nil {
					fmt.Println("read:", err)
					cancel()
					return
				}
				buf <- Message{MType: mt, Content: message}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-buf:
				err = c.WriteMessage(msg.MType, msg.Content)
				if err != nil {
					fmt.Println("write:", err)
					return
				}
			}
		}
	}()
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
