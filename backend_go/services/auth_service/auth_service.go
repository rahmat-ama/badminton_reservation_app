package authservice

import (
	"errors"
	"fmt"

	"github.com/rahmat-ama/badminton_reservation/db"
	authdto "github.com/rahmat-ama/badminton_reservation/dto/auth_dto"
	userdto "github.com/rahmat-ama/badminton_reservation/dto/user_dto"
	"github.com/rahmat-ama/badminton_reservation/models"
	userservice "github.com/rahmat-ama/badminton_reservation/services/user_service"
	"github.com/rahmat-ama/badminton_reservation/utils"
	"gorm.io/gorm"
)

func Register(req *authdto.RegisterRequest) (*authdto.LoginResponse, error) {
	var existUser models.User
	if err := db.GetDB().Where("username = ?", req.Username).First(&existUser).Error; err == nil {
		return nil, errors.New("username sudah terdaftar")
	}

	var role models.Role
	if err := db.GetDB().Where("name = ?", "Customer").First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			role = models.Role{Name: "Customer"}
			if err := db.GetDB().Create(&role).Error; err != nil {
				return nil, fmt.Errorf("gagal membuat default role: %w", err)
			}
		} else {
			return nil, fmt.Errorf("gagal mencari role: %w", err)
		}
	}

	createReq := &userdto.CreateUserRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Kontak:   req.Kontak,
		RoleID:   role.ID,
	}

	user, err := userservice.CreateUser(createReq)
	if err != nil {
		return nil, err
	}

	accessToken, err := utils.GenerateJWT(user.ID, user.Username, user.Email, user.RoleID)
	if err != nil {
		return nil, fmt.Errorf("gagal generate access token: %w", err)
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("gagal generate refresh token: %w", err)
	}

	db.GetDB().Preload("Role").First(user, user.ID)

	response := &authdto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    86400, // 24 jam
		User: authdto.UserInfo{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			RoleID:   user.RoleID,
			RoleName: user.Role.Name,
		},
	}

	return response, nil
}

func Login(req *authdto.LoginRequest) (*authdto.LoginResponse, error) {
	var user models.User
	if err := db.GetDB().Preload("Role").Where("username = ?", req.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("username atau password salah")
		}
		return nil, fmt.Errorf("gagal mencari user: %w", err)
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("username atau password salah")
	}

	accessToken, err := utils.GenerateJWT(user.ID, user.Username, user.Email, user.RoleID)
	if err != nil {
		return nil, fmt.Errorf("gagal generate access token: %w", err)
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("gagal generate refresh token: %w", err)
	}

	response := &authdto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    86400, // 24 jam
		User: authdto.UserInfo{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			RoleID:   user.RoleID,
			RoleName: user.Role.Name,
		},
	}

	return response, nil
}

func GetUserFromToken(userID uint) (*models.User, error) {
	var user models.User
	if err := db.GetDB().Preload("Role").First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mencari user: %w", err)
	}
	return &user, nil
}
