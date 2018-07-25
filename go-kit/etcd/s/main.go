package main

import (
	"context"
	"fmt"
	"os/signal"

	"github.com/go-kit/kit/log"
	kitetcd "github.com/go-kit/kit/sd/etcd"
	"os"
	"syscall"
	//"time"
)

func main() {
	// Let's say this is a service that means to register itself.
	// First, we will set up some context.
	var (
		etcdServer = "http://192.168.1.51:2379" // don't forget schema and port!
		prefix     = "/user/login/"             // known at compile time
		instance   = "192.168.1.51:8080"        // taken from runtime or platform, somehow
		key        = prefix + instance          // should be globally unique
		value      = "http://" + instance       // based on our transport
		ctx        = context.Background()
	)

	// Build the client.
	client, err := kitetcd.NewClient(ctx, []string{etcdServer}, kitetcd.ClientOptions{})
	if err != nil {
		panic(err)
	}

	// Build the registrar.
	registrar := kitetcd.NewRegistrar(client, kitetcd.Service{
		Key:   key,
		Value: value,
	}, log.NewNopLogger())

	// Register our instance.
	registrar.Register()

	// At the end of our service lifecycle, for example at the end of func main,
	// we should make sure to deregister ourselves. This is important! Don't
	// accidentally skip this step by invoking a log.Fatal or os.Exit in the
	// interim, which bypasses the defer stack.
	defer registrar.Deregister()

	// Interrupt handler.
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	//select {
	//case <-time.After(10 * time.Second):
	//	return
	//}
	//time.After(10 * time.Second)

	fmt.Println("exit", <-errc)
}
