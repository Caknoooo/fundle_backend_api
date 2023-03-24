package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/mashingan/smapping"
)

type TransaksiService interface {
	CreateTransaksi(ctx context.Context, transaksiDTO dto.TransaksiCreateDTO) (entities.Transaksi, error)
	CreateTransaksiUser(ctx context.Context, transaksiDTO dto.TransaksiUserCreateDTO) (entities.User_Transaksi, error)
	GetAllTransaksi(ctx context.Context) ([]entities.Transaksi, error)
}

type transaksiService struct {
	transaksiRepository repository.TransaksiRepository
}

func NewTransaksiService(ur repository.TransaksiRepository) TransaksiService {
	return &transaksiService{
		transaksiRepository: ur,
	}
}

func (ts *transaksiService) CreateTransaksi(ctx context.Context, transaksiDTO dto.TransaksiCreateDTO) (entities.Transaksi, error) {
	user := entities.Transaksi{}
	err := smapping.FillStruct(&user, smapping.MapFields(transaksiDTO))
	if err != nil {
		return entities.Transaksi{}, err
	}
	return ts.transaksiRepository.CreateTransaksi(ctx, user)
}

func (ts *transaksiService) CreateTransaksiUser(ctx context.Context, transaksiDTO dto.TransaksiUserCreateDTO) (entities.User_Transaksi, error) {
	user := entities.User_Transaksi{}
	err := smapping.FillStruct(&user, smapping.MapFields(transaksiDTO))
	if err != nil {
		return entities.User_Transaksi{}, err
	}
	return ts.transaksiRepository.CreateTransaksiUser(ctx, user)
}

func (ts *transaksiService) GetAllTransaksi(ctx context.Context) ([]entities.Transaksi, error) {
	return ts.transaksiRepository.GetAllTransaksi(ctx)
}
