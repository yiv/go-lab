package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

type GoogleIAP struct {
	Kind               string `json:"kind"`
	PurchaseTimeMillis string `json:"purchaseTimeMillis"`
	PurchaseState      string `json:"purchaseState"`
	ConsumptionState   bool   `json:"consumptionState"`
	DeveloperPayload   string `json:"developerPayload"`
}

func main() {
	data, err := ioutil.ReadFile("downloaded.json")
	if err != nil {
		log.Fatal(err)
	}
	//conf, err := google.JWTConfigFromJSON(data, "https://www.google.com/m8/androidpublisher")
	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/androidpublisher")

	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(oauth2.NoContext)

	res, err := client.Get("https://www.googleapis.com/androidpublisher/v2/applications/com.daily.poker.teenpatti/purchases/products/101/tokens/goomkddanebpljoieleeooni.AO-J1Ox9IpWeRD7wEoDmOq-5EO_cpy2b3Ss_C109iwEAYsMwc4k6Hy8aZ45jnTh4A-pC9JDNYa0IVfb3MPt3IVQviHwlZ8YcmgulZ6T2HB7sSzcBDRjCwCk")
	if err != nil {
		log.Println("err:", err)
	}

	fmt.Println("edwin 55", err)

	body, err := ioutil.ReadAll(res.Body)
	log.Println(string(body))

	appResult := &GoogleIAP{}

	err = json.Unmarshal(body, &appResult)
	log.Printf("Receipt return %+v \n", appResult)

	//To transfer to purchase time millisecond to time.Time
	time_duration, _ := strconv.ParseInt(appResult.PurchaseTimeMillis, 10, 64)
	log.Println(time_duration)

	time_purchase := time.Unix(time_duration/1000, 0)
	log.Println(time_purchase.Local())

	// Compare with receipt to make sure the receipt if valid.

	// 1. appResult.PurchaseTimeMillis need equal to purchaseTime get from App.
	// 2. appResult.ConsumptionState need to be 1. (User already consumed)
	// 3. appResult.PurchaseState need to be 0. (Order is completed.)

}
