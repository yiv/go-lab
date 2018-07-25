package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/sh", "-c", "sudo docker start go_gameagent_1")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)
}
