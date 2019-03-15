package main

import (
	"os"
	"fmt"
)

func main()  {
	AppendLine()
}

func AppendLine()  {
	_, err := os.Stat("a/2")
	if os.IsNotExist(err) {
		os.MkdirAll("a/2/test.txt", 0755)
	}
	fd, err := os.OpenFile("a/2/test.txt",os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	for i := 0; i < 50; i++{
		fd.WriteString(fmt.Sprintf("append line %v\n", i))
	}

	if err != nil{
		fmt.Println(err)
	}
}

func Write()  {
	_, err := os.Stat("a/2")
	if os.IsNotExist(err) {
		os.MkdirAll("a/2/test.txt", 0755)
	}
	fd, err := os.OpenFile("a/2/test.txt",os.O_RDWR|os.O_CREATE, 0755)
	fd.Write([]byte{55})
	if err != nil{
		fmt.Println(err)
	}
}