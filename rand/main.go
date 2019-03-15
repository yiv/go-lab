package main

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	mrand "math/rand"
	"time"
)

func main() {
	RandUid()
}

func cryptoRand() {
	nBig, err := crand.Int(crand.Reader, big.NewInt(10))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()
	fmt.Printf("%d\t", n)

}

func mathRand() {
	for i := 0; i < 100; i ++ {
		fmt.Printf("%d\t", mrand.Intn(2))
	}
}

func mathRandPrintSeed() {
	for i := 0; i < 10; i ++ {
		seed := time.Now().UnixNano()
		mrand.Seed(seed)
		fmt.Printf("seed %v , %d\n", seed, mrand.Intn(100))
	}
}

func RandUid() {
	mrand.Seed(time.Now().UnixNano())
	//var uids string
	for i := 0; i < 973; i++ {
		r := mrand.Int31n(100000000) + 900000000
		uid := fmt.Sprintf(`"%v"`, r)
		fmt.Println(uid)
	}

}
