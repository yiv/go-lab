package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format("2006-01-02 15:04:05 +0000 UTC"))
	fmt.Println(time.Now().Format("20060102"))
	fmt.Println(time.Now().Unix())

	fmt.Println(time.Now().UnixNano())

	timestr := strconv.FormatInt(time.Now().UnixNano(), 10)
	fmt.Println(timestr)
	md5val := md5.Sum([]byte("1497410603940131300"))
	md5str := fmt.Sprintf("%x", md5val)
	fmt.Println(md5str)
	fmt.Println(("909DC39ADD7F1E11A6CB2BACFC39E552"))

	fmt.Println(time.Now().Format("xxx"))

	fmt.Println("unix sec to time : ", time.Unix(1483228800, 0))

	t, err := time.Parse("2006-01-02 15:04:05", "2017-08-22 15:03:05")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Parse string to time", t)
	}

	//取两天的时间差，天数
	d1, _ := time.Parse("2006-01-02 15:04:05", "2017-08-22 15:03:05")
	d2, _ := time.Parse("2006-01-02 15:04:05", "2017-08-25 15:00:05")
	fmt.Println("sub time : ", int64(d2.Sub(d1).Hours())/24)

	//duration 转 seconds
	fmt.Printf("duration 转 seconds : %v\n", int64(time.Duration(time.Hour*24*30).Seconds()))

	//取相差的天数
	ctime := int64(1503880003)
	c, _ := time.Parse("2006-01-02", time.Unix(ctime, 0).Format("2006-01-02"))
	n, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	fmt.Println(c)
	fmt.Println(n)
	fmt.Printf("%d\n", int64(n.Sub(c).Hours())/24)

}
