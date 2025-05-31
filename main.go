package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.Handle("GET", "/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := engine.Run(":8887")
	if err != nil {
		panic(err)
	}
}
