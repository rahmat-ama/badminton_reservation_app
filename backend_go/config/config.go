package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	AppName    string
	Port       string
	JWT_SECRET string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Kesalahan memuat file .env atau file .env tidak ditemukan")
	}

	AppName = getEnv("APP_NAME", "Badminton Reservation App")
	Port = getEnv("APP_PORT", ":8000")
	JWT_SECRET = getEnv("JWT_SECRET", "7a9f3b8c2e1d4f6a5b9c8e2f1d3a7b9c4e6f8a1b3c5d7e9f2a4b6c8d0e1f3a5b7")
	DBHost = getEnv("DB_HOST", "db")
	DBPort = getEnv("DB_PORT", "5432")
	DBUser = getEnv("DB_USER", "postgres")
	DBPassword = getEnv("DB_PASSWORD", "postgres")
	DBName = getEnv("DB_NAME", "badminton_reservation_db")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("Variabel %s kosong, menggunakan nilai default: %s", key, defaultValue)
		return defaultValue
	}
	return value
}
