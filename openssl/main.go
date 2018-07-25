package main

import (
	"fmt"

	"github.com/spacemonkeygo/openssl"
	"encoding/base64"
)

const PaytmMerchantKey = `7EOxauh@GSecdiZ!`

func main() {
	out, _ := encrypt([]byte("test"))
	fmt.Printf("out %v\n", base64.StdEncoding.EncodeToString(out))
	out, _ = decrypt(out)
	fmt.Printf("%s\n", out)
}

func encrypt(input []byte) (output []byte, err error) {
	iv := "@@@@&&&&####$$$$"
	crypter, _ := NewCrypter([]byte(PaytmMerchantKey), []byte(iv))
	output, err = crypter.Encrypt(input)
	return
}

func decrypt(input []byte) (output []byte, err error) {
	iv := "@@@@&&&&####$$$$"
	crypter, _ := NewCrypter([]byte(PaytmMerchantKey), []byte(iv))
	output, err = crypter.Decrypt(input)
	return
}

type Crypter struct {
	key    []byte
	iv     []byte
	cipher *openssl.Cipher
}

func NewCrypter(key []byte, iv []byte) (*Crypter, error) {
	cipher, err := openssl.GetCipherByName("aes-128-cbc")
	if err != nil {
		return nil, err
	}

	return &Crypter{key, iv, cipher}, nil
}

func (c *Crypter) Encrypt(input []byte) ([]byte, error) {
	ctx, err := openssl.NewEncryptionCipherCtx(c.cipher, nil, c.key, c.iv)
	if err != nil {
		return nil, err
	}

	cipherbytes, err := ctx.EncryptUpdate(input)
	if err != nil {
		return nil, err
	}

	finalbytes, err := ctx.EncryptFinal()
	if err != nil {
		return nil, err
	}

	cipherbytes = append(cipherbytes, finalbytes...)
	return cipherbytes, nil
}

func (c *Crypter) Decrypt(input []byte) ([]byte, error) {
	ctx, err := openssl.NewDecryptionCipherCtx(c.cipher, nil, c.key, c.iv)
	if err != nil {
		return nil, err
	}

	cipherbytes, err := ctx.DecryptUpdate(input)
	if err != nil {
		return nil, err
	}

	finalbytes, err := ctx.DecryptFinal()
	if err != nil {
		return nil, err
	}

	cipherbytes = append(cipherbytes, finalbytes...)
	return cipherbytes, nil
}
