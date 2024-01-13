package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()
	route.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	route.Run() // listen and serve on 0.0.0.0:8080
}