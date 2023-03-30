package dto

type PembayaranDTO struct {
	Jumlah             float64 `gorm:"type:float" json:"jumlah" binding:"required"`
	BuktiPembayaran    string  `json:"bukti_pembayaran" binding:"required"`
	StatusPembayaranID uint    `json:"status_pembayaran_id" binding:"required"`
	ListBankID         uint    `json:"list_bank_id" binding:"required"`
}
