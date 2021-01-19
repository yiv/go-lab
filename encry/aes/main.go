package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

func main() {

	key := "myverystrongpasswordo32bitlength"
	plainText := "Hello 8gwifi.org"
	ct := Encrypt([]byte(key), plainText)
	fmt.Printf("Original Text:  %s\n", plainText)
	fmt.Printf("AES Encrypted Text:  %s\n", ct)
	Decrypt([]byte(key), ct)
}

func Encrypt(key []byte, plaintext string) string {
	c, err := aes.NewCipher(key)
	if err != nil {
		_ = fmt.Errorf("NewCipher(%d bytes) = %s", len(key), err)
		panic(err)
	}
	out := make([]byte, len(plaintext))
	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out)
}

func Decrypt(key []byte, ct string) {
	ciphertext, _ := hex.DecodeString(ct)
	c, err := aes.NewCipher(key)
	if err != nil {
		_ = fmt.Errorf("NewCipher(%d bytes) = %s", len(key), err)
		panic(err)
	}
	plain := make([]byte, len(ciphertext))
	c.Decrypt(plain, ciphertext)
	s := string(plain[:])
	fmt.Printf("AES Decrypyed Text:  %s\n", s)
}
