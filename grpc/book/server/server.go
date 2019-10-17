package main

import (
	"avg_service/user/pkg/user"
	"avg_service/x/mysql"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/yiv/go-lab/grpc/book/editorpb"
)

var (
	port      = flag.Int("port", 11000, "The server port")
	debugAddr = flag.String("debug.addr", ":11100", "Debug and metrics listen address")
	runTime   float64
)

type UserServer struct {
	Logger log.Logger
}

func NewUserServer(logger log.Logger) UserServer {
	us := UserServer{Logger: logger}
	return us
}
func (u UserServer) Characters(context context.Context, req *pb.CharactersReq) (res *pb.CharactersRes, err error) {
	
	return
}


func main() {
	flag.Parse()

	var logger log.Logger
	{
		logLevel := level.AllowInfo()
		logLevel = level.AllowDebug()
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = level.NewFilter(logger, logLevel)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		level.Error(logger).Log("err", err.Error())
		os.Exit(0)
	}

	errc := make(chan error)
	go func() {
		grpcServer := grpc.NewServer()
		grpcHandler := NewUserServer(logger)
		pb.RegisterEditorServer(grpcServer, grpcHandler)
		errc <- grpcServer.Serve(lis)
	}()

	

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("exit", <-errc)
	level.Info(logger).Log("server-mysql", runTime)
}
