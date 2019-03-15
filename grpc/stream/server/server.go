package main

import (
	"context"
	"net"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/yiv/go-lab/grpc/stream/pb"
)

type server struct {
}

func (s *server) Stream(stream pb.Game_StreamServer) (err error) {
	log.Info("new stream")
	defer func(begin time.Time) {
		log.Info("all done, took: ", time.Since(begin))
	}(time.Now())
	buf := make(chan *pb.Frame, 5000)
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				f, e := stream.Recv()
				if e != nil {
					err = e
					cancel()
					log.Error("err on recv stream: ", e)
					return
				}
				buf <- f
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case f := <-buf:
				stream.Send(f)
			}
		}
	}()

	wg.Wait()
	return
}
func main() {
	lis, err := net.Listen("tcp", ":7788")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.MaxConcurrentStreams(50000))
	grpcHandler := new(server)
	pb.RegisterGameServer(grpcServer, grpcHandler)
	grpcServer.Serve(lis)
}
