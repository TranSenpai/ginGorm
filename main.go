package main

import (
	"fmt"
	router "main/internal/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create main router and run server
	server := gin.Default()
	server.Use()
	router.InitRouter(server)
	server.Run(":8080")

	var num uint8
	var str string

	fmt.Print(num)
	fmt.Print(str)
}
