package main

import (
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := gin.Default()
	router.POST("/somePost", func(c *gin.Context) {
		fmt.Println(c.Request.ContentLength)
	})
	router.Run(":6060")
}
