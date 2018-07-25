package main

import (
	"net"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"../pb"
)

type server struct {
}

func (s *server) Stream(stream pb.Game_StreamServer) error {
	log.Info("new stream")
	defer func(begin time.Time) {
		log.Info("all done, took: ", time.Since(begin))
	}(time.Now())
	count := 0
	for {
		f, e := stream.Recv()
		log.Info("stream recv")
		if e != nil {
			log.Error("err on recv stream: ", e)
			return e
		}
		stream.Send(f)
		log.Info("send out")
		count++
		if count >= 100000 {
			break
		}
	}
	return nil
}
func main() {
	lis, err := net.Listen("tcp", ":7788")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	grpcHandler := new(server)
	pb.RegisterGameServer(grpcServer, grpcHandler)
	grpcServer.Serve(lis)
}
