package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
)

type SeederService interface {
	GetAllBank(ctx context.Context)([]entities.ListBank, error)
	GetBankByID(ctx context.Context, bankID uint)(entities.ListBank, error) 
	GetAllCategory(ctx context.Context) ([]entities.CategoryEvent,error)
	GetCategoryByID(ctx context.Context, categoryID uint) (entities.CategoryEvent, error)
	GetAllStatusPembayaran(ctx context.Context)([]entities.StatusPembayaran, error)
	GetStatusPembayaranByID(ctx context.Context, statusID uint) (entities.StatusPembayaran, error)
}

type seederService struct {
	seederRepository repository.SeederRepository
}

func NewSeederService(sr repository.SeederRepository) SeederService {
	return &seederService{
		seederRepository: sr,
	}
}

func (ss *seederService) GetAllBank(ctx context.Context)([]entities.ListBank, error){
	return ss.seederRepository.GetAllBank(ctx)
} 

func (ss *seederService) GetBankByID(ctx context.Context, bankID uint)(entities.ListBank, error) {
	return ss.seederRepository.GetBankByID(ctx, bankID)
}

func (ss *seederService) GetAllCategory(ctx context.Context) ([]entities.CategoryEvent,error) {
	return ss.seederRepository.GetAllCategory(ctx)
}

func (ss *seederService) GetCategoryByID(ctx context.Context, categoryID uint) (entities.CategoryEvent, error){
	return ss.seederRepository.GetCategoryByID(ctx, categoryID)
}

func(ss *seederService) GetAllStatusPembayaran(ctx context.Context)([]entities.StatusPembayaran, error) {
	return ss.seederRepository.GetAllStatusPembayaran(ctx)
}

func (ss *seederService) GetStatusPembayaranByID(ctx context.Context, statusID uint) (entities.StatusPembayaran, error) {
	return ss.seederRepository.GetStatusPembayaranByID(ctx, statusID)
}