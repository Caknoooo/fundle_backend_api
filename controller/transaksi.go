package controller

import (
	"net/http"

	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
)

type TransaksiController interface {
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
