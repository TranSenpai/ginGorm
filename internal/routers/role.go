package router

import (
	"main/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoleRoute(server *gin.RouterGroup) {
	api := server.Group("/role")
	roleController := controller.RoleController{}
	api.POST("/role", roleController.CreateRole)
	api.DELETE("/:id", roleController.Delete)
	api.PUT("/:id", roleController.UpdateRole)
	api.GET("/:id", roleController.Search)
	api.GET("/roles", roleController.SearchAll)
}
