package entities

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`

	// Postingan
	RekeningEvent  string    `gorm:"type:varchar(100)" json:"rekening_event"`
	JudulEvent     string    `gorm:"type:varchar(100)" json:"judul_event"`
	DeskripsiEvent string    `gorm:"type:text" json:"deskripsi_event"`
	JenisEvent     string    `gorm:"type:varchar(100)" json:"jenis_event"`
	FotoEvent      string    `gorm:"type:varchar(100)" json:"foto_event"`
	MaxDonasi      float64   `gorm:"type:float" json:"max_donasi"`
	JumlahDonasi   float64   `gorm:"type:float" json:"jumlah_donasi"`
	LikeCount      uint64    `json:"like_count"`
	ExpiredDonasi  time.Time `gorm:"datetime" json:"expired_donasi"`
	IsDone         uint64    `json:"is_done"`
	Is_target_full bool      `gorm:"type:boolean" json:"is_target_full"`
	Is_expired     bool      `gorm:"type:boolean" json:"is_expired"`

	// Pembuat Event
	NamaDepanPembuat    string `gorm:"type:varchar(100)" json:"nama_depan_pembuat"`
	NamaBelakangPembuat string `gorm:"type:varchar(100)" json:"nama_belakang_pembuat"`
	NomorKTP            string `gorm:"type:varchar(20)" json:"nomor_ktp"`
	NomorTeleponPembuat string `gorm:"type:varchar(20)" json:"nomor_telepon_pembuat"`
	Pekerjaan           string `gorm:"type:varchar(100)" json:"pekerjaan"`
	AsalInstansi        string `gorm:"type:varchar(100)" json:"asal_pekerjaan"`

	// Penerima Event
	NamaDepanPenerima    string    `gorm:"type:varchar(100)" json:"nama_depan_penerima"`
	NamaBelakangPenerima string    `gorm:"type:varchar(100)" json:"nama_belakang_penerima"`
	TujuanGalangDana     string    `gorm:"type:varchar(100)" json:"tujuan_galang_dana"`
	LokasiTujuan         string    `gorm:"type:varchar(100)" json:"lokasi_tujuan"`

	UserID           uuid.UUID          `gorm:"type:uuid" json:"user_id"`
	User             User               `gorm:"foreignKey:UserID" json:"-"`

	Likes            []Like             `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	HistoryPenarikan []HistoryPenarikan `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"history_penarikan"`
	Transaksi        []Transaksi        `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"transaksi,omitempty"`
}
	// PembuatDonasiID  uuid.UUID          `gorm:"type:uuid" json:"pembuat_donasi_id"`
	// PembuatDonasi    PembuatDonasi      `gorm:"foreignKey:PembuatDonasiID" json:"pembuat_donasi"`
	// PenerimaDonasiID uuid.UUID          `gorm:"type:uuid" json:"penerima_donasi_id"`
	// PenerimaDonasi   PenerimaDonasi     `gorm:"foreignKey:PenerimaDonasiID" json:"penerima_donasi"`