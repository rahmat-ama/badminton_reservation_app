package utils

import (
	"fmt"
	"time"

	"github.com/rahmat-ama/badminton_reservation/models"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func StartScheduler(db *gorm.DB) {
	// Set Timezone Jakarta agar Cron berjalan sesuai waktu WIB
	jakartaTime, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("Gagal load lokasi waktu, menggunakan UTC")
		jakartaTime = time.UTC
	}

	// Buat Cron instance dengan lokasi Jakarta
	c := cron.New(cron.WithLocation(jakartaTime))

	// --- JADWALKAN TUGAS ---
	// Format Cron: "Menit Jam TanggalBulan Bulan Hari"
	// "0 0 * * 1" artinya: Setiap Senin jam 00:00 malam

	_, err = c.AddFunc("0 0 * * 1", func() {
		fmt.Println("[CRON] Memulai Reset & Seed Jadwal Mingguan...")
		ResetAndSeedTimeSlots(db)
	})

	if err != nil {
		fmt.Println("Error menjadwalkan cron:", err)
		return
	}

	// Jalankan Cron di background
	c.Start()
	fmt.Println("[CRON] Scheduler Berjalan. Menunggu hari Senin pukul 00:00 WIB...")
}

func ResetAndSeedTimeSlots(db *gorm.DB) {
	// 1. HARD RESET: Hapus SEMUA data timecourt lama
	// Menggunakan Unscoped() untuk mengabaikan soft delete (hapus permanen)
	// Atau pakai TRUNCATE agar bersih total dan reset ID (opsional)
	err := db.Exec("TRUNCATE TABLE timeslots RESTART IDENTITY CASCADE").Error
	if err != nil {
		// Fallback jika Truncate gagal (misal karena constraint), pakai Delete biasa
		db.Unscoped().Where("1 = 1").Delete(&models.Timeslot{})
	}
	fmt.Println("[CRON] Data lama berhasil dihapus.")

	// 2. SEEDING BARU: Loop 7 hari ke depan mulai hari ini (Senin)
	startDate := time.Now()

	// Kita ingin Selasa (1), Rabu (2), Jumat (4), Sabtu (5) hari kedepan relatif dari Senin
	// Tapi logika loop lebih aman: Cek tiap hari

	for i := 0; i < 7; i++ {
		// Tambahkan i hari dari hari ini
		currentDate := startDate.AddDate(0, 0, i)
		weekday := currentDate.Weekday()

		// Filter Hari: Selasa, Rabu, Jumat, Sabtu
		if weekday == time.Tuesday || weekday == time.Wednesday || weekday == time.Friday || weekday == time.Saturday {

			// Generate 3 Slot Waktu: 09:00, 13:00, 19:00
			hours := []int{9, 13, 19}

			for _, h := range hours {
				// Atur jam mulai
				startTime := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), h, 0, 0, 0, time.Local)
				// Atur jam selesai (durasi 2 jam)
				endTime := startTime.Add(2 * time.Hour)

				// Create Object
				slot := models.Timeslot{
					StartTime:    startTime.String()[:25],
					EndTime:      endTime.String()[:25],
					PriceWeekday: 50000,
					PriceWeekend: 70000,
				}

				// Simpan ke DB
				if err := db.Create(&slot).Error; err != nil {
					fmt.Printf("Gagal buat slot %v: %v\n", startTime, err)
				}
			}
		}
	}
	fmt.Println("[CRON] Selesai! Jadwal minggu ini telah dibuat.")
}
