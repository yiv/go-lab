package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("patti.config")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
	fmt.Println(viper.GetInt64("patti.coin"))

}
