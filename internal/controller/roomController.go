package controller

import (
	model "main/internal/models"
	service "main/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoomController struct{}

func (rc RoomController) CreateRoom(c *gin.Context) {
	id := c.Param("id")
	var room model.Room
	serviceRoom := service.GetRoomService()

	err := c.ShouldBindJSON(&room)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Some field are not right"})
		return
	}

	_, err = serviceRoom.Search(id)
	if err == nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "This room has already exist"})
		return
	}

	err = serviceRoom.CreateRoom(&room)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Can not create room"})
		return
	}

	data, err := serviceRoom.Search(room.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not retrieve created room"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Your created room",
		"created room": data,
	})
}

func (rc RoomController) UpdateRoom(c *gin.Context) {
	id := c.Param("id")
	var room model.Room
	serviceRoom := service.GetRoomService()

	// Ép data từ JSON từ request HTTP sang type struct
	err := c.ShouldBindJSON(&room)
	if err != nil {
		// Có một lỗi cú pháp trong yêu cầu và yêu cầu bị từ chối.
		c.JSON(http.StatusBadRequest, gin.H{"message": "Some field are not right"})
		return
	}

	_, err = serviceRoom.Search(room.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "This room does not exist"})
		return
	}

	err = serviceRoom.UpdateRoom(id, &room)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not update room"})
		return
	}

	data, err := serviceRoom.Search(room.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not retrieve updated room"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Your updated room",
		"updated room": data,
	})
}

func (rc RoomController) Delete(c *gin.Context) {
	id := c.Param("id")
	serviceRoom := service.GetRoomService()

	role, err := serviceRoom.Search(id)
	if err.Error() == "not found" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	err = serviceRoom.DeleteRoom(role.ID)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "role does not exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete successfully"})
}

func (rc RoomController) Search(c *gin.Context) {
	id := c.Param("id")
	var room model.Room
	serviceRoom := service.GetRoomService()

	room, err := serviceRoom.Search(id)

	// gin.H is a type alias for map[string]interface{}.
	// It's a shorthand notation used to create maps that store data
	// with string keys and can hold values of any type.
	// This type allows you to easily define and use dynamic
	// JSON-like structures within your Gin applications.

	if err.Error() == "not found" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "room found",
		"data":    room,
	})
}

func (rc RoomController) SearchAll(c *gin.Context) {
	var rooms []model.Room
	serviceRoom := service.GetRoomService()

	rooms, err := serviceRoom.SearchAll()

	if err.Error() == "dont have any room in table" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Empty room table"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "rooms found",
		"data":    rooms,
	})
}
