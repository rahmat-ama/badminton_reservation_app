package userservice

import (
	"errors"
	"fmt"

	"github.com/rahmat-ama/badminton_reservation/db"
	userdto "github.com/rahmat-ama/badminton_reservation/dto/user_dto"
	"github.com/rahmat-ama/badminton_reservation/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	err := db.GetDB().Preload("Role").Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("gagal memuat data user: %w", err)
	}
	return users, nil
}

func CreateUser(req *userdto.CreateUserRequest) (*models.User, error) {
	tx := db.GetDB().Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("gagal memulai transaction database: %w", tx.Error)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var role models.Role
	if err := tx.First(&role, req.RoleID).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("role dengan ID %d tidak ditemukan", req.RoleID)
		}
		return nil, fmt.Errorf("gagal validasi role: %w", err)
	}

	var Userexist models.User
	if err := tx.Where("username = ?", req.Username).First(&Userexist).Error; err == nil {
		tx.Rollback()
		return nil, errors.New("username sudah ada")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("gagal enkripsi password: %w", err)
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		Kontak:   req.Kontak,
		RoleID:   req.RoleID,
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("gagal membuat user: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("gagal commit trasaction database: %w", err)
	}

	db.GetDB().Preload("Role").First(&user, user.ID)

	return &user, nil
}

func GetByID(id uint) (*models.User, error) {
	var user models.User
	err := db.GetDB().Preload("Role").First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user dengan ID %d tidak ditemukan", id)
		}
		return nil, fmt.Errorf("gagal memuat data user: %w", err)
	}
	return &user, nil
}

func UpdateUser(id uint, req *userdto.UpdateUserRequest) (*models.User, error) {
	tx := db.GetDB().Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("gagal memulai transaksi: %w", tx.Error)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var user models.User
	if err := tx.First(&user, id).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user dengan ID %d tidak ditemukan", id)
		}
		return nil, fmt.Errorf("gagal memuat user: %w", err)
	}

	if req.RoleID > 0 {
		var role models.Role
		if err := tx.First(&role, req.RoleID).Error; err != nil {
			tx.Rollback()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("role dengan ID %d tidak ditemukan", req.RoleID)
			}
			return nil, fmt.Errorf("gagal validasi role: %w", err)
		}
		user.RoleID = req.RoleID
	}

	if req.Username != "" && req.Username != user.Username {
		var Userexist models.User
		if err := tx.Where("username = ? AND id != ?", req.Username, id).First(&Userexist).Error; err == nil {
			tx.Rollback()
			return nil, errors.New("username sudah ada")
		}
		user.Username = req.Username
	}

	if req.Kontak != "" {
		user.Kontak = req.Kontak
	}

	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("gagal hash password: %w", err)
		}
		user.Password = string(hashedPassword)
	}

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("gagal memperbarui user: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("gagal commit transaction database: %w", err)
	}

	db.GetDB().Preload("Role").First(&user, user.ID)

	return &user, nil
}

func DeleteUser(id uint) error {
	var user models.User
	if err := db.GetDB().First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user dengan ID %d tidak ditemukan", id)
		}
		return fmt.Errorf("gagal memuat user: %w", err)
	}

	if err := db.GetDB().Delete(&user).Error; err != nil {
		return fmt.Errorf("gagal hapus user: %w", err)
	}

	return nil
}
