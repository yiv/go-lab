package main

import (
	"flag"
	"fmt"
)

var (
	grpcAddr = flag.String("grpc.addr", ":8500", "GRPC server listen address")
)

func main() {
	flag.Parse()
	fmt.Println(*grpcAddr)

}
