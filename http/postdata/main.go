package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var (
	host = flag.String("addr", "https://s2.tensafe.net", "http service address")
	//host   = flag.String("addr", "http://10.0.19.6", "http service address")
	path   = flag.String("path", "/api/user/getUser", "http service address")
	worker = flag.Int("worker", 1, "count of concurrency")
	delay  = flag.Int("delay", 10, "delay of new connection(Microsecond)")
	freq   = flag.Int("freq", 50, "frequency of data send(Second)")
	round  = flag.Int("round", 1, "round of data send/receive")
	start  = flag.Int64("start", 100, "the time to send a request")
)

var (
	successMtx sync.Mutex
	failMtx    sync.Mutex
	bytesMtx   sync.Mutex
	workerMtx  sync.Mutex
	allBytes   int
	success    int
	fail       int
	workerSum  int
)

func main() {
	flag.Parse()
	PostForm()
}

func updateWorkerSum(i int) {
	workerMtx.Lock()
	workerSum += i
	workerMtx.Unlock()
}

func updateBytes(i int) {
	//bytesMtx.Lock()
	allBytes += i
	//bytesMtx.Unlock()
}

func updateSuccess(i int) {
	successMtx.Lock()
	success += i
	successMtx.Unlock()
}

func updateFail(i int) {
	failMtx.Lock()
	fail += i
	failMtx.Unlock()
}

func PostForm() {
	wg := sync.WaitGroup{}
	startCh := make(chan struct{})
	endCh := make(chan struct{})
	for i := 0; i < *worker; i++ {
		wg.Add(1)
		go func(id int, startCh, endCh chan struct{}) {

			updateWorkerSum(1)

			var (
				successSub = 0
				failSub    = 0
				bytesSub   = 0
			)

			defer func() {
				wg.Done()
				updateBytes(bytesSub)
				updateSuccess(successSub)
				updateFail(failSub)
			}()

			log.Println("worker", workerSum)
			<-startCh
			for x := 1; x < *round+1; x++ {
				select {
				case <-endCh:
					break
				default:
					form := url.Values{}
					form.Add("user_id", fmt.Sprintf("%v", 83905686261604352))
					st := time.Now()
					//res, err := http.PostForm("https://s2.tensafe.net/api/user/getUser", form)
					u := *host + *path
					res, err := http.PostForm(u, form)
					code := 0
					if err != nil {
						log.Error(err.Error())
						updateFail(1)
					} else {
						code = res.StatusCode
						if res.StatusCode == 200 {
							body, _ := ioutil.ReadAll(res.Body)
							bytesSub += len(body)
							successSub += 1
						} else {
							failSub += 1
						}
					}
					if x%20 == 0 {
						log.Printf("id: %v, round: %v, code: %v, time used: %v", id, x, code, time.Now().Sub(st))
					}
					if *freq > 0 {
						time.Sleep(time.Microsecond * time.Duration(*freq))
					}
				}

			}
		}(i, startCh, endCh)
		if *delay > 0 {
			time.Sleep(time.Microsecond * time.Duration(*delay))
		}
	}

	requestTime := time.Unix(*start, 0)
	log.Println("time", "up", "after", requestTime.Sub(time.Now()))

	startTimer := time.NewTimer(requestTime.Sub(time.Now()))
	var startPost time.Time
	go func() {
		<-startTimer.C
		close(startCh)
		startPost = time.Now()
	}()

	wg.Wait()
	log.Println("success:", success, "fail:", fail, "all time:", time.Now().Sub(startPost), "all bytes", allBytes, "throughput", *worker**round)
}
