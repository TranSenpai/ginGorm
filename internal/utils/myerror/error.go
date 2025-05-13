package myerror

import (
	"fmt"
	"time"
)

// const (
// 	// 2xx Success
// 	StatusOK        = 200 // Successfully
// 	StatusCreated   = 201 // Create successfully
// 	StatusNoContent = 204 // Create successfully, no content return

// 	// 4xx Client Error
// 	StatusBadRequest          = 400 // Request wrong format or lack of field
// 	StatusUnauthorized        = 401 // Not logging or wrong token
// 	StatusForbidden           = 403 // Don't have the necessary permissions to access the requested resource.
// 	StatusNotFound            = 404 // Resource not found
// 	StatusConflict            = 409 // Duplicate data
// 	StatusUnprocessableEntity = 422 // Validate data (invalid or doesn't meet the server's expectations. )
// 	StatusTooManyRequests     = 429 // Too many requests to a server within a given time period

// 	// 5xx Server Error
// 	StatusInternalServerError = 500 // Server encountered an unexpected condition preventing it from fulfilling the client's request
// 	StatusBadGateway          = 502 // Error upstream server (means the server couldn't fulfill the request because the server it needed to reach sent back a faulty or invalid response.)
// 	StatusServiceUnavailable  = 503 // Server overload or scheduled maintenance. Server is temporarily unable to handle the request due to a temporary issue.
// )

type MyErr struct {
	Status    int       `json:"-"`
	Message   string    `json:"message,omitempty"`
	Err       error     `json:"error,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (e *MyErr) Error() string {
	return fmt.Sprintf("Status: %d, Message: %s, Error: %v", e.Status, e.Message, e.Err)
}

func NewMyError(status int, msg string, err error, time time.Time) error {
	return &MyErr{
		Status:    status,
		Message:   msg,
		Err:       err,
		Timestamp: time,
	}
}
