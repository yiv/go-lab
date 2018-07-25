package main

import (
	"crypto/rc4"
	"fmt"
	"encoding/base64"
)

func main() {
	//rc4bytes := RC4Crypt([]byte(`{"key":"中"}`))
	rc4bytes := RC4Crypt([]byte(`中`))
	fmt.Println(string(rc4bytes))
	fmt.Println(base64.StdEncoding.EncodeToString(rc4bytes))
	fmt.Println(rc4bytes)
}
func RC4Crypt(s []byte) []byte {
	key := []byte("f63dfeafe6bd2f74fedcf754c89d25ad")
	c, _ := rc4.NewCipher(key)
	d := make([]byte, len(s))
	c.XORKeyStream(d, s)
	return d
}
