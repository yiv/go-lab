/* GetHeadInfo
 */
package main

import (
	"net"
	"os"
	"fmt"
	//"io/ioutil"
	"time"
)

func main() {


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

