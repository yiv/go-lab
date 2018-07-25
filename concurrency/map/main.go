package main

import (
	"sync"
	"github.com/yiv/go-lab/concurrency/map/pb"
	"github.com/golang/protobuf/proto"
	"time"
	"fmt"
)

type User struct {
	mtx     sync.RWMutex
	Records map[string]int64 `json:"records"`
}

func (u *User) copy() *User {
	fmt.Println("User", "copy")
	u.mtx.Lock()
	defer u.mtx.Unlock()
	c := &User{
		Records: u.Records,
	}
	return c
}

func main() {
	user := &User{Records: make(map[string]int64)}
	user.Records["win"] = 500
	user.Records["lost"] = 5000
	Marshal(user.copy())
	//for x := 0; x < 100; x++ {
	//	for y := 0; y < 100; y++ {
	//		Marshal(user.copy())
	//	}
		time.Sleep(time.Second * 0)
	//}
}

func Marshal(user *User) {
	fmt.Println("Marshal", "Marshal")
	u := &pb.User{Records: user.Records}
	proto.Marshal(u)
}
