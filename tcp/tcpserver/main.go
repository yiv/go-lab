package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	l, err := net.Listen("tcp", "127.0.0.1:4315")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	log.Println("new connection")
	for {
		msg := make([]byte, 4086)
		n, err := conn.Read(msg)
		if err != nil {
			fmt.Println(err)
			conn.Close()
			return
		}

		log.Println("receive ", n, "msg", string(msg))

		_, err = conn.Write(msg)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}

}
