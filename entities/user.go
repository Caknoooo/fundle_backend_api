package entities

import (
	"github.com/Caknoooo/golang-clean_template/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Nama            string    `gorm:"type:varchar(100)" json:"nama"`
	NoTelp          string    `gorm:"type:varchar(30)" json:"no_telp"`
	Email           string    `gorm:"type:varchar(100)" json:"email"`
	Password        string    `gorm:"type:varchar(100)" json:"password"`
	ConfirmPassword string    `gorm:"type:varchar(100)" json:"confirm_password"`
	Role            string    `gorm:"type:varchar(100)" json:"role"`

	Transaksi []Transaksi `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"transaksi,omitempty"`
	Events    []Event     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"events,omitempty"`
	Likes     []Like      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"likes,omitempty"`

	Timestamp
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) CheckNil(tx *gorm.DB) error {
	if u.Events == nil {
		u.Events = []Event{}
	}
	return nil
}
