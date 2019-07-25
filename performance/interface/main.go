package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	data := &Data{}
	var sum int
	begin := time.Now()
	wg := &sync.WaitGroup{}

	for x := 0; x < 5; x++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100000000; i++ {
				data.X = i
				sum += ByStrict(data)
			}
		}()
	}
	wg.Wait()
	fmt.Println("time consume: ", time.Since(begin).Seconds())

	//errc := make(chan error)
	//go func() {
	//	logger := log.NewLogfmtLogger(os.Stdout)
	//	m := http.NewServeMux()
	//	m.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	//	m.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	//	m.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	//	m.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	//	m.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	//	m.Handle("/metrics", promhttp.Handler())
	//
	//	logger.Log("addr", *debugAddr)
	//	errc <- http.ListenAndServe(*debugAddr, m)
	//}()

}

type Data struct {
	X int
}

func ByStrict(d *Data) (r int) {

	return d.X
}

func ByAssert(d interface{}) int {
	data := d.(*Data)
	return data.X
}
