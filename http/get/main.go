package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"crypto/tls"
	"sync"
	"time"
	"flag"
)

var (
	logger log.Logger
	//配置文件
	host = flag.String("host", "https://api.qmovies.tv:8081/room/list2", "the quest url")
	num = flag.Int("num", 1000, "the number of concurrency request")
)
func main() {
	flag.Parse()
	logger = log.NewLogfmtLogger(os.Stdout)
	TestGet()
}

func SetHeader() {
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://205.147.109.119:18181/user_monitor/edit_user/0/74264742", nil)
	if err != nil {
		level.Error(logger).Log("err", err.Error())
		return
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "PHPSESSID=mivfh11ovtt8bujld35rc1fpe2; ci_session=a%3A4%3A%7Bs%3A10%3A%22session_id%22%3Bs%3A32%3A%229c9f90e234b50e20ebf85991d1d64633%22%3Bs%3A10%3A%22ip_address%22%3Bs%3A12%3A%2245.32.14.148%22%3Bs%3A10%3A%22user_agent%22%3Bs%3A114%3A%22Mozilla%2F5.0+%28Windows+NT+10.0%3B+Win64%3B+x64%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Chrome%2F62.0.3202.89+Safari%2F537.36%22%3Bs%3A13%3A%22last_activity%22%3Bi%3A1525257327%3B%7D864c033d78a04ab93c67bcc9731edd73")
	req.Header.Add("Host", "205.147.109.119:18181")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.89 Safari/537.36")

	if resp, err := client.Do(req); err != nil {
		level.Error(logger).Log("err", err.Error())
		return
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		level.Info(logger).Log("body", string(body))
	}
}

func GetYouStarRoomList() {
	get := func() (code int, err error) {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		//resp, err := http.Get("https://api.qmovies.tv:8081/room/list2")
		resp, err := http.Get(*host)
		if err != nil {

			return 0, err
		}
		return resp.StatusCode, nil
	}

	wg := sync.WaitGroup{}
	start := time.Now()
	for i := 0; i < *num; i++ {
		wg.Add(1)
		go func() {
			code, err := get()
			if err != nil {
				level.Error(logger).Log("err", err.Error())
			}
			if code != 200 {
				level.Error(logger).Log("code", code)
			}
			wg.Done()
		}()

	}
	wg.Wait()
	level.Info(logger).Log("Done:", time.Now().Sub(start))
}

func TestGet()  {
	
	codes := map[int]int{}
	seconds := map[int64]int{}
	st := time.Now()
	for i := 0; i < 1000; i++{
		start := time.Now()
		resp, err := http.Get("https://xms-dev-1251001060.file.myqcloud.com/2019/8/1565964059148/book_main.json")
		if err != nil {
			_= level.Error(logger).Log("err", err.Error())
		}else{
			codes[resp.StatusCode]++
		}
		used := time.Now().Sub(start).Nanoseconds() / 1000000
		
		seconds[used]++
	}
	
	fmt.Printf("%#v \n", codes)
	fmt.Printf("%#v \n", seconds)
	fmt.Println("total time used: ",time.Now().Sub(st))
}
