package bookingcontroller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rahmat-ama/badminton_reservation/db"
	"github.com/rahmat-ama/badminton_reservation/dto/booking_dto"
	"github.com/rahmat-ama/badminton_reservation/models"
	"github.com/rahmat-ama/badminton_reservation/utils"
)

func Get(ctx *gin.Context) {
	var bookings []models.Booking
	roleName := ctx.GetString("role_name")
	userID := ctx.GetUint("user_id")
	query := db.GetDB().Preload("User").Preload("Court").Preload("Timeslot").Order("booking_date DESC, created_at DESC")

	if roleName != "Admin" {
		query = query.Where("user_id = ?", userID)
	}

	if err := query.Find(&bookings).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal memuat data booking"))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Berhasil memuat data booking", bookings))
}

func Create(ctx *gin.Context) {
	var req booking_dto.CreateBookingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Data tidak valid: "+err.Error()))
		return
	}

	var existBooking models.Booking
	bookingDate, err := time.Parse("2006-01-02", req.BookingDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Format tanggal tidak valid: "+err.Error()))
		return
	}

	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse("User tidak terautentikasi"))
		return
	}

	result := db.GetDB().Where("court_id = ? AND timeslot_id = ? AND booking_date = ? AND status != ?",
		req.CourtID, req.TimeslotID, bookingDate, "cancelled").
		First(&existBooking)

	if result.Error == nil {
		ctx.JSON(http.StatusConflict, utils.ErrorResponse("Booking untuk court dan timeslot ini pada tanggal tersebut sudah ada"))
		return
	}

	var court models.Court
	if err := db.GetDB().First(&court, req.CourtID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Court tidak ditemukan"))
		return
	}

	var timeslot models.Timeslot
	if err := db.GetDB().First(&timeslot, req.TimeslotID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Timeslot tidak ditemukan"))
		return
	}

	totalPrice := timeslot.PriceWeekday
	weekday := bookingDate.Weekday()
	if weekday == time.Saturday || weekday == time.Sunday {
		totalPrice = timeslot.PriceWeekend
	}

	booking := models.Booking{
		UserID:      userID.(uint),
		CourtID:     req.CourtID,
		TimeslotID:  req.TimeslotID,
		BookingDate: bookingDate,
		Status:      "pending",
		TotalPrice:  totalPrice,
	}

	if err := db.GetDB().Create(&booking).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal membuat booking"))
		return
	}
	db.GetDB().Preload("User").Preload("Court").Preload("Timeslot").First(&booking, booking.ID)

	ctx.JSON(http.StatusCreated, utils.SuccessResponse("Booking berhasil dibuat", booking))
}

func Show(ctx *gin.Context) {
	id := ctx.Param("id")
	var booking models.Booking
	if err := db.GetDB().Preload("User").Preload("Court").Preload("Timeslot").First(&booking, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Booking tidak ditemukan"))
		return
	}

	roleName := ctx.GetString("role_name")
	userID := ctx.GetUint("user_id")

	if roleName != "Admin" && booking.UserID != userID {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse("Anda hanya dapat melihat booking milik sendiri"))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Berhasil memuat data booking", booking))
}

func Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var req booking_dto.UpdateBookingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Data tidak valid: "+err.Error()))
		return
	}

	var booking models.Booking
	if err := db.GetDB().First(&booking, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Booking tidak ditemukan"))
		return
	}

	roleName := ctx.GetString("role_name")
	userID := ctx.GetUint("user_id")
	if roleName != "Admin" && booking.UserID != userID {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse("Anda hanya dapat mengupdate booking milik sendiri"))
		return
	}

	if req.Status != "" {
		booking.Status = req.Status
	}

	if req.PaymentToken != "" {
		booking.PaymentToken = req.PaymentToken
	}

	if err := db.GetDB().Save(&booking).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal memperbarui booking"))
		return
	}
	db.GetDB().Preload("User").Preload("Court").Preload("Timeslot").First(&booking, id)

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Booking berhasil diperbarui", booking))
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var booking models.Booking
	if err := db.GetDB().First(&booking, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Booking tidak ditemukan"))
		return
	}
	if err := db.GetDB().Delete(&booking).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal menghapus booking: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Booking berhasil dihapus", nil))
}
