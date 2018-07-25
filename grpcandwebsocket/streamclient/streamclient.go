package main

import (
	"flag"

	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/yiv/go-lab/grpcandwebsocket/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

var (
	serverAddr = flag.String("server_addr", "192.168.1.51:7799", "The server address in the format of host:port")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewGMServiceClient(conn)
	stream, err := client.Stream(context.Background())
	if err != nil {
		log.Fatalf("%v.RouteChat(_) = _, %v", client, err)
	}
	go func() {
		for {
			err := stream.Send(&pb.Frame{Payload: toBytes(500, []byte("hello"))})
			if err != nil {
				log.Error("send err :", err)
			}

			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			f, err := stream.Recv()
			if err != nil {
				log.Error("recv err :", err)
				break
			}
			fmt.Printf("recv[code: %v, msg: %s]\n", frameCode(f), framebytes(f))

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

func toBytes(i uint32, pbBytes []byte) (payload []byte) {
	payload = append(payload, byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
	payload = append(payload, pbBytes...)
	return
}
func frameCode(f *pb.Frame) uint32 {
	payload := f.Payload
	return bytes2code(payload[0:4])
}
func framebytes(f *pb.Frame) []byte {
	payload := f.Payload
	return payload[4:]
}
func code2bytes(i uint32) (b []byte) {
	b = append(b, byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
	return
}
func bytes2code(b []byte) (i uint32) {
	i = uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
	return
}
