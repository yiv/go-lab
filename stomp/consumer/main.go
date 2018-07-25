package main

import (
	"fmt"
	"github.com/go-stomp/stomp"
	"time"
)

func main() {
	// conn, err := stomp.Dial("tcp", "192.168.223.128:61613")
	fmt.Println("kk")
	conn, err := stomp.Dial("tcp", "192.168.10.166:61613", stomp.ConnOpt.HeartBeat(time.Hour*48, time.Hour*48))

	if err != nil {
		fmt.Println(err)
	}
	sub, err := conn.Subscribe("/queue/shaoyong-test", stomp.AckClient)
	// sub, err := conn.Subscribe("/queue/sys_event_recr", stomp.AckClient)

	if err != nil {
		fmt.Println(conn)
	}

	// receive 5 messages and then quit
	for i := 0; i < 5000000000; i++ {
		fmt.Println("#1")
		msg := <-sub.C
		fmt.Println("#2")
		if msg.Err != nil {
			fmt.Println("##", msg.Err)
			return
		}
		fmt.Println("#3")

		// acknowledge the message
		err = conn.Ack(msg)
		fmt.Println("#4")
		if err != nil {
			fmt.Println(msg)
		}
		fmt.Println("#5")
	}
	time.Sleep(time.Second * 10000)
	fmt.Printf("edwin done\n")

}

/**

func main() {
	// conn, err := stomp.Dial("tcp", "192.168.223.128:61613")
	conn, err := stomp.Dial("tcp", "192.168.10.166:61613")
	if err != nil {
		fmt.Println(err)
	}
	sub, err := conn.Subscribe("/queue/shaoyong-test", stomp.AckClient)
	if err != nil {
		fmt.Println(conn)
	}

	// receive 5 messages and then quit
	for {
		msg := <-sub.C
		if msg.Err != nil {
			fmt.Println("1##", msg.Err)
			return
		}

		fmt.Println(string(msg.Body))

		// acknowledge the message
		err = conn.Ack(msg)
		if err != nil {
			fmt.Println(msg)
		}
		time.Sleep(time.Microsecond * 10)
	}
	//fmt.Println(conn)
}

**/
