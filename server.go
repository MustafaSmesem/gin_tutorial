package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const PORT = ":1234"

func main() {
	fmt.Println("Hello gin")
	server := gin.Default()

	server.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello world from gin!!",
		})
	})

	server.Run(PORT)
}
