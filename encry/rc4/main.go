package main

import (
	"crypto/rc4"
	"fmt"
)

func main() {
	jsonstr := []byte("{\"name\":\"edwin\"}")
	cipher, err := rc4.NewCipher([]byte("f63dfeafe6bd2f74fedcf754c89d25ad"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("json string: ", string(jsonstr))
	cryptodst := make([]byte, len(jsonstr))
	cipher.XORKeyStream(cryptodst, jsonstr)
	fmt.Println("crypto []byte: ", cryptodst)
	fmt.Println("crypto string: ", string(cryptodst))
	uncryptodst := make([]byte, len(cryptodst))

	c2, err := rc4.NewCipher([]byte("f63dfeafe6bd2f74fedcf754c89d25ad"))
	c2.XORKeyStream(uncryptodst, cryptodst)
	fmt.Println("uncrypto: []byte", uncryptodst)
	fmt.Println("uncrypto string: ", string(uncryptodst))

}
