package main

import (
	"fmt"
	"time"
)

func main() {
	//TimePrint()
	for {
		time.Sleep(time.Second * 1)
		fmt.Printf("\r %v", time.Now().UnixNano())
	}
}

func StringToTime() {
	t, err := time.Parse("2006-01-02 15:04:05", "2017-08-22 15:03:05")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Parse string to time", t)
	}
	
}

func UnixSecondToTime() {
	fmt.Println("unix sec to time : ", time.Unix(1483228800, 0))
}

func TimeToUnixSecond() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
}

func TimePrint() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format("2006-01-02 15:04:05 +0000 UTC"))
	fmt.Println(time.Now().Format("20060102"))
}
func DurConvert() {
	//duration 转 seconds
	fmt.Printf("duration 转 seconds : %v\n", int64(time.Duration(time.Hour * 24 * 30).Seconds()))
}

func TimeSub() {
	//取两天的时间差，天数
	d1, _ := time.Parse("2006-01-02 15:04:05", "2017-08-22 15:03:05")
	d2, _ := time.Parse("2006-01-02 15:04:05", "2017-08-25 15:00:05")
	fmt.Println("sub time : ", int64(d2.Sub(d1).Hours())/24)
}
