package dto

import (
	"github.com/google/uuid"
)

type LikeDTO struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	EventID uuid.UUID `gorm:"type:uuid" form:"event_id" json:"event_id" binding:"required"`
	UserID  uuid.UUID `gorm:"type:uuid" form:"user_id" json:"user_id" binding:"required"`
}