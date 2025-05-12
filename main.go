package main

import (
	"main/internal/midllewares"
	router "main/internal/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create main router and run server
	server := gin.Default()
	// Register a middleware, Use() takes a gin.handlerFunc slice and return IRoute interface
	server.Use(
		midllewares.ErrorHander(),
	)
	router.InitRouter(server)
	server.Run(":8080")
}
