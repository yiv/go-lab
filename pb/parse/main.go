package main

import (
	"fmt"
	"os"

	"github.com/emicklei/proto"
)

func main() {
	reader, _ := os.Open("test.proto")
	defer reader.Close()

	parser := proto.NewParser(reader)
	definition, _ := parser.Parse()

	proto.Walk(definition, proto.WithService(handleService), proto.WithMessage(handleMessage), proto.WithRPC(handleRPC))
}

func handleService(s *proto.Service) {
	fmt.Println(s.Name)
}
func handleRPC(m *proto.RPC) {
	fmt.Println(m.Name)
	fmt.Println(m.RequestType)
	fmt.Println(m.ReturnsType)
}

func handleMessage(m *proto.Message) {
	fmt.Println(m.Name)
	fmt.Printf("%#v", m)
	for _, v := range m.Elements {
		x := v.(*proto.NormalField)
		fmt.Println(x.Name)
	}
}
