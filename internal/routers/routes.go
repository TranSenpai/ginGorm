package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	// Quy ước /api/version
	apiv1 := engine.Group("/api/v1")
	RegisterContract(apiv1)
}
