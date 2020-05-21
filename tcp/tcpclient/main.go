/* GetHeadInfo
 */
package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
	"sync"

	//"io/ioutil"
	"time"
)

var (
	addr  = flag.String("addr", "127.0.0.1:4315", "http service address")
	count = flag.Int("count", 1, "count of concurrency")
	delay = flag.Int("delay", 10, "delay of new connection(Microsecond)")
	freq  = flag.Int("freq", 10, "frequency of data send(Second)")
	round = flag.Int("round", 1000000, "round of data send/receive")
)

var (
	sum int
	mtx sync.Mutex
)

func main() {
	flag.Parse()
	wg := sync.WaitGroup{}

	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			log.Println("sum", sum)
			time.Sleep(time.Second * 3)
		}
	}()

	for i := 0; i < *count; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			echo(id)
		}(i)
		time.Sleep(time.Microsecond * time.Duration(*delay))
	}

	wg.Wait()
}

func updateSum(i int) {
	mtx.Lock()
	sum += i
	mtx.Unlock()
}

func echo(id int) {
	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		log.Error(err.Error())
		return
	}

	updateSum(1)
	defer func() {
		updateSum(-1)
	}()

	var (
		bytes int
		st    = time.Now()
	)

	log.Println("open", id, "sum ", sum)
	send := []byte("abcdefghijklmnopqrstuvwxyz123456abcdefghijklmnopqrstuvwxyz123456")
	for i := 0; i < *round; i++ {
		_, err = conn.Write(send)
		if err != nil {
			log.Error(err.Error())
			return
		}

		//log.Println("send ", string(msg))
		receive := make([]byte, 64)
		_, err := conn.Read(receive)
		//msg, err = ioutil.ReadAll(conn)
		if err != nil {
			log.Error(err.Error())
			return
		}

		//log.Println("receive", n, "msg", string(msg))

		bytes += len(receive)
		if *freq > 0 {
			time.Sleep(time.Second * time.Duration(*freq))
		}

	}

	log.Println("id", id, "round", *round, "bytes", bytes, "time", time.Now().Sub(st))
}

func test() {

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(err)
	req := "GET /QOS_device_info.htm?ts=1507182875903 HTTP/1.1\r\nHost: 10.0.0.1\r\nAuthorization: Basic YWRtaW46TEpMbnhhcjk4OA==\r\n\r\n"
	req = "hello"
	fmt.Println("req len ", len(req))

	for {
		fmt.Println(time.Now().Unix())

		_, err = conn.Write([]byte(req))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		//result, err := readFully(conn)
		//result, err := ioutil.ReadAll(conn)
		//checkError(err)
		//
		//fmt.Println(string(result))
		time.Sleep(time.Second)
	}

}
