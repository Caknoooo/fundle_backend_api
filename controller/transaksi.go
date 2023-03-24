package controller

import (
	"net/http"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
)

type TransaksiController interface {
	CreateTransaksi(ctx *gin.Context)
	CreateTransaksiUser(ctx *gin.Context)
	GetAllTransaksi(ctx *gin.Context)
}

type transaksiController struct {
	transaksiService services.TransaksiService
}

func NewTransaksiController(us services.TransaksiService) TransaksiController {
	return &transaksiController{
		transaksiService: us,
	}
}

func (tc *transaksiController) CreateTransaksi(ctx *gin.Context) {
	var transaksi dto.TransaksiCreateDTO
	if err := ctx.ShouldBind(&transaksi); err != nil {
		panic(err) // harus diperbaiki
	}

	result, err := tc.transaksiService.CreateTransaksi(ctx.Request.Context(), transaksi)

	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan User", "Failed", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menambahkan User", result)
	ctx.JSON(http.StatusOK, res)
}

func (tc *transaksiController) CreateTransaksiUser(ctx *gin.Context) {
	var transaksi dto.TransaksiUserCreateDTO
	if err := ctx.ShouldBind(&transaksi); err != nil {
		panic(err) // harus diperbaiki
	}

	result, err := tc.transaksiService.CreateTransaksiUser(ctx.Request.Context(), transaksi)

	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan User", "Failed", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menambahkan User", result)
	ctx.JSON(http.StatusOK, res)
}

func (tc *transaksiController) GetAllTransaksi(ctx *gin.Context) {
	result, err := tc.transaksiService.GetAllTransaksi(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan List User", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("Berhasil Mendapatkan List User", result)
	ctx.JSON(http.StatusOK, res)
}
