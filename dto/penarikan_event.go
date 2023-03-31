package dto

import (
	"github.com/google/uuid"
)

type PenarikanEventDTO struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Jumlah_Penarikan  float64   `json:"jumlah_penarikan" json:"jumlah_penarikan" binding:"required"`
	
	BankID  uint      `json:"bank_id" form:"bank_id" binding:"required"`
	EventID uuid.UUID `json:"event_id" form:"event_id" binding:"required"`
}
