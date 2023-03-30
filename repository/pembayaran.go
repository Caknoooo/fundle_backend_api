package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type PembayaranRepository interface {
	CreatePembayaran(ctx context.Context, pembayaran entities.Pembayaran) (entities.Pembayaran, error)
}

type pembayaranRepository struct {
	connection *gorm.DB
}

func NewPembayaranRepository(db *gorm.DB) PembayaranRepository {
	return &pembayaranRepository{
		connection: db,
	}
}

func (pr *pembayaranRepository) CreatePembayaran(ctx context.Context, pembayaran entities.Pembayaran) (entities.Pembayaran, error) {
	if err := pr.connection.Create(&pembayaran).Error; err != nil {
		return entities.Pembayaran{}, err
	}
	return pembayaran, nil
}
