package controller

import (
	"net/http"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransaksiController interface {
	CreateTransaksi(ctx *gin.Context)
	GetAllTransaksi(ctx *gin.Context)
	GetTransaksiByID(ctx *gin.Context)
	GetTransaksiByUserID(ctx *gin.Context)
}

type transaksiController struct {
	jwtService       services.JWTService
	transaksiService services.TransaksiService
}

func NewTransaksiController(us services.TransaksiService, jwt services.JWTService) TransaksiController {
	return &transaksiController{
		transaksiService: us,
		jwtService:       jwt,
	}
}

func (tc *transaksiController) CreateTransaksi(ctx *gin.Context) {
	var transaksiDTO dto.TransaksiCreateDTO
	if err := ctx.ShouldBind(&transaksiDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	transaksi, err := tc.transaksiService.CreateTransaksi(ctx.Request.Context(), transaksiDTO)

	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan User", "Failed", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menambahkan Transaksi", transaksi)
	ctx.JSON(http.StatusOK, res)
}

func (tc *transaksiController) GetAllTransaksi(ctx *gin.Context) {
	result, err := tc.transaksiService.GetAllTransaksi(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan List Transaksi", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("Berhasil Mendapatkan List Transaksi", result)
	ctx.JSON(http.StatusOK, res)
}

func (tc *transaksiController) GetTransaksiByID(ctx *gin.Context) {
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := tc.transaksiService.GetTransaksiByID(ctx, uuid)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Transaksi", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Transaksi", result)
	ctx.JSON(http.StatusOK, res)
}

func (tc *transaksiController) GetTransaksiByUserID(ctx *gin.Context) {
	id := ctx.Param("user_id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := tc.transaksiService.GetTransaksiByID(ctx, uuid)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan List Transaksi", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan List Transaksi", result)
	ctx.JSON(http.StatusOK, res)
}
