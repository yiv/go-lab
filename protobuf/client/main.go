package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "github.com/yiv/go-lab/protobuf/tutorial"
	"io/ioutil"
)

func main() {

	in, _ := ioutil.ReadFile("../book")
	newp := &pb.Person{}

	fmt.Println(in)

	_ = proto.Unmarshal(in, newp)

	fmt.Println(newp.Name)

}
