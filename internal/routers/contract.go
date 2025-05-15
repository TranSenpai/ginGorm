package router

import (
	"main/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterContract(server *gin.RouterGroup) {
	api := server.Group("/contract")
	contractController := controller.NewContractController()
	api.POST("", contractController.CreateContract)
	api.DELETE("", contractController.Delete)
	api.PUT("", contractController.UpdateContract)
	api.GET("", contractController.Search)
	api.GET("/rooms", contractController.Search)
	api.GET("/total/rooms", contractController.SearchContractInRoom)
}
