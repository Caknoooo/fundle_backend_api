package entities

import (
	"github.com/google/uuid"
)

type HistoryTransaksiUser struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Status           string    `gorm:"type:varchar(50)" json:"status"`
	Jumlah_Transaksi float64   `gorm:"type:float" json:"jumlah_transaksi"`
	
	TransaksiID      uuid.UUID `gorm:"type:uuid" json:"transaksi_id"`
}
