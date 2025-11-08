package timeslotcontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rahmat-ama/badminton_reservation/db"
	"github.com/rahmat-ama/badminton_reservation/dto/timeslot_dto"
	"github.com/rahmat-ama/badminton_reservation/models"
	"github.com/rahmat-ama/badminton_reservation/utils"
)

func Index(ctx *gin.Context) {
	var timeslots []models.Timeslot
	if err := db.GetDB().Order("start_time ASC").Find(&timeslots).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal mengambil data timeslot"))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Berhasil mengambil data timeslot", timeslots))
}

func Create(ctx *gin.Context) {
	var req timeslot_dto.CreateTimeslotRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Data tidak valid: "+err.Error()))
		return
	}

	if req.PriceWeekday <= 0 || req.PriceWeekend <= 0 {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Harga harus lebih dari 0"))
		return
	}

	var existTimeslot models.Timeslot
	if err := db.GetDB().Where("start_time = ? AND end_time = ?", req.StartTime, req.EndTime).First(&existTimeslot).Error; err == nil {
		ctx.JSON(http.StatusConflict, utils.ErrorResponse("Timeslot dengan waktu tersebut sudah ada"))
		return
	}

	timeslot := models.Timeslot{
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
		PriceWeekday: req.PriceWeekday,
		PriceWeekend: req.PriceWeekend,
	}

	if err := db.GetDB().Create(&timeslot).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal membuat timeslot: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.SuccessResponse("Timeslot berhasil dibuat", timeslot))
}

func Show(ctx *gin.Context) {
	id := ctx.Param("id")
	var timeslot models.Timeslot
	if err := db.GetDB().First(&timeslot, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Timeslot tidak ditemukan"))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Berhasil mengambil data timeslot", timeslot))
}

func Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var req timeslot_dto.UpdateTimeslotRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Data tidak valid: "+err.Error()))
		return
	}

	var timeslot models.Timeslot
	if err := db.GetDB().First(&timeslot, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Timeslot tidak ditemukan"))
		return
	}

	if req.StartTime != "" {
		var existTimeslot models.Timeslot
		timeslotID, _ := strconv.ParseUint(id, 10, 32)
		endTime := req.EndTime
		if endTime == "" {
			endTime = timeslot.EndTime
		}
		if err := db.GetDB().Where("start_time = ? AND end_time = ? AND id != ?", req.StartTime, endTime, timeslotID).First(&existTimeslot).Error; err == nil {
			ctx.JSON(http.StatusConflict, utils.ErrorResponse("Timeslot dengan waktu tersebut sudah ada"))
			return
		}
		timeslot.StartTime = req.StartTime
	}
	if req.EndTime != "" {
		timeslot.EndTime = req.EndTime
	}
	if req.PriceWeekday > 0 {
		timeslot.PriceWeekday = req.PriceWeekday
	}
	if req.PriceWeekend > 0 {
		timeslot.PriceWeekend = req.PriceWeekend
	}

	if err := db.GetDB().Save(&timeslot).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal mengupdate timeslot"))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Timeslot berhasil diupdate", timeslot))
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var timeslot models.Timeslot
	if err := db.GetDB().First(&timeslot, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Timeslot tidak ditemukan"))
		return
	}

	var bookingCount int64
	db.GetDB().Model(&models.Booking{}).Where("timeslot_id = ? AND status != ?", id, "cancelled").Count(&bookingCount)

	if bookingCount > 0 {
		ctx.JSON(http.StatusConflict, utils.ErrorResponse("Tidak dapat menghapus timeslot yang memiliki booking aktif"))
		return
	}

	if err := db.GetDB().Delete(&timeslot).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal menghapus timeslot"))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Timeslot berhasil dihapus", nil))
}
