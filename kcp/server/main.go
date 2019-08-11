package main

import (
	"log"
	"net"

	"github.com/xtaci/kcp-go"
)

func main() {

	if listener, err := kcp.Listen("127.0.0.1:12345"); err == nil {
		// spin-up the client
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
			}
			go handleEcho(conn)
		}
	} else {
		log.Fatal(err)
	}
}
func handleEcho(conn net.Conn) {
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}

		n, err = conn.Write(buf[:n])
		if err != nil {
			log.Println(err)
			return
		}
	}
}
