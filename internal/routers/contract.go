package router

import (
	"main/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterContract(server *gin.RouterGroup) {
	api := server.Group("/contract")
	contractController := controller.ContractController{}
	api.POST("", contractController.CreateContract)
	api.DELETE("/:studentcode", contractController.Delete)
	api.PUT("/:studentcode", contractController.UpdateContract)
	api.GET("/:studentcode", contractController.Search)
	api.GET("/all", contractController.SearchAll)
	api.GET("/byname/:fullname", contractController.SearchByName)

}
