package entities

import (
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/type/decimal"
)

type Pembayaran struct {
	ID              uuid.UUID       `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Jumlah          decimal.Decimal `gorm:"type:decimal(15,2)" json:"jumlah"`
	BuktiPembayaran string          `gorm:"type:varchar(50)" json:"bukti_pembayaran"`

	StatusPembayaranID uint `gorm:"type:uint" json:"status_pembayaran_id"`
	ListBankID         uint `gorm:"type:uint" json:"list_bank_id"`
}
