package main

import (
	"fmt"
	"github.com/gmallard/stompngo"
	"net"
	"time"
)

var (
	host       = "192.168.10.166"       // default host
	port       = "61613"                // default port
	protocol   = "1.0"                  // Default protocol level
	login      = "guest"                // default login
	passcode   = "guest"                // default passcode
	vhost      = "localhost"            // default vhost
	heartbeats = "0,0"                  // default (no) heartbeats
	dest       = "/queue/shaoyong-test" // default destination
	scc        = 1                      // Subchannel capacity
	nmsgs      = 1                      // default number of messages (useful at times)
	maxbl      = -1                     // Max body length to dump (-1 => no limit)
	// ackMode    = "auto"
	ackMode = "client"
	// ackMode    = "client-individual"
)

func main() {
	conn, err := net.Dial(stompngo.NetProtoTCP, net.JoinHostPort(host, port))
	if err != nil {
		fmt.Println(err)
	}

	stompConn, err := stompngo.Connect(conn, stompngo.Headers{"accept-version", protocol, "host", vhost, "heart-beat", heartbeats, "id", stompngo.Uuid()})
	fmt.Printf("%#v\n", stompConn)
	id := stompngo.Uuid()
	mdc, err := stompConn.Subscribe(stompngo.Headers{"destination", dest, "ack", ackMode, "id", id})
	var md stompngo.MessageData
	start := time.Now()
	for {
		md = <-mdc
		_ = md
		fmt.Print(".")
		// fmt.Println(string(md.Message.Body))
		// fmt.Println(id)
		fmt.Printf("%#v\n", md.Message.Headers)
		// fmt.Println(md.Message.Headers.Value("message-id"))
		err = stompConn.Ack(stompngo.Headers{"message-id", md.Message.Headers.Value("message-id")})
		// err = stompConn.Ack(stompngo.Headers{"message-id", md.Message.Headers.Value("message-id"), "subscription", id})
		// err = stompConn.Ack(stompngo.Headers{"id", md.Message.Headers.Value("ack")})
		if err != nil {
			fmt.Println(err)
		}
		// time.Sleep(time.Microsecond * 500)
		// break

	}
	fmt.Println(time.Now().Sub(start).Seconds())

}
