package main

import (
	"flag"
)

var (
	Version   = "No Version Provided"
	BuildTime = "No BuildTime Provided"
)

var (
	version   = flag.String("ver", Version, "Consul agent address")
	buildTime = flag.String("build", BuildTime, "Consul agent address")
)

func main() {
	flag.Parse()
}
