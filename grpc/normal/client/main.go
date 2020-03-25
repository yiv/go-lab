package main

import (
	"flag"
	"time"

	log "github.com/sirupsen/logrus"
	pb "github.com/yiv/go-lab/grpc/normal/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	//"os"
	//"os/signal"
	//"syscall"
	//"time"
)

var (
	svc = flag.String("svc", "topme:9080", "The server address in the format of host:port")
)

func main() {
	flag.Parse()
	log.Println("dial ", *svc)
	conn, err := grpc.Dial(*svc, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserClient(conn)

	go func() {
		for {
			res, err := client.GetUserInfo(context.Background(), &pb.Req{Uid: time.Now().Unix()})
			if err != nil {
				log.Error("edwin #36", err.Error())
			}
			log.Println("edwin #37", res)
			time.Sleep(time.Second * 2)
		}
	}()

	errc := make(chan error)
	log.Info("terminated", <-errc)
}
