package main

import (
	"crypto/rc4"
	"fmt"
	"encoding/base64"
	"net/http"
	"bytes"
	"os"
	"io/ioutil"
	"time"
)

const (
	Uid int64 = 671845128699906
	//Uid int64 = 946190157611011

	//host string = "http://192.168.1.200:10070"
	//host string = "http://192.168.1.205:10070"
	host = "http://192.168.1.51:10070"
	//host  = "http://frontapi.poker666.in"
	//host string = "http://poker666.in:10070"
	jwt = "eyJhbGciOiJIUzI1NiIsImtpZCI6ImtpZC1oZWFkZXIiLCJ0eXAiOiJKV1QifQ.eyJleHAiOjE1MTAxNDQ2NzcsInVpZCI6ODE1OTE1NzcxMzYzMzI5fQ.nrVxdV3RTTC1DboyhjUeOJmMLj4sGO41XF_DamozDQE"
)

func main() {
	TestGetDeviceID()
}

func rc4Crypt(s []byte) []byte {
	key := []byte("f63dfeafe6bd2f74fedcf754c89d25ad")
	c, _ := rc4.NewCipher(key)
	d := make([]byte, len(s))
	c.XORKeyStream(d, s)
	return d
}

func postCli(u string, j []byte, jwt string) []byte {
	reqUrl := host + u
	fmt.Println("URL: ", reqUrl)
	s := rc4Crypt(j)
	fmt.Printf("edwin #4 %s \n", base64.StdEncoding.EncodeToString(s))
	var (
		req  *http.Request
		resp *http.Response
		err  error
	)
	req, err = http.NewRequest("POST", reqUrl, bytes.NewBuffer(s))
	if err != nil {
		fmt.Println("err", err.Error())
		os.Exit(0)
	}
	req.Header.Set("Content-Type", "text/plain")
	if jwt != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwt))
	}
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("err", err.Error())
		os.Exit(0)
	}
	fmt.Println("Request Header: ", req.Header)

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body string:", base64.StdEncoding.EncodeToString(body))
	body = rc4Crypt(body)
	fmt.Println("response Body string:", base64.StdEncoding.EncodeToString(body))
	return body
}

func TestGetDeviceID() {
	start := time.Now()
	u := "/device/v1/getDeviceID"
	//j := []byte(`{"Imei":"k2k2k","Imsi":"ddderw","Mac":"223adsd"}`)
	j := []byte(`{"imsi":"460077054189144","imei":"86547302505633","mac":"24:09:95:37:70:44"}`)
	r := postCli(u, j, "")

	fmt.Println("response Body string:", string(r))
	fmt.Println(time.Now().Sub(start))
	//tdid = res.Did
}
