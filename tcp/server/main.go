package main

import (
	"log"
	"net"
	"io/ioutil"
	"fmt"
	"time"
)

func main() {
	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	l, err := net.Listen("tcp", "127.0.0.1:8000")
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

func handle(conn net.Conn)  {
	conn.SetDeadline(time.Now().Add(time.Second*2))
	r ,err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println(err)
		conn.Close()
		return
	}

	fmt.Printf("%s",r)
}