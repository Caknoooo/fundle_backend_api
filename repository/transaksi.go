package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransaksiRepository interface {
	CreateTransaksi(ctx context.Context, transaksi entities.Transaksi) (entities.Transaksi, error)
	GetAllTransaksi(ctx context.Context) ([]entities.Transaksi, error)
	GetTransaksiByID(ctx context.Context, transaksiID uuid.UUID) (entities.Transaksi, error)
	GetAllTransaksiByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Transaksi, error)
	GetAllEventLastTransaksi(ctx context.Context, eventID uuid.UUID) ([]entities.Transaksi, error)
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

	if err:= tr.connection.Preload("Pembayaran").Preload("Event").Preload("User").First(&transaksi).Error;err != nil {
		return entities.Transaksi{}, err
	}
	return transaksi, nil
}

func (tr *transaksiRepository) GetAllTransaksi(ctx context.Context) ([]entities.Transaksi, error) {
	var transaksi []entities.Transaksi
	if err := tr.connection.Preload("Pembayaran").Preload("Event").Preload("User").Find(&transaksi).Error; err != nil {
		return nil, err
	}
	return transaksi, nil
}

func (tr *transaksiRepository) GetTransaksiByID(ctx context.Context, transaksiID uuid.UUID) (entities.Transaksi, error) {
	var transaksi entities.Transaksi
	if err := tr.connection.Preload("Pembayaran").Preload("Event").Preload("User").Where("id = ?", transaksiID).Find(&transaksi).Error; err != nil {
		return entities.Transaksi{}, err
	}
	return transaksi, nil
}

func (tr *transaksiRepository) GetAllTransaksiByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Transaksi, error) {
	var transaksi []entities.Transaksi
	if err := tr.connection.Preload("User").Where("user_id = ?", userID).Find(&transaksi).Error; err != nil {
		return nil, err
	}
	return transaksi, nil
}

// func (tr *transaksiRepository) GetAllEventLastTransaksi(ctx context.Context, eventID uuid.UUID) ([]entities.Transaksi, error) {
// 	var transaksi []entities.Transaksi
// 	if err := tr.connection.Where("event_id = ?", eventID).Order("Tanggal_Transaksi desc").Limit(3).Find(&transaksi).Error; err != nil {
// 		return nil, err
// 	}
// 	return transaksi, nil
// }

func (tr *transaksiRepository) GetAllEventLastTransaksi(ctx context.Context, eventID uuid.UUID) ([]entities.Transaksi, error) {
	var transaksi []entities.Transaksi
	if err := tr.connection.
		Joins("JOIN users ON transaksis.user_id = users.id").
		Where("transaksis.event_id = ?", eventID).
		Order("transaksis.Tanggal_Transaksi desc").
		Limit(3).
		Select("transaksis.*, users.nama as Nama_User").Preload("Pembayaran").Preload("Event").Preload("User").
		Find(&transaksi).
		Error; err != nil {
		return nil, err
	}
	return transaksi, nil
}