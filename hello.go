package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/hello", func(context *gin.Context) {
		context.String(200, "Hello World!")
	})
	router.Run(":8080")
}
