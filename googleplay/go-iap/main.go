package main

import (
	"github.com/dogenzaka/go-iap/playstore"

	"io/ioutil"
	"log"
)

func main() {
	// You need to prepare a public key for your Android app's in app billing
	// at https://console.developers.google.com.
	jsonKey, err := ioutil.ReadFile("downloaded.json")
	if err != nil {
		log.Fatal(err)
	}

	client, err := playstore.New(jsonKey)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.VerifySubscription("com.daily.poker.teenpatti", "101", "goomkddanebpljoieleeooni.AO-J1Ox9IpWeRD7wEoDmOq-5EO_cpy2b3Ss_C109iwEAYsMwc4k6Hy8aZ45jnTh4A-pC9JDNYa0IVfb3MPt3IVQviHwlZ8YcmgulZ6T2HB7sSzcBDRjCwCk")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
