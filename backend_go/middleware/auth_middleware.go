package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rahmat-ama/badminton_reservation/db"
	"github.com/rahmat-ama/badminton_reservation/models"
	"github.com/rahmat-ama/badminton_reservation/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Header Authorization kosong",
			})
			ctx.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Format token tidak valid. Gunakan: Bearer <token>",
			})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Token tidak valid atau expired",
				"error":   err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Set("username", claims.Username)
		ctx.Set("email", claims.Email)
		ctx.Set("role_id", claims.RoleID)

		var role models.Role
		if err := db.GetDB().First(&role, claims.RoleID).Error; err == nil {
			ctx.Set("role_name", role.Name)
		}

		ctx.Next()
	}
}

func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roleNameInterface, exists := ctx.Get("role_name")
		if !exists {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":  "error",
				"message": "Role tidak ditemukan dalam token",
			})
			ctx.Abort()
			return
		}

		roleName, ok := roleNameInterface.(string)
		if !ok {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":  "error",
				"message": "Format role tidak valid",
			})
			ctx.Abort()
			return
		}

		isAllowed := false
		for _, allowedRole := range allowedRoles {
			if strings.EqualFold(roleName, allowedRole) {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":        "error",
				"message":       "Akses ditolak. Role Anda tidak memiliki izin untuk aksi ini",
				"required_role": allowedRoles,
				"your_role":     roleName,
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return RequireRole("Admin")
}

func CustomerOnly() gin.HandlerFunc {
	return RequireRole("Customer")
}

func AdminOrCustomer() gin.HandlerFunc {
	return RequireRole("Admin", "Customer")
}
