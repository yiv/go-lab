package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	ketcd "github.com/go-kit/kit/sd/etcd"
	"github.com/go-kit/kit/sd/lb"

	"git.ifunbow.com/tpserver/gameserver/game"
	ucgrpccli "git.ifunbow.com/tpserver/usercenter/client"
	"git.ifunbow.com/tpserver/usercenter/handling"
)

var (
	httpAddr          = flag.String("http.addr", ":10070", "Address for HTTP (JSON) server")
	debugAddr         = flag.String("debug.addr", ":10071", "Debug and metrics listen address")
	etcdAddr          = flag.String("etcd.addr", "http://192.168.1.51:2379", "Consul agent address")
	retryMax          = flag.Int("retry.max", 3, "per-request retries to different instances")
	retryTimeout      = flag.Duration("retry.timeout", 500*time.Millisecond, "per-request timeout, including retries")
	serviceUsercenter = flag.String("server.usercenter", "usercenter", "service name for user center")
)

type userCenter struct {
}

func (u *userCenter) GetInfo(id game.UserID) (info *game.PlayerInfo, err error) {

	return
}
func (u *userCenter) AdjustCoin(id game.UserID, coins int64) (err error)         { return }
func (u *userCenter) AdjustGift(id game.UserID, gid int, amount int) (err error) { return }
func (u *userCenter) AdjustGem(id game.UserID, coins int) (err error)            { return }

func main() {
	flag.Parse()

	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	logger.Log("msg", "hello")
	defer logger.Log("msg", "goodbye")

	// Mechanical domain.
	errc := make(chan error)
	// Interrupt handler.
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// Service discovery domain
	var ctx = context.Background()
	client, err := ketcd.NewClient(ctx, []string{*etcdAddr}, ketcd.ClientOptions{})
	if err != nil {
		panic(err)
	}

	// Transport domain.
	tracer := stdopentracing.GlobalTracer() // no-op

	// server routes.
	{

		instancer, err := ketcd.NewInstancer(client, *serviceUsercenter, logger)
		if err != nil {
			panic(err)
		}
		edps := service.Endpoints{}
		{
			factory := factoryFor(service.MakeGetDeviceIDEndpoint, tracer, logger)
			endpointer := sd.NewEndpointer(instancer, factory, logger)
			balancer := lb.NewRoundRobin(endpointer)
			retry := lb.Retry(*retryMax, *retryTimeout, balancer)
			edps.GetDeviceIDEndpoint = retry
		}
		request :=
		edps.GetDeviceIDEndpoint(ctx,)
	}


	// Run!
	logger.Log("exit", <-errc)
}

func factoryFor(makeEndpoint func(service.Service) endpoint.Endpoint, tracer stdopentracing.Tracer, logger log.Logger) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		service := ucgrpccli.New(conn, tracer, logger)
		ep := makeEndpoint(service)

		return ep, conn, nil
	}
}
