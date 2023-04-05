package repository

import (
	"context"
	"errors"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PenarikanRepository interface {
	CreatePenarikan(ctx context.Context, penarikan entities.HistoryPenarikan) (entities.HistoryPenarikan, error)
	GetPenarikanByUser(ctx context.Context, userID uuid.UUID) ([]entities.HistoryPenarikan, error)
}

type penarikanRepository struct {
	connection *gorm.DB
}

func NewPenarikanRepository(db *gorm.DB) PenarikanRepository {
	return &penarikanRepository{
		connection: db,
	}
}

func (pr *penarikanRepository) CreatePenarikan(ctx context.Context, penarikan entities.HistoryPenarikan) (entities.HistoryPenarikan, error){
	var updateEvent entities.Event
	if err := pr.connection.First(&updateEvent, penarikan.EventID).Error; err != nil {
		return entities.HistoryPenarikan{}, err
	}

	if updateEvent.SisaDonasi - penarikan.Jumlah_Penarikan < 0 {
		return entities.HistoryPenarikan{}, errors.New("Penarikan Melebihi Batas Saldo Yang Tersisa")
	} else {
		updateEvent.SisaDonasi -= penarikan.Jumlah_Penarikan
		pr.connection.Save(&updateEvent)
	}

	if err := pr.connection.Create(&penarikan).Error; err != nil {
		return entities.HistoryPenarikan{}, err
	}
	return penarikan, nil
}

func (pr *penarikanRepository) GetPenarikanByUser(ctx context.Context, userID uuid.UUID) ([]entities.HistoryPenarikan, error) {
	var penarikan []entities.HistoryPenarikan
	if err := pr.connection.Preload("Event").Where("user_id = ?", userID).Find(&penarikan).Error; err != nil {
		return nil, err
	}
	return penarikan, nil
}