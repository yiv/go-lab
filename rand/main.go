package main

import (
	"fmt"
	mrand "math/rand"
	"time"
	crand "crypto/rand"
	"math/big"
)

func main() {
	fmt.Println("rand number from crypto rand")
	for i := 0; i < 10; i ++ {
		cryptoRand()
	}
	fmt.Println("rand number from crypto rand")

	fmt.Println("rand number from math rand")
	mrand.Seed(time.Now().UnixNano())
	mathRand()
	fmt.Println("rand number from math rand")

}

func cryptoRand()  {
	nBig, err := crand.Int(crand.Reader, big.NewInt(10))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()
	fmt.Printf("%d\t", n)

}

func mathRand()  {
	for i := 0; i < 100; i ++{
		fmt.Printf("%d\t", mrand.Intn(2))
	}
}
