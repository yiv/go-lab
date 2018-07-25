package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strings"
	"bytes"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

type Query struct {
	Query string `csv:"query"`
	Uid   string `csv:"uid"`
}

type RInfo struct {
	Query   string `csv:"query"`
	Uid     string `csv:"uid"`
	Country string `csv:"country"`
	City    string `csv:"city"`
}

func main() {
	var querys []Query
	var fd *os.File
	var err error
	fd, err = os.OpenFile("./in.csv", os.O_RDWR, 0755)
	if err != nil {
		fmt.Println("edwin #1", err)
		return
	}


	if err = gocsv.UnmarshalFile(fd, &querys); err != nil { // Load clients from file
		fmt.Println("edwin #2", err)
		return
	}
	var queryArr []Query
	var RInfoArr []RInfo
	infoMap := make(map[string]RInfo)
	IP2UidMap := make(map[string]string)

	for _, q := range querys {
		q.Uid = string(bytes.TrimSpace([]byte(q.Uid)))
		q.Query = string(bytes.TrimSpace([]byte(q.Query)))
		if q.Uid != "" {
			IP2UidMap[q.Query] = q.Uid
			queryArr = append(queryArr, q)
		}
	}

	fmt.Println("edwin #73 ", IP2UidMap)

	urlstr := "http://ip-api.com/batch?fields=query,country,city"

	start := 0
	cli := http.Client{}
	for {
		if start >= len(queryArr) {
			break
		}
		end := start + 100
		var poststr []byte
		if end > len(queryArr) {
			poststr, _ = json.Marshal(queryArr[start:])
		} else {
			poststr, _ = json.Marshal(queryArr[start:end])
		}
		res := query(cli, urlstr, string(poststr))
		for _, r := range res {
			r.Uid = IP2UidMap[r.Query]
			infoMap[r.Query] = r
			fmt.Println("edwin #5555 ", r, r.Query)
		}
		start = end
		time.Sleep(time.Millisecond * 500)
		fmt.Fprintln(os.Stderr, start)
	}



	for _, v := range infoMap {
		fmt.Printf("infoMap \"%s\",\"%s\",\"%s\",\"%s\"\n", v.Query, v.Uid, v.Country, v.City)
		RInfoArr = append(RInfoArr, v)
	}

	var csvContent string
	csvContent, err = gocsv.MarshalString(&RInfoArr)

	fd, err = os.OpenFile("./out.csv", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("edwin #3", err)
		return
	}
	var n int
	if n, err = fd.Write([]byte(csvContent)); err != nil {
		fmt.Println("edwin #4", err)
		return
	} else {
		fmt.Println(csvContent)
		fmt.Println("success", n)

	}

}
func query(cli http.Client, urlstr, query string) []RInfo {
	var err error
	var resp *http.Response
	resp, err = cli.Post(urlstr, "", strings.NewReader(string(query)))
	if err != nil {
		fmt.Println("edwin #5", err.Error())
		return nil
	}

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("edwin #6", err.Error())
		return nil
	}else{
		fmt.Println("edwin #66", string(body))
	}

	var res []RInfo
	if err = json.Unmarshal(body, &res); err != nil {
		fmt.Println(err.Error())
	}
	//for _, q := range res {
	//	fmt.Printf("ip,%s,country,%s,city,%s\n", q.Query, q.Country, q.City)
	//}
	return res
}
