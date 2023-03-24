package dto

import "google.golang.org/genproto/googleapis/type/decimal"

type PembayaranDTO struct {
	Jumlah             decimal.Decimal `json:"jumlah" binding:"required"`
	BuktiPembayaran    string          `json:"bukti_pembayaran" binding:"required"`
	StatusPembayaranID uint            `json:"status_pembayaran_id" binding:"required"`
	ListBankID         uint            `json:"list_bank_id" binding:"required"`
}