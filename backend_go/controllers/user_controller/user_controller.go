package usercontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	authdto "github.com/rahmat-ama/badminton_reservation/dto/auth_dto"
	userdto "github.com/rahmat-ama/badminton_reservation/dto/user_dto"
	authservice "github.com/rahmat-ama/badminton_reservation/services/auth_service"
	userservice "github.com/rahmat-ama/badminton_reservation/services/user_service"
	"github.com/rahmat-ama/badminton_reservation/utils"
)

func RegisterUser(ctx *gin.Context) {
	var req authdto.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Request tidak valid: "+err.Error()))
		return
	}

	response, err := authservice.Register(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Request tidak valid: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.SuccessResponse("Registrasi berhasil, Anda sudah login", response))
}

func LoginUser(ctx *gin.Context) {
	var req authdto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Request tidak valid: "+err.Error()))
		return
	}

	response, err := authservice.Login(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Request tidak valid: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Login berhasil", response))
}

func Get(ctx *gin.Context) {
	users, err := userservice.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Request tidak valid: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse("Data user berhasil dimuat", users))
}

func Show(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Request tidak valid: "+err.Error()))
		return
	}

	userID := ctx.GetUint("user_id")
	roleName := ctx.GetString("role_name")

	if roleName != "Admin" && userID != uint(id) {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status":  "error",
			"message": "Anda hanya dapat melihat profile sendiri",
		})
		return
	}

	user, err := userservice.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Request tidak valid: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("User berhasil dimuat", user))
}

func Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Request tidak valid: "+err.Error()))
		return
	}

	userID := ctx.GetUint("user_id")
	roleName := ctx.GetString("role_name")

	if roleName != "Admin" && userID != uint(id) {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status":  "error",
			"message": "Anda hanya dapat mengubah profile sendiri",
		})
		return
	}

	var req userdto.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Request tidak valid: "+err.Error()))
		return
	}

	user, err := userservice.UpdateUser(uint(id), &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Request tidak valid: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("User berhasil diperbarui", user))
}

func Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Request tidak valid: "+err.Error()))
		return
	}

	if err := userservice.DeleteUser(uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Request tidak valid: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("User berhasil dihapus", nil))
}
