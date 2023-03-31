package controller

import (
	"net/http"
	"strconv"

	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
)

type SeederController interface {
	GetAllBank(ctx *gin.Context)
	GetBankByID(ctx *gin.Context)
	GetAllCategory(ctx *gin.Context)
	GetCategoryByID(ctx *gin.Context)
	GetAllStatusPembayaran(ctx *gin.Context)
	GetStatusPembayaranByID(ctx *gin.Context)
}

type seederController struct {
	seederService services.SeederService
}

func NewSeederController(ss services.SeederService) SeederController{
	return &seederController{
		seederService: ss, 
	}
}

func (sc *seederController) GetAllBank(ctx *gin.Context) {
	banks, err := sc.seederService.GetAllBank(ctx)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan List Bank", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan List Bank", banks)
	ctx.JSON(http.StatusOK, res)
}

func (sc *seederController) GetBankByID(ctx *gin.Context){
	id := ctx.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := sc.seederService.GetBankByID(ctx, uint(uintID))
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Bank", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Bank", result)
	ctx.JSON(http.StatusOK, res)
}

func (sc *seederController) GetAllCategory(ctx *gin.Context) {
	categories, err := sc.seederService.GetAllCategory(ctx)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan List Category", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan List Category", categories)
	ctx.JSON(http.StatusOK, res)
}

func (sc *seederController) GetCategoryByID(ctx *gin.Context){
	id := ctx.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := sc.seederService.GetCategoryByID(ctx, uint(uintID))
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Category", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Category", result)
	ctx.JSON(http.StatusOK, res)
}

func (sc *seederController) GetAllStatusPembayaran(ctx *gin.Context) {
	statusPembayaran, err := sc.seederService.GetAllStatusPembayaran(ctx)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan List Status Pembayaran", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan List Status Pembayaran", statusPembayaran)
	ctx.JSON(http.StatusOK, res)
}

func (sc *seederController) GetStatusPembayaranByID(ctx *gin.Context){
	id := ctx.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := sc.seederService.GetStatusPembayaranByID(ctx, uint(uintID))
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Status Pembayaran", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Status Pembayaran", result)
	ctx.JSON(http.StatusOK, res)
}