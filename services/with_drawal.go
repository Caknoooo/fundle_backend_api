package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/google/uuid"
)

type PenarikanService interface {
	CreatePenarikan(ctx context.Context, penarikan entities.HistoryPenarikan) (entities.HistoryPenarikan, error)
	GetPenarikanByUser(ctx context.Context, userID uuid.UUID) ([]entities.HistoryPenarikan, error)
}

type penarikanService struct {
	penarikanRepository repository.PenarikanRepository
}

func NewPenarikanService(pr repository.PenarikanRepository) PenarikanService {
	return &penarikanService{
		penarikanRepository: pr,
	}
}

func(ps *penarikanService) CreatePenarikan(ctx context.Context, penarikan entities.HistoryPenarikan) (entities.HistoryPenarikan, error) {
	return ps.penarikanRepository.CreatePenarikan(ctx, penarikan)
}
 
func (ps *penarikanService) GetPenarikanByUser(ctx context.Context, userID uuid.UUID) ([]entities.HistoryPenarikan, error) {
	return ps.penarikanRepository.GetPenarikanByUser(ctx, userID)
}