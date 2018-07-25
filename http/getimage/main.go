package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	httpCli := http.Client{Timeout: time.Duration(10 * time.Second)}
	res, _ := httpCli.Get("https://fb-s-d-a.akamaihd.net/h-ak-fbx/v/t1.0-1/p160x160/20620883_104195513616818_1460809289114780922_n.jpg?oh=a63cc41897990f945327839173aeebe0&oe=5A14E7B9&__gda__=1511437097_4de514c3f107233c58cc9b5c95a279c5")
	//res, _ := httpCli.Get("https://fb-s-d-a.akamaihd.net/h-ak-fbx/v/t1.0-1/p50x50/20620883_104195513616818_1460809289114780922_n.jpg?oh=490de084711e4c3e412a3dc975875036&oe=5A124728&__gda__=1511356191_09ff03466af6137bafb7fcb78f5090a3")
	//res, _ := httpCli.Get("https://fb-s-d-a.akamaihd.net/h-ak-fbx/v/t1.0-1/p100x100/20620883_104195513616818_1460809289114780922_n.jpg?oh=ca0923866ce4830336ac8ba8cfd7c89a&oe=5A1F6B26&__gda__=1512581046_3e4455499b8f007cf0c5736e1a9a94aa")

	buffer, _ := ioutil.ReadAll(res.Body)
	//buffer := make([]byte, res.ContentLength)
	//res.Body.Read(buffer)
	fmt.Println(buffer)
	fmt.Println(res.ContentLength)

	localPath := "a/b"
	filePath := fmt.Sprintf("%s/x.jpg", localPath)
	if err := os.MkdirAll(localPath, 0664); err != nil {
		fmt.Println(err)
		return
	}
	if err := ioutil.WriteFile(filePath, buffer, 0664); err != nil {
		fmt.Println(err)
		return
	}
}
