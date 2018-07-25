package main

import (
	"fmt"
	"time"

	"github.com/anvie/port-scanner"
)

func main() {

	ps := portscanner.NewPortScanner("localhost", time.Second)

	// get opened port
	fmt.Printf("scanning port %d-%d...\n", 20, 30000)

	openedPorts := ps.GetOpenedPort(53276, 53278)

	for i := 0; i < len(openedPorts); i++ {
		port := openedPorts[i]
		fmt.Print(" ", port, " [open]")
		fmt.Println("  -->  ", ps.DescribePort(port))
	}
}
