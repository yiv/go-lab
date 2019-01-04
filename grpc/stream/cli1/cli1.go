package main

import (
	"context"
	"github.com/gogo/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"math/rand"
	"sync"
	"time"

	"github.com/yiv/go-lab/grpc/stream/pb"
)

func main() {
	multiStream()
}
func multiStream() {
	var pool []pb.GameClient
	for m := 0; m < 100; m++ {
		conn, err := grpc.Dial(":7788", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("fail to dial: %v", err)
		}
		cli := pb.NewGameClient(conn)
		pool = append(pool, cli)
	}

	wg := sync.WaitGroup{}

	now := time.Now()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cli := pool[rand.Int31n(100)]
			stream, err := cli.Stream(context.Background())
			if err != nil {
				log.Fatalf("fail to open stream: %v", err)
			}
			readandwrite(stream)
			stream.CloseSend()
		}()
	}

	wg.Wait()

	log.Info("all done took ", time.Now().Sub(now))
}

//
//func singleStream() {
//	conn, err := grpc.Dial(":7788", grpc.WithInsecure())
//	if err != nil {
//		log.Fatalf("fail to dial: %v", err)
//	}
//	log.Info("dial done")
//
//	cli := pb.NewGameClient(conn)
//	stream, err := cli.Stream(context.Background())
//	if err != nil {
//		log.Fatalf("fail to open stream: %v", err)
//	}
//	log.Info("stream done")
//
//	go func() {
//		log.Info("send start")
//		defer func(begin time.Time) {
//			log.Info("all done, took: ", time.Since(begin))
//		}(time.Now())
//		for i := 0; i < 100000; i++ {
//			e := stream.Send(&pb.Message{Body: fmt.Sprintf("body_%d", i)})
//			if e != nil {
//				log.Fatal("fail to send stream: %v", e)
//			}
//			log.Info("send out")
//		}
//
//	}()
//	select {}
//}

func readandwrite(stream pb.Game_StreamClient) {
	wg := sync.WaitGroup{}
	pl, _ := proto.Marshal(&pb.UpdateNetReq{Delay: 50})
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				e := stream.Send(&pb.Frame{Payload: pl})
				if e != nil {
					log.Fatal("fail to send stream: %v", e)
				}
			}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		count := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				_, e := stream.Recv()
				if e != nil {
					log.Fatal("fail to send stream: %v", e)
				}
				count++
				if count > 100 {
					log.Info("stream recv, ", count)
					cancel()
					return
				}
			}
		}
	}()
	wg.Wait()
}
