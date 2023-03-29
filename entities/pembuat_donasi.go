package entities

import "github.com/google/uuid"

type PembuatDonasi struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	// Biodata pembuat donasi
	NamaDepanPembuat    string `gorm:"type:varchar(100)" json:"nama_depan_pembuat"`
	NamaBelakangPembuat string `gorm:"type:varchar(100)" json:"nama_belakang_pembuat"`
	NomorKTP            string `gorm:"type:varchar(20)" json:"nomor_ktp"`
	NomorTeleponPembuat string `gorm:"type:varchar(20)" json:"nomor_telepon_pembuat"`
	Pekerjaan           string `gorm:"type:varchar(100)" json:"pekerjaan"`
	AsalInstansi        string `gorm:"type:varchar(100)" json:"asal_pekerjaan"`
}
