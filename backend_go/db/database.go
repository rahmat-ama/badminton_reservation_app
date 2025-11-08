package db

import (
	"fmt"
	"log"
	"time"

	"github.com/rahmat-ama/badminton_reservation/config"
	"github.com/rahmat-ama/badminton_reservation/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Gagal mendapatkan instance database:", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Koneksi database berhasil!")

	if err := DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Court{},
		&models.Timeslot{},
		&models.Booking{},
	); err != nil {
		log.Fatal("Gagal migrate database:", err)
	}

	log.Println("Migrasi databse selesai!")
}

func GetDB() *gorm.DB {
	return DB
}
