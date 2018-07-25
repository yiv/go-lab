package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	h := sha256.New()
	str := "HDFC Bank|4036217121962950|INR|HDFC|Stonem85623985998921|27000364888888|DC|01|Txn Success|TXN_SUCCESS|1.00|2018-07-10 15:04:56.0|20180710111212800110168868500018912|np_X"
	h.Write([]byte(str))

	fmt.Printf("%x", h.Sum(nil))
}
