package router

import (
	"main/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoomRoute(server *gin.RouterGroup) {
	api := server.Group("/room")
	roomController := controller.RoomController{}
	api.POST("/room", roomController.CreateRoom)
	api.DELETE("/:id", roomController.Delete)
	api.PUT("/:id", roomController.UpdateRoom)
	api.GET("/:id", roomController.Search)
	api.GET("/rooms", roomController.SearchAll)
}
