package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type TransaksiService interface {
	CreateTransaksi(ctx context.Context, transaksiDTO dto.TransaksiCreateDTO) (entities.Transaksi, error)
	GetAllTransaksi(ctx context.Context) ([]entities.Transaksi, error)
	GetTransaksiByID(ctx context.Context, transaksiID uuid.UUID) (entities.Transaksi, error)
	GetAllTransaksiByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Transaksi, error)
	GetAllEventLastTransaksi(ctx context.Context, eventID uuid.UUID) ([]entities.Transaksi, error)
}

type transaksiService struct {
	transaksiRepository repository.TransaksiRepository
}

func NewTransaksiService(tr repository.TransaksiRepository) TransaksiService {
	return &transaksiService{
		transaksiRepository: tr,
	}
}

func (ts *transaksiService) CreateTransaksi(ctx context.Context, transaksiDTO dto.TransaksiCreateDTO) (entities.Transaksi, error) {
	transaksi := entities.Transaksi{}
	err := smapping.FillStruct(&transaksi, smapping.MapFields(transaksiDTO))
	if err != nil {
		return entities.Transaksi{}, err
	}
	return ts.transaksiRepository.CreateTransaksi(ctx, transaksi)
}

func (ts *transaksiService) GetAllTransaksi(ctx context.Context) ([]entities.Transaksi, error) {
	return ts.transaksiRepository.GetAllTransaksi(ctx)
}

func (ts *transaksiService) GetTransaksiByID(ctx context.Context, transaksiID uuid.UUID) (entities.Transaksi, error) {
	return ts.transaksiRepository.GetTransaksiByID(ctx, transaksiID)
}

func (ts *transaksiService) GetAllTransaksiByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Transaksi, error) {
	return ts.transaksiRepository.GetAllTransaksiByUserID(ctx, userID)
}

func (ts *transaksiService) GetAllEventLastTransaksi(ctx context.Context, eventID uuid.UUID) ([]entities.Transaksi, error) {
	return ts.transaksiRepository.GetAllEventLastTransaksi(ctx, eventID)
}
