package controller

import (
	"net/http"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EventController interface {
	CreateEvent(ctx *gin.Context)
	GetAllEvent(ctx *gin.Context)
	GetAllEventByUserID(ctx *gin.Context)
	GetEventByID(ctx *gin.Context)
	LikeEventByEventID(ctx *gin.Context)
	UpdateEvent(ctx *gin.Context)
	DeleteEvent(ctx *gin.Context)
}

type eventController struct {
	jwtService   services.JWTService
	eventService services.EventService
}

func NewEventController(es services.EventService, jwt services.JWTService) EventController {
	return &eventController{
		jwtService:   jwt,
		eventService: es,
	}
}

func (ec *eventController) CreateEvent(ctx *gin.Context) {
	var eventDTO dto.EventCreateDTO
	if err := ctx.ShouldBind(&eventDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	event, err := ec.eventService.CreateEvent(ctx, eventDTO)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan Event", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menambahkan Event", event)
	ctx.JSON(http.StatusOK, res)
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
	if err := ec.eventService.UpdateEvent(ctx, eventDTO); err != nil {
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
