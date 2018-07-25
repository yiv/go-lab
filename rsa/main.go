package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"log"
	//"math/rand"
	"net/http"
	"sort"
	"strings"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	log.Printf("%#v \n", r)
	//if body, err := ioutil.ReadAll(r.Body); err != nil {
	//	log.Println(err.Error())
	//} else {
	//	log.Printf("#%v# \n", string(body))
	//}
	r.ParseForm()
	sign, err := base64.StdEncoding.DecodeString(r.PostForm.Get("sign"))
	//sign := r.PostForm.Get("sign")

	var keys = make([]string, 0, 0)
	for key, value := range r.PostForm {
		log.Printf("PostForm #%v=%v#\n", key, value)
		if len(value) > 0 {
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)

	var pList = make([]string, 0, 0)
	var allList = make([]string, 0, 0)
	for _, key := range keys {

		var value = strings.TrimSpace(r.PostForm.Get(key))
		if len(value) > 0 {
			if key == "sign" {
				allList = append(allList, key+"="+value)
				continue
			}
			allList = append(allList, key+"="+value)
			pList = append(pList, key+"="+value)
		}
	}
	var s = strings.Join(pList, "&")
	var all = strings.Join(allList, "&")
	log.Printf("all #%v#\n", all)
	log.Printf("s #%v#\n", s)

	key := `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDMVuZ899/BKhHQ17mTYh8d8qGVttfGNHY46kjEzCRE93h9nq51jxcQ2lfaDw5eRr9E0z3RyNrys1pFh69du/bxzmxiC0Y1cQM1Y+UH+KbLV2eJ/xoNWVyjyEPHqtcvECZB0Ipcsx5MwyeYTv4+QR1QupB5kLkWh/xFS7ioZcyiuwIDAQAB
-----END PUBLIC KEY-----`

	err = VerifyPKCS1v15([]byte(s), []byte(sign), []byte(key), crypto.SHA1)

	if err != nil {
		log.Println("edwin #5", err.Error())
	}

}

func VerifyPKCS1v15(src, sig, key []byte, hash crypto.Hash) error {
	var h = hash.New()
	h.Write(src)
	var hashed = h.Sum(nil)

	var err error
	var block *pem.Block
	block, _ = pem.Decode(key)
	if block == nil {
		return errors.New("public key error")
	}

	var pubInterface interface{}
	pubInterface, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	var pub = pubInterface.(*rsa.PublicKey)
	err = rsa.VerifyPKCS1v15(pub, hash, hashed, sig)
	if err != nil {
		log.Println("edwin #88", err.Error())
	} else {
		log.Println("edwin #77 success")
	}
	return err
}
func main() {
	http.HandleFunc("/shop/v1/confirmUCOrder", HandleIndex)
	log.Fatal(http.ListenAndServe(":1080", nil))
}
