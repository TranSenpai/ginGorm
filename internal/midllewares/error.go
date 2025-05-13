package midllewares

import (
	errorx "main/internal/utils/myerror"
	"net/http"

	"github.com/gin-gonic/gin"
)

// In context struct in gin package have Errors field type errorMsgs
// |
// v
// errorMsgs is type []*Error in gin package
// |
// v
// Error represents a error's specification.
// |
// v
// type Error struct {
// 	Err  error
// 	Type ErrorType
// 	Meta any
// }
// |
// v
// type error is a build interface in lib builtin in golang:
// type error interface {
// 	Error() string
// }

func ErrorHander() gin.HandlerFunc {
	return func(c *gin.Context) {
		// allow request run to controller
		c.Next()
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case *errorx.MyErr:
				c.AbortWithStatusJSON(e.Status, e)
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, e)
			}
		}
	}
}
