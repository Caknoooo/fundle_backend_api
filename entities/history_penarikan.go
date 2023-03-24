package entities

import (
	"time"

	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/type/decimal"
)

type HistoryPenarikan struct {
	ID                uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	Jumlah_Penarikan  decimal.Decimal `gorm:"type:decimal(15,2)" json:"jumlah_penarikan"`
	Tanggal_Transaksi time.Time       `gorm:"datetime" json:"tangal_transaksi"`

	UserID  uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User    User      `gorm:"foreignKey:UserID" json:"user"`
	EventID uuid.UUID `gorm:"type:uuid" json:"event_id"`
	Event   Event     `gorm:"foreignKey:EventID" json:"event"`
}
