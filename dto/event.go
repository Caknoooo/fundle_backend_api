package dto

import (
	"time"

	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/type/decimal"
)

type EventCreateDTO struct {
	RekeningEvent  string          `json:"rekening_event" form:"rekening_event" binding:"required"`
	Nama           string          `json:"nama" form:"nama" binding:"required"`
	DeskripsiEvent string          `json:"deskripsi_event" form:"deskripsi_event" binding:"required"`
	JenisEvent     string          `json:"jenis_event" form:"jenis_event" binding:"required"`
	JumlahDonasi   decimal.Decimal `json:"jumlah_donasi" form:"jumlah_donasi" binding:"required"`
	FotoEvent      string          `json:"foto_event" form:"foto_event" binding:"required"`
	ExpiredDonasi  time.Time       `json:"expired_donasi" form:"expired_donasi" binding:"required"`
	UserID         uuid.UUID       `json:"user_id" form:"user_id" binding:"required"`
}

type EventUpdateDTO struct {
	RekeningEvent  *string          `json:"rekening_event" form:"rekening_event"`
	Nama           *string          `json:"nama" form:"nama"`
	DeskripsiEvent *string          `json:"deskripsi_event" form:"deskripsi_event"`
	JenisEvent     *string          `json:"jenis_event" form:"jenis_event"`
	JumlahDonasi   *decimal.Decimal `json:"jumlah_donasi" form:"jumlah_donasi"`
	FotoEvent      *string          `json:"foto_event" form:"foto_event"` // Ini kalau mau dibuat banyak foto, harus one to many
	UserID         *string          `json:"user_id" form:"user_id"`
	IsTargetFull   *bool            `json:"is_target_full" form:"is_target_full"`
	IsExpired      *bool            `json:"is_expired" form:"is_expired"`
}