package paymentcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPayment(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Payment retrieved successfully",
	})
}

func CreatePayment(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Payment created successfully",
	})
}

func ShowPayment(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Payment retrieved successfully",
	})
}

func UpdatePayment(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Payment updated successfully",
	})
}

func DeletePayment(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Payment deleted successfully",
	})
}
