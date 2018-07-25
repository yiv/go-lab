package main

import (
	"crypto/tls"
	"fmt"
	"time"

	"gopkg.in/gomail.v2"
)

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "79743781@qq.com")
	m.SetHeader("To", "79743781@qq.com")
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Server Reboot Warning")
	m.SetBody("text/html", fmt.Sprintf("Server has rebooted at %s", time.Now().Format("2006-01-02 15:04")))

	d := gomail.NewDialer("smtp.qq.com", 587, "79743781@qq.com", "gkxkmxfdcrehbghd")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
