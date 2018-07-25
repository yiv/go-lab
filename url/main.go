package main

import (
	"fmt"
	"net/url"
)

func main() {
	r, _ := url.Parse("sign=cufDJKe2D75awwHRcAu9SthAx3UhMTdquCHUqhjohm%2BdaeZL2b7tZTja%2BY1WUN2JnWVksPeEJQs8QOmWU0Cvb3m8bBpWrBFph%2BSfUY%2BrmkaZ%2Bh51n1VPqZESKNwaDMONHUU7Gi%2Fe76lRe47HWHMXFXdo2IQbnlaK741tiCLl%2FGo%3D&notify_data=%7B%22order_status%22%3A%22S%22%2C%22trade_id%22%3A%2220171108210959880979%22%2C%22prd_name%22%3A%22several+chips%22%2C%22prd_type%22%3A%221%22%2C%22order_amt%22%3A%2210.000%22%2C%22pay_type%22%3A%22401%22%2C%22order_time%22%3A%222017-11-08+21%3A09%3A59%22%2C%22prd_id%22%3A%22f34d0f73bf024d578c61226b4b494b2e%22%2C%22order_id%22%3A%22295161651673071617%22%2C%22currency_id%22%3A%22INR%22%2C%22finish_time%22%3A%222017-11-08+21%3A10%3A05%22%7D&version=1.0&app_id=110073017&secure_mode=RSA")
	fmt.Println(r.Path)
}
