package router

import (
	"main/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterContract(server *gin.RouterGroup) {
	api := server.Group("/contract")
	contractController := controller.ContractController{}
	api.POST("/contract", contractController.CreateContract)
	api.DELETE("/:id", contractController.Delete)
	api.PUT("/:id", contractController.UpdateContract)
	api.GET("/:id", contractController.Search)
	api.GET("/rooms", contractController.SearchAll)
}
