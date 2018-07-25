package main

import (
	"flag"
	"fmt"
	"github.com/yiv/go-lab/grpcandwebsocket/pb"
	"google.golang.org/grpc"
	"io"
	"net"
)

var (
	port = flag.Int("port", 7799, "The server port")
)

type gameServer struct {
}

func (s *gameServer) Stream(stream pb.GMService_StreamServer) error {
	fmt.Println("stream start to revc .......")
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("err", err.Error(), "msg", "stream EOF err")
				return
			}
			if err != nil {
				fmt.Println("err", err.Error(), "msg", "stream receive err")
				return
			}
			code := frameCode(in)
			bytes := framebytes(in)
			fmt.Printf("[收到][code = %v, msg = %s]\n", code, bytes)
		}
	}()
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("192.168.1.51:%d", *port))
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
	}
	fmt.Printf("listen on : %v\n", *port)
	grpcServer := grpc.NewServer()
	grpcHandler := new(gameServer)
	pb.RegisterGMServiceServer(grpcServer, grpcHandler)
	fmt.Printf("grpc running.....\n")
	grpcServer.Serve(lis)
	fmt.Printf("server end\n")
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
