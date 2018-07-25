package main

import (
	"flag"
	"io"

	log "github.com/sirupsen/logrus"
	pb "github.com/yiv/go-lab/grpc/route_guide/routeguide"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	//"os"
	//"os/signal"
	//"syscall"
	"time"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:11000", "The server address in the format of host:port")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRouteGuideClient(conn)
	stream, err := client.RouteChat(context.Background())
	if err != nil {
		log.Fatalf("%v.RouteChat(_) = _, %v", client, err)
	}
	go func() {
		for {
			err := stream.Send(&pb.RouteNote{Message: "i am client 2"})
			if err != nil {
				log.Error(err)
			}
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				log.Fatal("io.EOF")
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Info("server: ", in.Message)
			time.Sleep(time.Second * 3)
		}
	}()

	//go func() {
	//	c := make(chan os.Signal)
	//	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	//	log.Fatal("client over")
	//}()
	errc := make(chan error)
	log.Info("terminated", <-errc)
}
