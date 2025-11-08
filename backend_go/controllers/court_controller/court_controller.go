package courtcontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rahmat-ama/badminton_reservation/db"
	"github.com/rahmat-ama/badminton_reservation/dto/court_dto"
	"github.com/rahmat-ama/badminton_reservation/models"
	"github.com/rahmat-ama/badminton_reservation/utils"
)

func Index(ctx *gin.Context) {
	var courts []models.Court
	if err := db.GetDB().Order("updated_at desc").Find(&courts).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal mengambil data court"))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Berhasil mengambil data court", courts))
}

func Create(ctx *gin.Context) {
	var req court_dto.CreateCourtRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Data tidak valid: "+err.Error()))
		return
	}

	var existCourt models.Court
	if err := db.GetDB().Where("court_name = ?", req.CourtName).First(&existCourt).Error; err == nil {
		ctx.JSON(http.StatusConflict, utils.ErrorResponse("Court dengan nama tersebut sudah ada"))
		return
	}

	court := models.Court{
		CourtName: req.CourtName,
		Type:      req.Type,
		Location:  req.Location,
	}

	if err := db.GetDB().Create(&court).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal membuat court"))
		return
	}

	ctx.JSON(http.StatusCreated, utils.SuccessResponse("Court berhasil dibuat", court))
}

func Show(ctx *gin.Context) {
	id := ctx.Param("id")
	var court models.Court
	if err := db.GetDB().First(&court, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Court tidak ditemukan"))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Berhasil mengambil data court", court))
}

func Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var req court_dto.UpdateCourtRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Data tidak valid: "+err.Error()))
		return
	}

	var court models.Court
	if err := db.GetDB().First(&court, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Court tidak ditemukan"))
		return
	}

	if req.CourtName != "" {
		var existCourt models.Court
		courtID, _ := strconv.ParseUint(id, 10, 32)
		if err := db.GetDB().Where("court_name = ? AND id != ?", req.CourtName, courtID).First(&existCourt).Error; err == nil {
			ctx.JSON(http.StatusConflict, utils.ErrorResponse("Court dengan nama tersebut sudah ada"))
			return
		}
		court.CourtName = req.CourtName
	}
	if req.Type != "" {
		court.Type = req.Type
	}
	if req.Location != "" {
		court.Location = req.Location
	}

	if err := db.GetDB().Save(&court).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal mengupdate court"))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Court berhasil diupdate", court))
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var court models.Court
	if err := db.GetDB().First(&court, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Court tidak ditemukan"))
		return
	}

	var bookingCount int64
	db.GetDB().Model(&models.Booking{}).Where("court_id = ? AND status != ?", id, "cancelled").Count(&bookingCount)

	if bookingCount > 0 {
		ctx.JSON(http.StatusConflict, utils.ErrorResponse("Tidak dapat menghapus court yang memiliki booking aktif"))
		return
	}

	if err := db.GetDB().Delete(&court).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal menghapus court"))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Court berhasil dihapus", nil))
}
