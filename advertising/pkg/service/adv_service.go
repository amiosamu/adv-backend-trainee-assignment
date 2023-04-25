package service

import (
	"context"
	"github.com/amiosamu/adv-backend-trainee-assignment/advertising/model"
	"github.com/amiosamu/adv-backend-trainee-assignment/advertising/pkg/repository"
	"github.com/google/uuid"
)

type AdvService struct {
	repo repository.Advertising
}

func NewAdvService(repo repository.Advertising) *AdvService {
	return &AdvService{
		repo: repo,
	}
}

func (a *AdvService) Create(ctx context.Context, advertising model.Advertising) (uuid.UUID, error) {
	return a.repo.Create(ctx, advertising)
}

func (a *AdvService) GetAll(ctx context.Context, uuid uuid.UUID) ([]model.Advertising, error) {
	return a.repo.GetAll(ctx, uuid)
}

func (a *AdvService) GetByID(ctx context.Context, uuid uuid.UUID) (model.Advertising, error) {
	return a.repo.GetByID(ctx, uuid)
}
