package main

import (
	router "main/internal/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	srv := gin.New()
	router.InitRouter(srv)
	server := gin.Default()
	router.InitRouter(server)
	server.Run(":8080")
}
