package main

import (
	"context"
	"fmt"
	"github.com/prometheus/common/log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"10.72.17.30:12379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer cli.Close()

	members, _ := cli.MemberList(context.Background())
	fmt.Println(members)

}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
