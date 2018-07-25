package main

import (
	"flag"
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/yiv/go-lab/grpc/normal/pb"
)

var (
	port = flag.Int("port", 11000, "The server port")
)

type userServer struct {
}

func (u userServer) GetUserInfo(context context.Context, req *pb.Req) (res *pb.Res, err error) {
	res = &pb.Res{Nick: "edwin"}
	return
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	grpcHandler := new(userServer)
	pb.RegisterUserServer(grpcServer, grpcHandler)
	grpcServer.Serve(lis)
}
