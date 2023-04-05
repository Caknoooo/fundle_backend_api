package controller

import (
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

type UserController interface {
	RegisterUser(ctx *gin.Context)
	GetAllUser(ctx *gin.Context)
	MeUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
	LogoutUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	CreateTransaksiUser(ctx *gin.Context)
	GetTransaksiUser(ctx *gin.Context)
}

type userController struct {
	jwtService        services.JWTService
	userService       services.UserService
	transaksiService  services.TransaksiService
	pembayaranService services.PembayaranService
	eventService      services.EventService
	db                *gorm.DB
}

func NewUserController(us services.UserService, ts services.TransaksiService, ps services.PembayaranService, es services.EventService, db *gorm.DB, jwt services.JWTService) UserController {
	return &userController{
		jwtService:        jwt,
		userService:       us,
		transaksiService:  ts,
		pembayaranService: ps,
		eventService:      es,
		db:                db,
	}
}

func (uc *userController) RegisterUser(ctx *gin.Context) {
	var user dto.UserCreateDTO
	if err := ctx.ShouldBind(&user); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if checkUser, _ := uc.userService.CheckUser(ctx.Request.Context(), user.Email); checkUser {
		res := utils.BuildResponseFailed("Email Sudah Terdaftar", "Failed", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, err := uc.userService.RegisterUser(ctx.Request.Context(), user)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan User", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menambahkan User", result)
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) GetAllUser(ctx *gin.Context) {
	result, err := uc.userService.GetAllUser(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan List User", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("Berhasil Mendapatkan List User", result)
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) MeUser(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetUserIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	result, err := uc.userService.GetUserByID(ctx.Request.Context(), userID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan User", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan User", result)
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) LoginUser(ctx *gin.Context) {
	var userLoginDTO dto.UserLoginDTO
	err := ctx.ShouldBind(&userLoginDTO)
	res, _ := uc.userService.Verify(ctx.Request.Context(), userLoginDTO.Email, userLoginDTO.Password)
	if !res {
		response := utils.BuildResponseFailed("Gagal Login", "Email atau Password Salah", utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, err := uc.userService.GetUserByEmail(ctx.Request.Context(), userLoginDTO.Email)
	if err != nil {
		response := utils.BuildResponseFailed("Gagal Login", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	token := uc.jwtService.GenerateToken(user.ID, user.Role)
	userResponse := entities.Authorization{
		Token:     token,
		Role:      user.Role,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	// ctx.SetCookie("token", token, 60*60*24, "/", "localhost", false, true)
	// ctx.String(http.StatusOK, "Token saved")
	response := utils.BuildResponseSuccess("Berhasil Login", userResponse)
	ctx.JSON(http.StatusOK, response)
}

func (uc *userController) LogoutUser(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		res := utils.BuildResponseFailed("Gagal Logout", "Token Kosong", utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err := uc.jwtService.InvalidateToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Logout", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	ctx.Header("Set-Cookie", "token=; Path=/; Max-Age=-1")
	ctx.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")

	res := utils.BuildResponseSuccess("Berhasil Logout", utils.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) UpdateUser(ctx *gin.Context) {
	var userDTO dto.UserUpdateDTO
	if err := ctx.ShouldBind(&userDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetUserIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userDTO.ID = userID
	if err = uc.userService.UpdateUser(ctx.Request.Context(), userDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Update User", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Update User", userDTO)
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) DeleteUser(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetUserIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if err = uc.userService.DeleteUser(ctx.Request.Context(), userID); err != nil {
		res := utils.BuildResponseFailed("Gagal Menghapus User", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menghapus User", utils.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) CreateTransaksiUser(ctx *gin.Context) {
	// Mendapatkan user ID dari token yang di-passing melalui context
	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetUserIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	// Mendapatkan data pembayaran dari request body
	var pembayaran dto.PembayaranDTO
	if err := ctx.ShouldBind(&pembayaran); err != nil {
		res := utils.BuildResponseFailed("Gagal Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var listBank entities.ListBank
	if err := uc.db.Where("id = ?", pembayaran.ListBankID).First(&listBank).Error; err != nil {
		res := utils.BuildResponseFailed("ID Bank Tidak Ditemukan", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Membuat pembayaran baru dan mendapatkan ID pembayaran
	resultPembayaran, err := uc.pembayaranService.CreatePembayaran(ctx.Request.Context(), pembayaran)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan Pembayaran", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Mendapatkan ID event dari path parameter
	eventID, err := uuid.Parse(ctx.Param("event_id"))
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Membuat transaksi baru dan mendapatkan ID transaksi
	transaksi := dto.TransaksiCreateDTO{
		NamaBank:            listBank.Nama,
		Jumlah_Donasi_Event: resultPembayaran.Jumlah,
		Tanggal_Transaksi:   time.Now(),
		EventID:             eventID,
		PembayaranID:        resultPembayaran.ID,
		UserID:              userID,
	}

	resultTransaksi, err := uc.transaksiService.CreateTransaksi(ctx.Request.Context(), transaksi)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan Transaksi", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Mendapatkan data event berdasarkan ID event
	event, err := uc.eventService.GetEventByID(ctx.Request.Context(), eventID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Event", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if event.Is_target_full {
		res := utils.BuildResponseFailed("Jumlah Donasi Telah Penuh", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	sisa_donasi := event.SisaDonasi + resultPembayaran.Jumlah
	
	// Menghitung jumlah donasi baru yang akan ditambahkan ke event
	newJumlahDonasi := event.JumlahDonasi + resultTransaksi.Jumlah_Donasi_Event
	if newJumlahDonasi >= event.MaxDonasi {
		newJumlahDonasi = event.MaxDonasi
		eventDTO := dto.EventUpdateDTO{
			JumlahDonasi: &newJumlahDonasi,
		}

		err = uc.eventService.PatchEvent(ctx.Request.Context(), eventDTO, eventID)
		if err != nil {
			res := utils.BuildResponseFailed("Gagal Mengupdate Jumlah Donasi Event", err.Error(), utils.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, res)
			return
		}

		res := utils.BuildResponseSuccess("Donasi Mencapai Batas Maksimum", event.MaxDonasi)
		ctx.JSON(http.StatusOK, res)
		return
	}

	eventDTO := dto.EventUpdateDTO{
		JumlahDonasi: &newJumlahDonasi,
		SisaDonasi: &sisa_donasi,
	}

	err = uc.eventService.PatchEvent(ctx.Request.Context(), eventDTO, eventID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mengupdate Jumlah Donasi Event", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res1 := utils.BuildResponseSuccess("Berhasil Menambahkan Transaksi", resultTransaksi)
	ctx.JSON(http.StatusOK, res1)
}

func (uc *userController) GetTransaksiUser(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetUserIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	result, err := uc.transaksiService.GetAllTransaksiByUserID(ctx.Request.Context(), userID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan User", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan User", result)
	ctx.JSON(http.StatusOK, res)
}
