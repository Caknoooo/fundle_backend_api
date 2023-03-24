package dto

import (
	"time"

	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/type/decimal"
)

type TransaksiCreateDTO struct {
	ID                  uuid.UUID       `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Jumlah_Donasi_Event decimal.Decimal `gorm:"type:decimal(15,2)" json:"jumlah_donasi" binding:"required"`
	Tanggal_Transaksi   time.Time       `gorm:"datetime" json:"tangal_transaksi" binding:"required"`
	EventID             uuid.UUID       `gorm:"type:uuid" json:"event_id" form:"user_id" binding:"required"`
	PembayaranID        uuid.UUID       `gorm:"type:uuid" json:"pembayaran_id" binding:"required"`
}

type TransaksiUserCreateDTO struct {
	UserID      uuid.UUID `gorm:"type:uuid" json:"user_id" binding:"required"`
	TransaksiID uuid.UUID `gorm:"type:uuid" json:"transaksi_id" binding:"required"`
}
