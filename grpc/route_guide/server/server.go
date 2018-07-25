package main

import (
	"flag"
	"fmt"
	"io"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "github.com/yiv/go-lab/grpc/route_guide/routeguide"
)

var (
	port = flag.Int("port", 11000, "The server port")
)

type routeGuideServer struct {
}

var usermap map[int64]pb.RouteGuide_RouteChatServer

func (s *routeGuideServer) RouteChat(stream pb.RouteGuide_RouteChatServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	grpcHandler := new(routeGuideServer)
	pb.RegisterRouteGuideServer(grpcServer, grpcHandler)
	grpcServer.Serve(lis)
}
