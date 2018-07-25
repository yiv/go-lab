package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "github.com/yiv/go-lab/protobuf/pb"
	"io/ioutil"
)

func main() {
	player := &pb.Player{
		Id:       550000,
		SeatCode: 90000001,
		Blind:    true,
		Seat:     1,
		Coin:     20000000,
		Nick:     "edwin",
		Avatar:   "/p/a.jpg",
		Cards:    []byte{5, 2, 3},
	}

	pbBytes, _ := proto.Marshal(player)
	out := append(code2bytes(10012), pbBytes...)

	fmt.Println("out ", out)

	if err := ioutil.WriteFile("../player", out, 0644); err != nil {
		fmt.Println("Failed to write address book:", err)
	}

}

func toframe(i uint32, pbBytes []byte) *pb.Frame {
	payload := append(code2bytes(i), pbBytes...)
	return &pb.Frame{Payload: payload}
}
func code2bytes(i uint32) (b []byte) {
	b = append(b, byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
	return
}
func bytes2code(b []byte) (i uint32) {
	i = uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
	return
}

func frameCode(f *pb.Frame) uint32 {
	payload := f.Payload
	return bytes2code(payload[0:4])
}
func framePBbytes(f *pb.Frame) []byte {
	payload := f.Payload
	return payload[4:]
}
