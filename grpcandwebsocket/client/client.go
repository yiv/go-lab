package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

var (
	agentAddr = "192.168.1.12:10050"
)

type client struct {
	sk        *websocket.Conn
	uid       string
	token     string
	roomClass int32
	seat      int32
}

func newClient() (c *client) {
	c = &client{
		sk:  newWebSocket(),
		uid: "55852618-040c-425c-b4e9-97fb16f55730",
	}
	return
}

func (c *client) recv() {
	for {
		_, message, err := c.sk.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		code, msg := splitBytes(message)

		fmt.Printf("[收到][code = %v, msg = %s]", code, msg)

	}
}
func (c *client) send() {
	var cmdCode uint32
	var msg string
	fmt.Println("=======命令行启动=====")
	for {
		fmt.Scanln(&cmdCode)
		//fmt.Scanln(&msg)
		fmt.Printf("输入命令: %v，消息：%s\n", cmdCode, msg)
		err := c.sk.WriteMessage(websocket.BinaryMessage, toBytes(cmdCode, []byte(msg)))
		if err != nil {
			fmt.Println("命令发送错误：", err)
			continue
		}
	}

}

func newWebSocket() *websocket.Conn {
	u := url.URL{Scheme: "ws", Host: agentAddr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	fmt.Println("socket connected ")
	return c
}
func splitBytes(payload []byte) (code uint32, pbBytes []byte) {
	code = uint32(payload[0])<<24 | uint32(payload[1])<<16 | uint32(payload[2])<<8 | uint32(payload[3])
	pbBytes = payload[4:]
	return
}
func toBytes(i uint32, pbBytes []byte) (payload []byte) {
	payload = append(payload, byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
	payload = append(payload, pbBytes...)
	return
}

func main() {
	cli := newClient()
	go cli.recv()
	go cli.send()
	select {}
}
