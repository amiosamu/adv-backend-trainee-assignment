package service

import (
	"context"
	"github.com/amiosamu/adv-backend-trainee-assignment/advertising/model"
	"github.com/amiosamu/adv-backend-trainee-assignment/advertising/pkg/repository"
	"github.com/google/uuid"
)

type Advertising interface {
	Create(ctx context.Context, advertising model.Advertising) (uuid.UUID, error)
	GetAll(ctx context.Context, uuid uuid.UUID) ([]model.Advertising, error)
	GetByID(ctx context.Context, uuid uuid.UUID) (model.Advertising, error)
}

type Service struct {
	Advertising
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Advertising: NewAdvService(repo.Advertising),
	}
}
