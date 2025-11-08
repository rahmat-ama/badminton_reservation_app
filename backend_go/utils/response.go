package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SuccessResponse(message string, data interface{}) gin.H {
	return gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	}
}

func ErrorResponse(message string) gin.H {
	return gin.H{
		"status":  "error",
		"message": message,
	}
}
