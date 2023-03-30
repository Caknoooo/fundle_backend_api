package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/mashingan/smapping"
)

type PembayaranService interface {
	CreatePembayaran(ctx context.Context, pembayaranDTO dto.PembayaranDTO) (entities.Pembayaran, error)
}

type pembayaranService struct {
	pembayaranRepository repository.PembayaranRepository
}

func NewPembayaranService(pr repository.PembayaranRepository) PembayaranService {
	return &pembayaranService{
		pembayaranRepository: pr,
	}
}

func (ps *pembayaranService) CreatePembayaran(ctx context.Context, pembayaranDTO dto.PembayaranDTO) (entities.Pembayaran, error) {
	pembayaran := entities.Pembayaran{}
	err := smapping.FillStruct(&pembayaran, smapping.MapFields(pembayaranDTO))
	if err != nil {
		return entities.Pembayaran{}, err
	}
	return ps.pembayaranRepository.CreatePembayaran(ctx, pembayaran)
}
