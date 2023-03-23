package entities

import (
	"time"

	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/type/decimal"
)

type Event struct {
	ID             uuid.UUID       `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	RekeningEvent  string          `gorm:"type:varchar(100)" json:"rekening_event"`
	Nama           string          `gorm:"type:varchar(100)" json:"nama"`
	DeskripsiEvent string          `gorm:"type:text" json:"deskripsi_event"`
	JenisEvent     string          `gorm:"type:varchar(100)" json:"jenis_event"`
	JumlahDonasi   decimal.Decimal `gorm:"type:decimal(15,2)" json:"jumlah_donasi"`
	FotoEvent      string          `gorm:"type:varchar(100)" json:"foto_event"` // Ini kalau mau dibuat banyak foto, harus one to many
	ExpiredDonasi  time.Time       `gorm:"datetime" json:"expired_donasi"`
	Is_target_full bool            `gorm:"type:boolean" json:"is_target_full"`
	Is_expired     bool            `gorm:"type:boolean" json:"is_expired"`
	
	UserID         uuid.UUID       `gorm:"type:uuid" json:"user_id"`
	User           User            `gorm:"foreignKey:UserID" json:"user"`

	// WithDrawals   []int	 `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"with_drawal,omitempty"` // Temp
	// Transaksi			[]int `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"transaksi,omitempty"`// Temp
}
