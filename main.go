package main

import (
	router "main/internal/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create main router and run server
	server := gin.Default()
	router.InitRouter(server)
	server.Run(":8080")
}
