package main

import (
	"github.com/xtaci/kcp-go"
	"io"
	"log"
	"time"
)

func main() {
	// dial to the echo server
	if sess, err := kcp.Dial("127.0.0.1:12345"); err == nil {
		for {
			data := time.Now().String()
			buf := make([]byte, len(data))
			log.Println("sent:", data)
			if _, err := sess.Write([]byte(data)); err == nil {
				// read back the data
				if _, err := io.ReadFull(sess, buf); err == nil {
					log.Println("recv:", string(buf))
				} else {
					log.Fatal(err)
				}
			} else {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	} else {
		log.Fatal(err)
	}
}
