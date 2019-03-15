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

const (
	PoolMax          = 10
	StreamMax        = 10
	MessagePerStream = 5000000
)

var sum int32

func main() {
	multiStream()
}

func multiStream() {
	var pool []pb.GameClient
	for m := 0; m < PoolMax; m++ {
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
	for i := 0; i < StreamMax; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cli := pool[rand.Int31n(PoolMax)]
			stream, err := cli.Stream(context.Background())
			if err != nil {
				log.Fatalf("fail to open stream: %v", err)
			}
			readandwrite(stream, MessagePerStream)
			stream.CloseSend()
		}()
	}

	wg.Wait()

	log.Info("all ", StreamMax*MessagePerStream, " sum ", sum, " done took ", time.Now().Sub(now))
}

func readandwrite(stream pb.Game_StreamClient, max int) {
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
				if count > max {
					//log.Info("stream recv, ", count)
					cancel()
					count--
					return
				}
				sum++
				count++
			}
		}
	}()
	wg.Wait()
}
