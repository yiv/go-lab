package main

import (
	"log"
	"sort"
	//"math/rand"
	"crypto/md5"
	"fmt"
	"net/http"
	"strings"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	log.Printf("%#v \n", r)
	r.ParseForm()
	var keys = make([]string, 0, 0)
	for key, value := range r.PostForm {
		log.Printf("PostForm #%v=%v#\n", key, value)
		if len(value) > 0 {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	var pList = make([]string, 0, 0)
	for _, key := range keys {
		var value = strings.TrimSpace(r.PostForm.Get(key))
		if len(value) > 0 {
			if key != "signature" {
				pList = append(pList, value)
			}
		}
	}
	pList = append(pList, "I2FAgut4LTv8o0AtdZmn7up9qb6KGJsR")
	var s = strings.Join(pList, "")
	md5string := fmt.Sprintf("%x", md5.Sum([]byte(s)))
	log.Printf("s #%v#\n", s)
	log.Printf("md5string #%v#\n", md5string)
}
func main() {
	http.HandleFunc("/shop/v1/confirmMOLOrder", HandleIndex)
	log.Fatal(http.ListenAndServe(":1080", nil))
}
