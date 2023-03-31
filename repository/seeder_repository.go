package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type SeederRepository interface {
	GetAllBank(ctx context.Context) ([]entities.ListBank, error)
	GetBankByID(ctx context.Context, bankID uint) (entities.ListBank, error)
	GetAllCategory(ctx context.Context) ([]entities.CategoryEvent, error) 
	GetCategoryByID(ctx context.Context, categoryID uint) (entities.CategoryEvent, error)
	GetAllStatusPembayaran(ctx context.Context) ([]entities.StatusPembayaran, error)
	GetStatusPembayaranByID(ctx context.Context, statusID uint)(entities.StatusPembayaran, error)
}

type seederRepository struct {
	connection *gorm.DB
}

func NewSeederRepository(db *gorm.DB) SeederRepository {
	return &seederRepository{
		connection: db,
	}
}

func (sr *seederRepository) GetAllBank(ctx context.Context) ([]entities.ListBank, error) {
	var banks []entities.ListBank
	if err := sr.connection.Find(&banks).Error; err != nil {
		return nil, err
	}
	return banks, nil
}

func (sr *seederRepository) GetBankByID(ctx context.Context, bankID uint) (entities.ListBank, error) {
	var bank entities.ListBank
	if err := sr.connection.Where("id = ?", bankID).Take(&bank).Error; err != nil {
		return entities.ListBank{}, err
	}
	return bank, nil
}

func (sr *seederRepository) GetAllCategory(ctx context.Context) ([]entities.CategoryEvent, error) {
	var categories []entities.CategoryEvent
	if err := sr.connection.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (sr *seederRepository) GetCategoryByID(ctx context.Context, categoryID uint) (entities.CategoryEvent, error) {
	var category entities.CategoryEvent
	if err := sr.connection.Where("id = ?", categoryID).Take(&category).Error; err != nil {
		return entities.CategoryEvent{}, err
	}
	return category, nil
} 

func (sr *seederRepository) GetAllStatusPembayaran(ctx context.Context) ([]entities.StatusPembayaran, error) {
	var StatusPembayaran []entities.StatusPembayaran
	if err := sr.connection.Find(&StatusPembayaran).Error; err != nil {
		return nil, err
	}
	return StatusPembayaran, nil
}

func (sr *seederRepository) GetStatusPembayaranByID(ctx context.Context, statusID uint)(entities.StatusPembayaran, error) {
	var StatusPembayaran entities.StatusPembayaran
	if err := sr.connection.Where("id = ?", statusID).Take(&StatusPembayaran).Error; err != nil {
		return entities.StatusPembayaran{}, err
	}
	return StatusPembayaran, nil
}
