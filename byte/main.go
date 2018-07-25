package main

import "fmt"

func main() {

	code := uint32(15000)
	//pb := []byte("hello")
	codebytes := code2bytes(code)
	fmt.Println(codebytes)
	c := bytes2code(codebytes)
	fmt.Println(c)

	long := []byte{1, 2, 3, 4,5,6,7}
	fmt.Println(long[0:4])
	fmt.Println(long[4:5])

}

func code2bytes(i uint32) (b []byte) {
	b = append(b, byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
	return
}
func bytes2code(b []byte) (i uint32) {
	i = uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
	return
}
