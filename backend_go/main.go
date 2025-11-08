package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rahmat-ama/badminton_reservation/config"
	"github.com/rahmat-ama/badminton_reservation/db"
	"github.com/rahmat-ama/badminton_reservation/routes"
	"github.com/rahmat-ama/badminton_reservation/seed"
)

func main() {

	db.InitDB()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routes.SetupRoutes(router)

	seed.SeedDB()

	log.Printf("App start %s, port %s", config.AppName, config.Port)
	log.Printf("API Endpoint: http://localhost%s/api", config.Port)
	if err := router.Run(config.Port); err != nil {
		log.Fatal("Gagal memulai server:", err)
	}
}
