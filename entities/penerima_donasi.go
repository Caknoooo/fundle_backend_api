package entities

import "github.com/google/uuid"

type PenerimaDonasi struct {
	ID                   uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	NamaDepanPenerima    string    `gorm:"type:varchar(100)" json:"nama_depan_penerima"`
	NamaBelakangPenerima string    `gorm:"type:varchar(100)" json:"nama_belakang_penerima"`
	TujuanGalangDana     string    `gorm:"type:varchar(100)" json:"tujuan_galang_dana"`
	LokasiTujuan         string    `gorm:"type:varchar(100)" json:"lokasi_tujuan"`
}
