package main

import (
	"context"
	"io"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	kitetcd "github.com/go-kit/kit/sd/etcd"
	"github.com/go-kit/kit/sd/lb"
)

func main() {
	// Let's say this is a service that means to register itself.
	// First, we will set up some context.
	var (
		etcdServer = "http://192.168.1.51:2379" // don't forget schema and port!
		ctx        = context.Background()
	)

	// Build the client.
	client, err := kitetcd.NewClient(ctx, []string{etcdServer}, kitetcd.ClientOptions{})
	if err != nil {
		panic(err)
	}

	// It's likely that we'll also want to connect to other services and call
	// their methods. We can build an Instancer to listen for changes from etcd,
	// create Endpointer, wrap it with a load-balancer to pick a single
	// endpoint, and finally wrap it with a retry strategy to get something that
	// can be used as an endpoint directly.
	barPrefix := "/user/login/"
	logger := log.NewNopLogger()
	instancer, err := kitetcd.NewInstancer(client, barPrefix, logger)
	if err != nil {
		panic(err)
	}
	endpointer := sd.NewEndpointer(instancer, barFactory, logger)
	balancer := lb.NewRoundRobin(endpointer)
	retry := lb.Retry(3, 3*time.Second, balancer)

	// And now retry can be used like any other endpoint.
	req := struct{}{}
	if _, err = retry(ctx, req); err != nil {
		panic(err)
	}
}

func barFactory(string) (endpoint.Endpoint, io.Closer, error) { return endpoint.Nop, nil, nil }
