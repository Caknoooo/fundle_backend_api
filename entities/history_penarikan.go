package entities

import (
	"time"

	"github.com/google/uuid"
)

type HistoryPenarikan struct {
	ID                uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Jumlah_Penarikan  float64   `gorm:"type:float" json:"jumlah_penarikan"`
	NamaBank          string    `gorm:"type:varchar(100)" json:"nama_bank"`
	Tanggal_Penarikan time.Time `gorm:"timestamp with time zone" json:"tangal_penarikan"`

	BankID  uint      `gorm:"type:uint" json:"bank_id"`
	Bank    ListBank  `gorm:"foreignKey:BankID" json:"-"`
	UserID  uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User    User      `gorm:"foreignKey:UserID" json:"-"`
	EventID uuid.UUID `gorm:"type:uuid" json:"event_id"`
	Event   Event     `gorm:"foreignKey:EventID" json:"event"`
}
	