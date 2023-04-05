package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PenarikanController interface {
	CreatePenarikan(ctx *gin.Context)
	GetPenarikanByUser(ctx *gin.Context)
}

type penarikanController struct {
	jwtService services.JWTService
	userService services.UserService
	penarikanService services.PenarikanService
	eventService services.EventService
	db *gorm.DB
}

func NewPenarikanController(us services.UserService, es services.EventService, ps services.PenarikanService, db *gorm.DB, jwt services.JWTService) PenarikanController {
	return &penarikanController{
		jwtService: jwt,
		userService: us,
		penarikanService: ps,
		eventService: es,
		db: db,
	}
}

func (pc *penarikanController) CreatePenarikan(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := pc.jwtService.GetUserIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	var penarikanDTO dto.PenarikanEventDTO
	if err := ctx.ShouldBind(&penarikanDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	event, err := pc.eventService.GetEventByID(ctx.Request.Context(), penarikanDTO.EventID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Event", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	balance := event.JumlahDonasi
	if balance < penarikanDTO.Jumlah_Penarikan {
		res := utils.BuildResponseFailed("Saldo Tidak Mencukupi", "Saldo anda tidak mencukupi untuk melakukan penarikan", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	
	var listBank entities.ListBank
	if err := pc.db.Where("id = ?", penarikanDTO.BankID).First(&listBank).Error; err != nil {
		res := utils.BuildResponseFailed("ID Bank Tidak Ditemukan", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	penarikan := entities.HistoryPenarikan{
		Jumlah_Penarikan: penarikanDTO.Jumlah_Penarikan,
		NamaBank: listBank.Nama, 
		Tanggal_Penarikan: time.Now(),
		BankID: penarikanDTO.BankID,
		EventID: penarikanDTO.EventID,
		UserID: userID,
	}

	result, err := pc.penarikanService.CreatePenarikan(ctx.Request.Context(), penarikan)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan Penarikan", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// newBalance := balance - penarikanDTO.Jumlah_Penarikan
	// if newBalance <= 0 {
	// 	event.SisaDonasi = 0
	// } else{
	// 	event.SisaDonasi = newBalance
	// }

	// eventDTO := dto.EventUpdateDTO {
	// 	JumlahDonasi: &event.JumlahDonasi,
	// }

	// err = pc.eventService.UpdateEvent(ctx, eventDTO, penarikan.EventID)
	// if err != nil {
	// 	res := utils.BuildResponseFailed("Gagal Mengupdate Saldo Event", err.Error(), utils.EmptyObj{})
	// 	ctx.JSON(http.StatusBadRequest, res)
	// 	return
	// }

	res := utils.BuildResponseSuccess("Penarikan Berhasil", result)
	ctx.JSON(http.StatusOK, res)
}

func (pc *penarikanController) GetPenarikanByUser(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := pc.jwtService.GetUserIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	result, err := pc.penarikanService.GetPenarikanByUser(ctx, userID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan History Penarikan User", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		fmt.Print(res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan History Penarikan User", result)
	ctx.JSON(http.StatusOK, res)
}