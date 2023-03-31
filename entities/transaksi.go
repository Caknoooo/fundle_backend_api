package entities

import (
	"time"

	"github.com/google/uuid"
)

type Transaksi struct {
	ID                  uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	NamaBank            string    `gorm:"type:varchar(100)" json:"nama_bank"`
	Jumlah_Donasi_Event float64   `gorm:"type:float" json:"jumlah_donasi"`
	Tanggal_Transaksi   time.Time `gorm:"timestamp with time zone" json:"tangal_transaksi"`

	// HistoryTransaksiUser HistoryTransaksiUser `gorm:"foreignKey:TransaksiID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"history_transaksi_users,omitempty"`
	UserID       uuid.UUID  `gorm:"type:uuid" json:"user_id"`
	User         User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
	EventID      uuid.UUID  `gorm:"type:uuid" json:"event_id"`
	Event        Event      `gorm:"foreignKey:EventID" json:"-"`
	PembayaranID uuid.UUID  `gorm:"type:uuid" json:"pembayaran_id"`
	Pembayaran   Pembayaran `gorm:"foreignKey:PembayaranID" json:"-"`
}
