package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type TransaksiRepository interface {
	CreateTransaksi(ctx context.Context, transaksi entities.Transaksi) (entities.Transaksi, error)
	CreateTransaksiUser(ctx context.Context, transaksi entities.User_Transaksi) (entities.User_Transaksi, error)
	GetAllTransaksi(ctx context.Context) ([]entities.Transaksi, error)
}

type transaksiRepository struct {
	connection *gorm.DB
}

func NewTransaksiRepository(db *gorm.DB) TransaksiRepository {
	return &transaksiRepository{
		connection: db,
	}
}

func (tr *transaksiRepository) CreateTransaksi(ctx context.Context, transaksi entities.Transaksi) (entities.Transaksi, error) {
	if err := tr.connection.Create(&transaksi).Error; err != nil {
		return entities.Transaksi{}, err
	}
	return transaksi, nil
}

func (tr *transaksiRepository) CreateTransaksiUser(ctx context.Context, transaksi entities.User_Transaksi) (entities.User_Transaksi, error) {
	if err := tr.connection.Create(&transaksi).Error; err != nil {
		return entities.User_Transaksi{}, err
	}
	return transaksi, nil
}

func (tr *transaksiRepository) GetAllTransaksi(ctx context.Context) ([]entities.Transaksi, error) {
	var user []entities.Transaksi
	if err := tr.connection.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
