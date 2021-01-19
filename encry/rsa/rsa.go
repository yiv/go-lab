package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

const PublicPerm = `-----BEGIN PUBLIC KEY----- 
###########
-----END PUBLIC KEY-----`

const PrivatePerm = `-----BEGIN PRIVATE KEY----- 
#######################
-----END PRIVATE KEY-----`

func main() {

	// Generate Alice RSA keys Of 2048 Buts

	var (
		privatePemBlock *pem.Block
		pemParseResult  interface{}
		privateKey      *rsa.PrivateKey
		err             error

		publicPemBlock *pem.Block
		publicKey      *rsa.PublicKey
	)

	privatePemBlock, _ = pem.Decode([]byte(PrivatePerm))
	pemParseResult, err = x509.ParsePKCS8PrivateKey(privatePemBlock.Bytes)
	privateKey = pemParseResult.(*rsa.PrivateKey)

	if err != nil {
		fmt.Println("edwin 44", err)
		os.Exit(1)
	}

	publicPemBlock, _ = pem.Decode([]byte(PublicPerm))
	pemParseResult, err = x509.ParsePKIXPublicKey(publicPemBlock.Bytes)
	publicKey = pemParseResult.(*rsa.PublicKey)
	if err != nil {
		fmt.Println("edwin 66", err)
		os.Exit(1)
	}

	secretMessage := "app_id=2v4P5czYUJZtdY0H0NR4&order_no=130954676521668608&third_order_no=162526189968162816&ts=1610625693"
	fmt.Println("Original Text  ", secretMessage)
	signature := SignPKCS1v15(secretMessage, *privateKey)
	fmt.Println("Signature :  ", signature)
	verify := VerifyPKCS1v15(signature, secretMessage, *publicKey)
	fmt.Println(verify)
}

func SignPKCS1v15(plaintext string, privKey rsa.PrivateKey) string {
	// crypto/rand.Reader is a good source of entropy for blinding the RSA
	// operation.
	rng := rand.Reader
	//hashed := sha256.Sum256([]byte(plaintext))
	hashed := sha1.Sum([]byte(plaintext))

	signature, err := rsa.SignPKCS1v15(rng, &privKey, crypto.SHA1, hashed[:])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
		return "Error from signing"
	}
	return base64.StdEncoding.EncodeToString(signature)
}

func VerifyPKCS1v15(signature string, plaintext string, pubkey rsa.PublicKey) string {
	sig, _ := base64.StdEncoding.DecodeString(signature)
	//hashed := sha256.Sum256([]byte(plaintext))
	hashed := sha1.Sum([]byte(plaintext))

	err := rsa.VerifyPKCS1v15(&pubkey, crypto.SHA1, hashed[:], sig)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error from verification: %s\n", err)
		return "Error from verification:"
	}
	return "Signature Verification Passed"
}
