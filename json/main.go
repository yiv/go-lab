package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	NilArray()
}

func js2structWithinMap() {
	type loginGuestReq struct {
		Did     string           `json:"did"`
		Unionid string           `json:"unionid"`
		Imei    string           `json:"imei"`
		Imsi    string           `json:"imsi"`
		Mac     string           `json:"mac"`
		Psystem string           `json:"psystem"`
		Pmodel  string           `json:"pmodel"`
		Nick    string           `json:"nick"`
		Gender  bool             `json:"gender"`
		Others  map[string]int32 `json:"others"`
	}
	req := loginGuestReq{}
	js := `{"did":"70f36b40-d1bc-44e3-9791-9615f3633b45","unionid":"ifunbow","imei":"55932fx23","imsi":"dkdkd12xf9222dk","psystem":"ios","pmode":"iphone5","nick":"一二三四五六七八九十","gender":true,"others":{"character":2}}`
	json.Unmarshal([]byte(js), &req)
	fmt.Println(req)

	res := loginGuestReq{
		Did:     "didxx",
		Unionid: "unxx",
		Imei:    "imei",
		Imsi:    "imsi",
		Psystem: "psystem",
		Pmodel:  "pmodel",
		Nick:    "nick",
		Gender:  true,
		Others:  map[string]int32{"character": 2},
	}
	s, _ := json.Marshal(res)
	fmt.Printf("%s\n", s)
}
func js2struct() {

	type P struct {
		Name string
	}
	js := "{\"name\":\"edwin\"}"
	p := P{}
	json.Unmarshal([]byte(js), &p)
	fmt.Println(p)
}
func js2array() {
	var str []string
	js := `["abc","123"]`
	json.Unmarshal([]byte(js), &str)
	fmt.Println(str)
}
func js2map() {
	var m map[string]int32
	js := `{"aaa":123,"bbb":456}`
	json.Unmarshal([]byte(js), &m)
	fmt.Println(m)
}

func NilArray() {
	ary := []string{}
	res := map[string][]string{
		"empty": ary,
	}

	r, _ := json.Marshal(res)
	fmt.Println(string(r))

	type Data struct {
		Name string
	}
	var ary2 []Data
	//ary2 = make([]Data, 0, 1)
	ary2 = []Data{}

	res2 := map[string][]Data{
		"empty": ary2,
	}
	r, _ = json.Marshal(res2)
	fmt.Println(string(r))

	//ary3 := []interface{}{}

}
