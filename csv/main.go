package main

import (
	"os"
	"fmt"
	//"io/ioutil"

	"github.com/gocarina/gocsv"
)

type Client struct { // Our example struct, you can use "-" to ignore a field
	Query      string `csv:"query"`
	Phone    string `csv:"phone"`
}

func main()  {
	fd, err := os.OpenFile("./xx.csv",os.O_RDWR, 0755)
	if err != nil{
		fmt.Println(err)
	}
	//buf, err := ioutil.ReadAll(fd)
	//if err != nil {
	//	fmt.Println("err", err.Error())
	//}else {
	//	fmt.Println("len", len(buf))
	//	fmt.Println(string(buf))
	//}

	clients := []*Client{}
	if err := gocsv.UnmarshalFile(fd, &clients); err != nil { // Load clients from file
		panic(err)
	}
	for _, client := range clients {
		fmt.Println("Hello", client.Query)
	}
}
