package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	id := int64(500)
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(id))
	fmt.Println(binary.BigEndian.Uint64(buf))
}
