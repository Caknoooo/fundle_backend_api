package entities

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	RekeningEvent  string    `gorm:"type:varchar(100)" json:"rekening_event"`
	Nama           string    `gorm:"type:varchar(100)" json:"nama"`
	DeskripsiEvent string    `gorm:"type:text" json:"deskripsi_event"`
	JenisEvent     string    `gorm:"type:varchar(100)" json:"jenis_event"`
	JumlahDonasi   float64   `gorm:"type:float" json:"jumlah_donasi"`
	MaxDonasi      float64   `gorm:"type:float" json:"max_donasi"`
	LikeCount      uint64    `json:"like_count"`
	FotoEvent      string    `gorm:"type:varchar(100)" json:"foto_event"`
	ExpiredDonasi  time.Time `gorm:"datetime" json:"expired_donasi"`
	IsDone         uint64    `json:"is_done"`
	Is_target_full bool      `gorm:"type:boolean" json:"is_target_full"`
	Is_expired     bool      `gorm:"type:boolean" json:"is_expired"`

	UserID           uuid.UUID          `gorm:"type:uuid" json:"user_id"`
	User             User               `gorm:"foreignKey:UserID" json:"user"`
	Likes            []Like             `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"likes"`
	HistoryPenarikan []HistoryPenarikan `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"history_penarikan"`
	Transaksi        []Transaksi        `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"transaksi,omitempty"`
}
