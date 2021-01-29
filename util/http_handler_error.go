package util

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// APIError struct
type APIError struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	ErrorMessage string `json:"error,omitempty"`
	Date         string `json:"date"`
}

func (e APIError) Error() string {
	return e.Message + ""
}

// CreateAPIError func
func CreateAPIError(code int, messages ...string) *APIError {
	var message, errorMessage string

	switch length := len(messages); {
	case length > 1:
		message = messages[0]
		errorMessage = messages[1]
		break
	case length == 1:
		message = messages[0]
		break
	}

	return &APIError{
		Code:         code,
		Message:      message,
		ErrorMessage: errorMessage,
	}
}

// HandlerError func
func HandlerError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			apiError, ok := c.Errors[0].Err.(*APIError)
			if !ok {
				apiError = &APIError{
					Code:    http.StatusInternalServerError,
					Message: c.Errors[0].Err.Error(),
				}
			}
			apiError.Date = time.Now().Format(time.RFC3339)
			c.AbortWithStatusJSON(apiError.Code, apiError)
			return
		}
	}
}
