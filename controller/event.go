package controller

import (
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventController interface {
	CreateEvent(ctx *gin.Context)
	GetAllEvent(ctx *gin.Context)
	GetAllEventByUserID(ctx *gin.Context)
	GetEventByID(ctx *gin.Context)
	LikeEventByEventID(ctx *gin.Context)
	UpdateEvent(ctx *gin.Context)
	DeleteEvent(ctx *gin.Context)
	GetAllEventLastTransaksi(ctx *gin.Context)
	Get3Event(ctx *gin.Context)
}

type eventController struct {
	jwtService       services.JWTService
	eventService     services.EventService
	transaksiService services.TransaksiService
	db               *gorm.DB
	page             uint
}

func NewEventController(es services.EventService, ts services.TransaksiService, jwt services.JWTService, db *gorm.DB) EventController {
	return &eventController{
		jwtService:       jwt,
		eventService:     es,
		transaksiService: ts,
		db:               db,
		page:             1,
	}
}

func (ec *eventController) CreateEvent(ctx *gin.Context) {
	var eventDTO dto.EventCreateDTO
	if err := ctx.ShouldBind(&eventDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Check if the category event exists
	var category entities.CategoryEvent
	if err := ec.db.Where("nama = ?", eventDTO.JenisEvent).First(&category).Error; err != nil {
		res := utils.BuildResponseFailed("Kategori Event Tidak Ditemukan", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}


	timeLeft := TimeLeft(eventDTO.ExpiredDonasi)
	eventDTO.SisaHariDonasi = &timeLeft
	// Create the event
	eventDTO.JenisEvent = category.Nama
	event, err := ec.eventService.CreateEvent(ctx, eventDTO)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan Event", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menambahkan Event", event)
	ctx.JSON(http.StatusOK, res)
}

func TimeLeft(expiredTime time.Time) string {
	timeLeft := expiredTime.Sub(time.Now())

	fmt.Print(timeLeft.Hours() / 24)
	if timeLeft.Hours() / 24 <= 0 {
		return "Waktu Habis"
	}

	if timeLeft.Hours() / 24 < 1 {
		return "<1"
	}
	
	dayLeft := int(math.Round(timeLeft.Hours() / 24))
	return fmt.Sprintf("%v", dayLeft)
}

func (ec *eventController) GetAllEvent(ctx *gin.Context) {
	events, err := ec.eventService.GetAllEvent(ctx)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan List Event", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan List Event", events)
	ctx.JSON(http.StatusOK, res)
}

func (ec *eventController) GetAllEventByUserID(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	uuid, err := uuid.Parse(userID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := ec.eventService.GetAllEventByUserID(ctx, uuid)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Event", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Event", result)
	ctx.JSON(http.StatusOK, res)
}

func (ec *eventController) GetEventByID(ctx *gin.Context) {
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := ec.eventService.GetEventByID(ctx, uuid)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Event", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Event", result)
	ctx.JSON(http.StatusOK, res)
}

func (ec *eventController) LikeEventByEventID(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	event_id := ctx.Param("event_id")

	user_uuid, err := uuid.Parse(user_id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	event_uuid, err := uuid.Parse(event_id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err = ec.eventService.LikeEventByEventID(ctx, user_uuid, event_uuid); err != nil {
		res := utils.BuildResponseFailed("Gagal Like Event", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Like Event", utils.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (ec *eventController) UpdateEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	var eventDTO dto.EventUpdateDTO
	if err := ctx.ShouldBindJSON(&eventDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	eventDTO.ID = uuid
	if err := ec.eventService.UpdateEvent(ctx, eventDTO, uuid); err != nil {
		res := utils.BuildResponseFailed("Gagal Mengupdate Event", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Update Event", eventDTO)
	ctx.JSON(http.StatusOK, res)
}

func (ec *eventController) DeleteEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := ec.eventService.DeleteEvent(ctx, uuid); err != nil {
		res := utils.BuildResponseFailed("Gagal Delete Event", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Delete Event", utils.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (ec *eventController) GetAllEventLastTransaksi(ctx *gin.Context) {
	id := ctx.Param("event_id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	events, err := ec.transaksiService.GetAllEventLastTransaksi(ctx, uuid)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan List Transaksi", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan List Transaksi", events)
	ctx.JSON(http.StatusOK, res)
}

func (ec *eventController) Get3Event(ctx *gin.Context) {
	limit := ec.page * 3

	events, err := ec.eventService.Get3Event(ctx, limit)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan List Event", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ec.page++
	res := utils.BuildResponseSuccess("Berhasil Mendapatkan List Event", events)
	ctx.JSON(http.StatusOK, res)
}
