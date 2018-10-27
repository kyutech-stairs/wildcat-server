package main

import (
	_ "fmt"

	"github.com/kyutech-stairs/wildcat-server/router"
)

func main() {
	// r := gin.Default()
	r := router.GetRouter()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.Run(":3000")
}
