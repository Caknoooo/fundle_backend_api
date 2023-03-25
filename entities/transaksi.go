package entities

import (
	"time"

	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/type/decimal"
)

type Transaksi struct {
	ID                  uuid.UUID       `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Jumlah_Donasi_Event decimal.Decimal `gorm:"type:decimal(15,2)" json:"jumlah_donasi"`
	Tanggal_Transaksi   time.Time       `gorm:"datetime" json:"tangal_transaksi"`

	EventID      uuid.UUID  `gorm:"type:uuid" json:"event_id"`
	Event        Event      `gorm:"foreignKey:EventID" json:"event"`
	PembayaranID uuid.UUID  `gorm:"type:uuid" json:"pembayaran_id"`
	Pembayaran   Pembayaran `gorm:"foreignKey:PembayaranID" json:"pembayaran"`
	UserID       uuid.UUID  `gorm:"type:uuid" json:"user_id"`
	User         User       `gorm:"foreignKey:UserID" json:"user"`
}
