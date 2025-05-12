package midllewares

import (
	errorx "main/internal/utils/myerror"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
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
			case errorx.MyErr:
				c.AbortWithStatusJSON(http.StatusBadRequest, e.Message)
			case *errorx.MyErr:
				c.AbortWithStatusJSON(http.StatusBadRequest, e.Message)
			case *mysql.MySQLError:
				// Error number: 1062; Symbol: ER_DUP_ENTRY; SQLSTATE: 23000
				// Message: Duplicate entry '%s' for key %d
				err := errorx.NewMyError(http.StatusConflict, "Modify database failed", e, time.Now())
				c.AbortWithStatusJSON(http.StatusBadRequest, err)
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": "Service Unavailable"})
			}
		}
	}
}

// func HandleError(c *gin.Context, err any) bool {
// 	if err == nil {
// 		return false
// 	}
// 	switch e := err.(type) {
// 	case errorx.MyErr:
// 		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
// 			Timestamp: time.Now(),
// 			Status:    e.Status,
// 			Message:   e.Message,
// 			Error:     e.Err.Error(),
// 		})
// 	case *errorx.MyErr:
// 		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
// 			Timestamp: time.Now(),
// 			Status:    e.Status,
// 			Message:   e.Message,
// 			Error:     e.Err.Error(),
// 		})
// 	case *mysql.MySQLError:
// 		// Error number: 1062; Symbol: ER_DUP_ENTRY; SQLSTATE: 23000
// 		// Message: Duplicate entry '%s' for key %d
// 		// The message returned with this error uses the format string for ER_DUP_ENTRY_WITH_KEY_NAME.
// 		if e.Number == 1062 {
// 			c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
// 				Timestamp: time.Now(),
// 				Status:    http.StatusConflict,
// 				Message:   "The student code or phone or mail or sign have been existed",
// 				Error:     e.Message,
// 			})
// 		} else {
// 			c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
// 				Timestamp: time.Now(),
// 				Status:    http.StatusInternalServerError,
// 				Message:   "Internal Server Error",
// 				Error:     e.Message,
// 			})
// 		}

// 	default:
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{
// 			Timestamp: time.Now(),
// 			Status:    http.StatusInternalServerError,
// 			Message:   "Internal Server Error",
// 			Error:     err,
// 		})
// 	}
// 	return true
// }
