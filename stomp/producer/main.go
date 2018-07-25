package main

import (
	"fmt"
	"github.com/go-stomp/stomp"
	"time"
)

func main() {
	conn, err := stomp.Dial("tcp", "192.168.10.166:61613")
	if err != nil {
		println("cannot connect to server", err.Error())
		return
	}
	for {
		text := fmt.Sprintf("Test message #%s#", time.Now().Format("2006-01-02 15:04:05.999999999"))
		err = conn.Send("/queue/shaoyong-test", "text/plain",
			[]byte(text), stomp.SendOpt.Header("persistent", "true"))
		if err != nil {
			println("failed to send to server", err)
			return
		}
	}
}

// func main() {
// 	conn, err := stomp.Dial("tcp", "192.168.10.166:61613")
// 	if err != nil {
// 		fmt.Println("1#", err)
// 	}
// 	for {
// 		err = conn.Send(
// 			"/queue/shaoyong-test", // destination
// 			"text/plain",           // content-type
// 			[]byte(fmt.Sprintf("Test message #%s#", time.Now().Format("2006-01-02 15:04:05.999999999")))) // body
// 		if err != nil {
// 			fmt.Println("2#", err)
// 		}
// 		time.Sleep(time.Microsecond * 1)
// 	}

// }
