package seed

import (
	"log"
	"time"

	"github.com/rahmat-ama/badminton_reservation/db"
	"github.com/rahmat-ama/badminton_reservation/models"
	"github.com/rahmat-ama/badminton_reservation/utils"
)

func SeedDB() {
	log.Println("Membuat roles...")
	adminRole := models.Role{Name: "Admin"}
	db.GetDB().FirstOrCreate(&adminRole, models.Role{Name: "Admin"})
	log.Printf("Role admin berhasil dibuat (ID: %d)", adminRole.ID)

	customerRole := models.Role{Name: "Customer"}
	db.GetDB().FirstOrCreate(&customerRole, models.Role{Name: "Customer"})
	log.Printf("Role customer berhasil dibuat (ID: %d)\n", customerRole.ID)

	log.Println("Membuat user dengan role admin...")
	hashedPassword, err := utils.HashPassword("admin123")
	if err != nil {
		log.Fatal("Gagal hashing password:", err)
	}
	adminUser := models.User{
		Username: "admin",
		Email:    "admin@badminton.com",
		Password: hashedPassword,
		Kontak:   "08123456789",
		RoleID:   adminRole.ID,
	}
	result := db.GetDB().Where("email = ?", "admin@badminton.com").FirstOrCreate(&adminUser)
	if result.Error != nil {
		log.Fatal("Gagal membuat user dengan role admin:", result.Error)
	}

	if result.RowsAffected > 0 {
		log.Println("Admin user berhasil dibuat!")
	} else {
		log.Println("Admin user sudah ada")
	}

	court1 := models.Court{CourtName: "Court 1", Type: "Indoor", Location: "Location 1"}
	court2 := models.Court{CourtName: "Court 2", Type: "Outdoor", Location: "Location 1"}
	court3 := models.Court{CourtName: "Court 3", Type: "Indoor", Location: "Location 1"}
	tanggal := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	timeslot1 := models.Timeslot{StartTime: tanggal + " 10:00:00 +0700", EndTime: tanggal + " 11:00:00 +0700",
		PriceWeekday: 50000, PriceWeekend: 70000}
	timeslot2 := models.Timeslot{StartTime: tanggal + " 13:00:00 +0700", EndTime: tanggal + " 15:00:00 +0700",
		PriceWeekday: 60000, PriceWeekend: 80000}
	timeslot3 := models.Timeslot{StartTime: tanggal + " 19:00:00 +0700", EndTime: tanggal + " 21:00:00 +0700",
		PriceWeekday: 55000, PriceWeekend: 70000}
	seedCourtOrTimeslot(court1)
	seedCourtOrTimeslot(court2)
	seedCourtOrTimeslot(court3)
	seedCourtOrTimeslot(timeslot1)
	seedCourtOrTimeslot(timeslot2)
	seedCourtOrTimeslot(timeslot3)

	log.Println("Kredensial Admin:")
	log.Println("	Username: admin")
	log.Println("	Email: admin@badminton.com")
	log.Println("	Password: admin123")
	log.Println("Start server dengan `go run cmd/server/main.go`")
}

func seedCourtOrTimeslot(data any) {
	switch model := data.(type) {
	case models.Court:
		result := db.DB.Where("court_name = ? AND type = ? AND location = ?", model.CourtName, model.Type, model.Location).
			FirstOrCreate(&model)
		if result.Error != nil {
			log.Fatal("Gagal membuat court")
		}
		if result.RowsAffected > 0 {
			log.Println("Berhasil membuat court")
		} else {
			log.Println("Court sudah ada")
		}
	case models.Timeslot:
		result := db.DB.Where("start_time = ? AND end_time = ?", model.StartTime, model.EndTime).
			FirstOrCreate(&model)
		if result.Error != nil {
			log.Fatal("Gagal membuat timeslot")
		}
		if result.RowsAffected > 0 {
			log.Println("Berhasil membuat timeslot")
		} else {
			log.Printf("Timeslot sudah ada")
		}
	}
}
