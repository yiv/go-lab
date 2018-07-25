package main

import (
	"context"
	"io"
	"time"

	//"github.com/gogo/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"../pb"
)

func main() {
	multiStream()
}
func multiStream() {
	conn, err := grpc.Dial(":7788", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	log.Info("dial done")

	cli := pb.NewGameClient(conn)
	stream, err := cli.Stream(context.Background())
	if err != nil {
		log.Fatalf("fail to open stream: %v", err)
	}
	dieChan := make(chan struct{})
	go func() {
		for {
			_, err := stream.Recv()
			if err == io.EOF { // 流关闭
				log.Debug(err)
				return
			}
			if err != nil {
				log.Error(err)
				return
			}
			select {
			case <-dieChan:
				log.Info("die chan recv")
				return
			}
		}

	}()
	time.Sleep(time.Second * 5)
	stream.CloseSend()
	select {}
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

//func readandwrite() {
//	conn, err := grpc.Dial(":7788", grpc.WithInsecure())
//	if err != nil {
//		log.Fatalf("fail to dial: %v", err)
//	}
//	log.Info("dial done")
//
//	cli := pb.NewGameServiceClient(conn)
//	stream, err := cli.Stream(context.Background())
//	bt, err := proto.Marshal(&pb.GeCall{Seat: 5, CallSeat: 4})
//	go func() {
//		for {
//			e := stream.Send(&pb.Frame{Payload: bt})
//			if e != nil {
//				log.Fatal("fail to send stream: %v", e)
//			}
//			log.Info("stream send, ")
//			time.Sleep(1 * time.Second)
//		}
//
//	}()
//	go func() {
//		for {
//			f, e := stream.Recv()
//			if e != nil {
//				log.Fatal("fail to send stream: %v", e)
//			}
//			log.Info("stream recv, ", f)
//		}
//
//	}()
//	select {}
//}
