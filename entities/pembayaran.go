package entities

import (
	"github.com/google/uuid"
)

type Pembayaran struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Jumlah          float64   `gorm:"type:float" json:"jumlah"`
	BuktiPembayaran string    `gorm:"type:varchar(50)" json:"bukti_pembayaran"`

	Transaksi          []Transaksi `gorm:"foreignKey:PembayaranID" json:"transaksi"`
	// StatusPembayaranID uint        `gorm:"type:uint" json:"status_pembayaran_id"`
	ListBankID         uint        `gorm:"type:uint" json:"list_bank_id"`
}
