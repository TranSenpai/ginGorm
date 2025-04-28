package myerror

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Status int

const (
	// 2xx Success
	StatusOK        = 200 // Successfully
	StatusCreated   = 201 // Create successfully
	StatusNoContent = 204 // Create successfully, no content return

	// 4xx Client Error
	StatusBadRequest          = 400 // Request wrong format or lack of field
	StatusUnauthorized        = 401 // Not logging or wrong token
	StatusForbidden           = 403 // Don't have the necessary permissions to access the requested resource.
	StatusNotFound            = 404 // Resource not found
	StatusConflict            = 409 // Duplicate data
	StatusUnprocessableEntity = 422 // Validate data (invalid or doesn't meet the server's expectations. )
	StatusTooManyRequests     = 429 // Too many requests to a server within a given time period

	// 5xx Server Error
	StatusInternalServerError = 500 // Server encountered an unexpected condition preventing it from fulfilling the client's request
	StatusBadGateway          = 502 // Error upstream server (means the server couldn't fulfill the request because the server it needed to reach sent back a faulty or invalid response.)
	StatusServiceUnavailable  = 503 // Server overload or scheduled maintenance. Server is temporarily unable to handle the request due to a temporary issue.
)

type MyErr struct {
	Status  Status
	Message string
	Err     error
}

func (e MyErr) Error() string {
	return fmt.Sprintf("Status: %d, Message: %s, Error: %v", e.Status, e.Message, e.Err)
}

type ErrorResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Status    Status    `json:"status"`
	Message   string    `json:"message"`
	Error     any       `json:"error,omitempty"`
}

func New(status Status, msg string, err error) MyErr {
	return MyErr{
		Status:  status,
		Message: msg,
		Err:     err,
	}
}

func HandleError(c *gin.Context, err any) {
	if err == nil {
		return
	}
	switch e := err.(type) {
	case MyErr:
	case *MyErr:
		// AbortWithStatusJSON stops the chain(prevent Gin from continuing to run the router),
		// writes the status code and return a JSON body
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Timestamp: time.Now(),
			Status:    e.Status,
			Message:   e.Message,
			Error:     e.Err.Error(),
		})
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{
			Timestamp: time.Now(),
			Status:    http.StatusInternalServerError,
			Message:   "Internal Server Error",
			Error:     err,
		})
	}
}
