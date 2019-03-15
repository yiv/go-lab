package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	stdjwt "github.com/dgrijalva/jwt-go"
)

// .\flag.exe -uid 16d033da-8c8e-4826-a842-670027612455 -host dev -action push   -coin 40000 -delta false
// .\flag.exe -uid 16d033da-8c8e-4826-a842-670027612455 -host dev -action pull

const (
	JwtHmacSecret = "cb1661c9b6c63c9ddabb60c3f3c64bfb"
)

var (
	actoin = flag.String("action", "push", "pull：获取玩家数据，push：修改玩家数据")
	uid    = flag.String("uid", "16d033da-8c8e-4826-a842-670027612455", "玩家ID")
	coin   = flag.Int64("coin", 100, "增减金币")
	delta  = flag.String("delta", "true", "是否增量修改")
	host   = flag.String("host", "test", "服务器，dev：开发服， test：测试服")
)

var HostAddr = map[string]string{
	"dev":  "http://192.168.1.12:10010",
	"test": "http://106.75.156.157:30601",
}

func main() {
	flag.Parse()
	server := HostAddr[*host]
	switch *actoin {
	case "pull":
		PullUserInfo(server, *uid)
	case "push":
		PushUserInfo(server, *uid, *coin)
	}

}

func PullUserInfo(server, uid string) {
	u := "/user/v1/pullUserInfo"
	j := []byte(`{"keys":["uid","nick", "coin", "gem"]}`)
	r := postCliHttps(server, u, j, NewJWTToken(uid))
	fmt.Println("request Body:", string(j))
	fmt.Println("response Body:", string(r))
}

func PushUserInfo(server, uid string, coin int64) {
	u := "/user/v1/pushUserInfo"
	j := []byte(fmt.Sprintf(`{"delta":%v, "key":"coin","adj": %v}`, *delta, coin))
	r := postCliHttps(server, u, j, NewJWTToken(uid))
	fmt.Println("request Body:", string(j))
	fmt.Println("response Body:", string(r))
}

func postCliHttps(host, u string, j []byte, jwt string) []byte {
	reqUrl := host + u
	fmt.Println("URL: ", reqUrl)
	s := rc4Crypt(j)
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
	tr := &http.Transport{
		//TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: time.Second * 5}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("err", err.Error())
		os.Exit(0)
	}
	//fmt.Println("Request Header: ", req.Header)

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	return rc4Crypt(body)
}

func rc4Crypt(s []byte) []byte {
	return s
}

func NewJWTToken(uid string) (tokenString string) {
	claims := stdjwt.MapClaims{"uid": uid, "exp": time.Now().Add(time.Hour * 24).Unix()}
	token := stdjwt.NewWithClaims(stdjwt.SigningMethodHS256, claims)
	token.Header["kid"] = "kid-header"
	tokenString, err := token.SignedString([]byte(JwtHmacSecret))
	if err != nil {
		return ""
	}
	return
}
