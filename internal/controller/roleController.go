package controller

import (
	model "main/internal/models"
	"main/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleController struct{}

// 1XX - Thông tin: Yêu cầu được chấp nhận hoặc quá trình tiếp tục.

// 2XX - Thành công: Xác nhận rằng hành động đã hoàn tất thành công hoặc đã được hiểu.

// 3XX - Chuyển hướng: Client phải thực hiện hành động bổ sung để hoàn thành yêu cầu.

// 4XX - Lỗi từ client chỉ ra rằng yêu cầu không thể hoàn thành hoặc chứa cú pháp sai.
// Mã lỗi 4xx sẽ hiện ra khi có lỗi từ phía người dùng, chủ yếu là do không đưa ra một yêu cầu hợp lệ.

// 5XX - Lỗi từ phía máy chủ: Cho biết máy chủ không thể hoàn tất yêu cầu được cho là hợp lệ.
// Khi duyệt web và bắt gặp các lỗi 5xx, bạn chỉ có thể chờ đợi, vì lúc này lỗi xuất phát từ phía
// máy chủ của dịch vụ web, không có cách nào can thiệp để sửa lỗi ngoài việc ngồi chờ bên máy chủ xử lý xong.

func (rc RoleController) CreateRole(c *gin.Context) {
	var role model.Role
	serviceRole := service.GetRoleService()

	err := c.ShouldBindJSON(&role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Some field are not right"})
		return
	}

	_, err = serviceRole.Search(role.ID)
	if err == nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "This role has already exist"})
		return
	}

	err = serviceRole.CreateRole(&role)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Can not create role"})
		return
	}

	data, err := serviceRole.Search(role.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not retrieve created role"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":      "Your created role",
		"updated role": data,
	})
}

func (rc RoleController) UpdateRole(c *gin.Context) {
	var role model.Role
	serviceRole := service.GetRoleService()

	// Ép data từ JSON từ request HTTP sang type struct
	err := c.ShouldBindJSON(&role)
	if err != nil {
		// Có một lỗi cú pháp trong yêu cầu và yêu cầu bị từ chối.
		c.JSON(http.StatusBadRequest, gin.H{"message": "Some field are not right"})
		return
	}

	_, err = serviceRole.Search(role.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "This role does not exist"})
		return
	}

	err = serviceRole.UpdateRole(role.ID, &role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not update role"})
		return
	}

	data, err := serviceRole.Search(role.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not retrieve updated role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Your updated role",
		"updated role": data,
	})
}

func (rc RoleController) Delete(c *gin.Context) {
	id := c.Param("id")
	serviceRole := service.GetRoleService()

	role, err := serviceRole.Search(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	err = serviceRole.DeleteRole(role.ID)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "role does not exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete successfully"})
}

func (rc RoleController) Search(c *gin.Context) {
	id := c.Param("id")
	serviceRole := service.GetRoleService()

	role, err := serviceRole.Search(id)

	// gin.H is a type alias for map[string]interface{}.
	// It's a shorthand notation used to create maps that store data
	// with string keys and can hold values of any type.
	// This type allows you to easily define and use dynamic
	// JSON-like structures within your Gin applications.

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "role found",
		"data":    role,
	})
}

func (rc RoleController) SearchAll(c *gin.Context) {
	serviceRole := service.GetRoleService()
	roles, err := serviceRole.SearchAll()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Empty role table"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "roles found",
		"data":    roles,
	})
}
