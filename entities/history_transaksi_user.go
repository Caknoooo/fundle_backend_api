package entities

import (
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/type/decimal"
)

type HistoryTransaksiUser struct {
	ID               uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	Status           string          `gorm:"type:varchar(50)" json:"status"`
	Jumlah_Transaksi decimal.Decimal `gorm:"type:decimal(15,2)" json:"jumlah_transaksi"`

	UserID      uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID" json:"user"`
	TransaksiID uuid.UUID `gorm:"type:uuid" json:"transaksi_id"`
	Transaksi   Transaksi `gorm:"foreignKey:TransaksiID" json:"transaksi"`
}
