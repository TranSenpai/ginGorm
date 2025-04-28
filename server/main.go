package main

import (
	router "main/internal/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Khởi tạo router chính và chạy server
	server := gin.Default()
	router.InitRouter(server)
	server.Run(":8080")
}
