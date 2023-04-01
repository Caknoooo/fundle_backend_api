package dto

import (
	"time"

	"github.com/google/uuid"
)

type EventCreateDTO struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	RekeningEvent  string    `json:"rekening_event" form:"rekening_event" binding:"required"`
	JudulEvent     string    `json:"judul_event" form:"judul_event" binding:"required"`
	DeskripsiEvent string    `json:"deskripsi_event" form:"deskripsi_event" binding:"required"`
	JenisEvent     string    `json:"jenis_event" form:"jenis_event" binding:"required"`
	MaxDonasi      float64   `json:"max_donasi" form:"max_donasi" binding:"required"`
	FotoEvent      string    `json:"foto_event" form:"foto_event" binding:"required"`
	ExpiredDonasi  time.Time `json:"expired_donasi" form:"expired_donasi" binding:"required"`

	NamaDepanPembuat    string `json:"nama_depan_pembuat" form:"nama_depan_pembuat" binding:"required"`
	NamaBelakangPembuat string `json:"nama_belakang_pembuat" form:"nama_belakang_pembuat" binding:"required"`
	NomorTeleponPembuat string `json:"nomor_telepon_pembuat" form:"nomor_telepon_pembuat" binding:"required"`
	NomorKTP            string `json:"nomor_ktp" form:"nomor_ktp" binding:"required"`
	Pekerjaan           string `json:"pekerjaan" form:"pekerjaan" binding:"required"`
	AsalInstansi        string `json:"asal_instansi" form:"asal_instansi" binding:"required"`

	NamaDepanPenerima    string `json:"nama_depan_penerima" form:"nama_depan_penerima" binding:"required"`
	NamaBelakangPenerima string `json:"nama_belakang_penerima" form:"nama_belakang_penerima" binding:"required"`
	TujuanGalangDana     string `json:"tujuan_galang_dana" form:"tujuan_galang_dana" binding:"required"`
	LokasiTujuan         string `json:"lokasi_tujuan" form:"lokasi_tujuan" binding:"required"`

	UserID uuid.UUID `json:"user_id" form:"user_id" binding:"required"`
}

type EventUpdateDTO struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	RekeningEvent  *string   `json:"rekening_event" form:"rekening_event"`
	Judul          *string   `json:"judul" form:"judul"`
	DeskripsiEvent *string   `json:"deskripsi_event" form:"deskripsi_event"`
	JenisEvent     *string   `json:"jenis_event" form:"jenis_event"`
	JumlahDonasi   *float64  `json:"jumlah_donasi" form:"jumlah_donasi"`
	FotoEvent      *string   `json:"foto_event" form:"foto_event"`
	UserID         *string   `json:"user_id" form:"user_id"`
	IsTargetFull   *bool     `json:"is_target_full" form:"is_target_full"`
	IsExpired      *bool     `json:"is_expired" form:"is_expired"`
}

type EventResponseServiceDTO struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Nama           string    `json:"nama" form:"nama"`
	DeskripsiEvent string    `json:"deskripsi_event" form:"deskripsi_event"`
	FotoEvent      string    `json:"foto_event" form:"foto_event"`
	ExpiredDonasi  time.Time `json:"expired_donasi" form:"expired_donasi" binding:"required"`
	IsExpired      bool      `json:"is_expired" form:"is_expired"`
}

type EventResponseDetailDonasiDTO struct {
	Nama           string    `json:"nama" form:"nama"`
	DeskripsiEvent string    `json:"deskripsi_event" form:"deskripsi_event"`
	FotoEvent      string    `json:"foto_event" form:"foto_event"`
	JumlahDonasi   float64   `json:"jumlah_donasi" form:"jumlah_donasi"`
	MaxDonasi      float64   `json:"max_donasi" form:"max_donasi"`
	ExpiredDonasi  time.Time `json:"expired_donasi" form:"expired_donasi" binding:"required"`
	IsExpired      bool      `json:"is_expired" form:"is_expired"`
	IsDone         uint64    `json:"is_done"`

	// Mengeluarkan 3 orang yang terakhir donasi
}

type EventResponseListDonasiDTO struct {
	Nama           string  `json:"nama" form:"nama"`
	DeskripsiEvent string  `json:"deskripsi_event" form:"deskripsi_event"`
	FotoEvent      string  `json:"foto_event" form:"foto_event"`
	JumlahDonasi   float64 `json:"jumlah_donasi" form:"jumlah_donasi"`
	MaxDonasi      float64 `json:"max_donasi" form:"max_donasi"`
}

type EventResponseMyEventDTO struct {
	Nama           string `json:"nama" form:"nama"`
	DeskripsiEvent string `json:"deskripsi_event" form:"deskripsi_event"`
	FotoEvent      string `json:"foto_event" form:"foto_event"`
}
