package main

import (
	"fmt"
	"net/http"
	//"net/url"
	"encoding/json"
	"bytes"
	"io/ioutil"
)

type HomeManageReq struct {
	UserToken string            `db:"user_token" json:"user_token"`
	HomeId    int            `db:"home_id" json:"home_id"`
	Phone     string            `db:"phone" json:"phone"`
	Status    string            `db:"status" json:"status"`
	NickName  string            `db:"nick_name" json:"nick_name"`
	Avatar    string            `db:"avatar" json:"avatar"`
}

func main() {
	req := HomeManageReq{UserToken:"546932ea651161d82453b475fe99b4c712f696fa7e73edc61198060cf2019277",HomeId:1854,Phone:"18576676197", NickName:"ed555", Avatar:"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"}

	jsonStr,_ := json.Marshal(req)
	body := bytes.NewBuffer([]byte(jsonStr))

	//resp, err := http.Post("http://192.168.223.128:8091/home_manage/invite_unregister_member/?app_id=aa56900b5638a893455ef80619bbdef2&ver=1.0.0&vc=facaa1efe5ca275a482d781e2d59b440&lang=zh-cn", "application/json", body)
	resp, err := http.Post("https://api.1719.com/V4/home_manage/invite_unregister_member/?app_id=aa56900b5638a893455ef80619bbdef2&ver=1.0.0&vc=facaa1efe5ca275a482d781e2d59b440&lang=zh-cn", "application/json", body)
	//resp, err = http.Post("http://127.0.0.1:6060/somePost/", "application/json", body)
	if err == nil {
		fmt.Println(resp.Status)
		body,_ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	} else {
		fmt.Println(err)
	}
}
