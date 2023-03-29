package dto

import "github.com/google/uuid"

type EventPembuatDonasiDTO struct {
	ID                  uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	NamaDepanPembuat    string    `json:"nama_depan_pembuat" form:"nama_depan" binding:"required"`
	NamaBelakangPembuat string    `json:"nama_belakang_pembuat" form:"nama_belakang" binding:"required"`
	NomorTeleponPembuat string    `json:"nomor_telepon_pembuat" form:"nomor_telepon_pembuat" binding:"required"`
	NomorKTP            string    `json:"nomor_ktp" form:"nomor_ktp" binding:"required"`
	Pekerjaan           string    `json:"pekerjaan" form:"pekerjaan" binding:"reqired"`
	AsalInstansi        string    `json:"asal_instansi" form:"asal_instansi" binding:"required"`
}

type EventPenerimaDonasiDTO struct {
	ID               uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	NamaDepan        string    `json:"nama_depan_peneima" form:"nama_depan_penerima" binding:"required"`
	NamaBelakang     string    `json:"nama_belakang_penerima" form:"nama_belakang_penerima" binding:"required"`
	TujuanGalangDana string    `json:"tujuan_galang_dana" form:"tujuan_galang_dana" binding:"required"`
	LokasiTujuan     string    `json:"lokasi_tujuan" form:"lokasi_tujuan" binding:"required"`
}
