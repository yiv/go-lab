package main

import (
	"./pb"
	"encoding/json"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"time"
)

type Account struct {
	Uid         int64
	Unionid     string
	Uuid        string
	Username    string
	Password    string
	Nick        string
	Gender      bool
	Addr        string
	Avatar      string
	Isguest     bool
	Condays     int32
	Signdate    int64
	Vipsigndate int64
	Status      bool
	Mtime       int64
	Ctime       int64
	Token       string
	Bankpwd     string
	Forbid      string
	Imsi        string
	Imei        string
	Mac         string
	Did         string
	Psystem     string
	Pmodel      string
}

const times int = 5000

func main() {
	info := pb.UserInfo{
		Uid:         888888,
		Unionid:     "xxxxxxxxxxxxxxxxxxx",
		Uuid:        "xxxxxxxxxxxxxxxxxxx",
		Username:    "xxxxxxxxxxxxxxxxxxx",
		Password:    "xxxxxxxxxxxxxxxxxxx",
		Nick:        "xxxxxxxxxxxxxxxxxxx",
		Gender:      false,
		Addr:        "xxxxxxxxxxxxxxxxxxx",
		Avatar:      "xxxxxxxxxxxxxxxxxxx",
		Isguest:     true,
		Condays:     5,
		Signdate:    5655555,
		Vipsigndate: 5655555,
		Status:      true,
		Mtime:       889855,
		Ctime:       889855,
		Token:       "xxxxxxxxxxxxxxxxxxx",
		Bankpwd:     "xxxxxxxxxxxxxxxxxxx",
		Forbid:      "xxxxxxxxxxxxxxxxxxx",
		Imsi:        "xxxxxxxxxxxxxxxxxxx",
		Imei:        "xxxxxxxxxxxxxxxxxxx",
		Mac:         "xxxxxxxxxxxxxxxxxxx",
		Did:         "xxxxxxxxxxxxxxxxxxx",
		Psystem:     "xxxxxxxxxxxxxxxxxxx",
		Pmodel:      "xxxxxxxxxxxxxxxxxxx",
	}

	b, err := json.Marshal(info)
	fmt.Printf("%s\n", b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("len of data to josn: ", len(b))

	p, err := proto.Marshal(&info)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("len of data to pb: ", len(p))

	pbf(info)
	jsonf(info)

}

func pbf(account pb.UserInfo) {
	defer func(begin time.Time) {
		fmt.Println("pb took: ", time.Since(begin))
	}(time.Now())
	for i := 0; i < times; i++ {
		_, err := proto.Marshal(&account)
		if err != nil {
			fmt.Println(err)
		}
	}
}
func jsonf(account pb.UserInfo) {
	defer func(begin time.Time) {
		fmt.Println("json took: ", time.Since(begin))
	}(time.Now())
	for i := 0; i < times; i++ {
		_, err := json.Marshal(account)
		if err != nil {
			fmt.Println(err)
		}
	}

}
