package repository

import (
	"context"
	"github.com/amiosamu/adv-backend-trainee-assignment/advertising/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Advertising interface {
	Create(ctx context.Context, adversary model.Advertising) (uuid.UUID, error)
	GetAll(ctx context.Context, uuid uuid.UUID) ([]model.Advertising, error)
	GetByID(ctx context.Context, uuid uuid.UUID) (model.Advertising, error)
}

type Repository struct {
	Advertising
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Advertising: NewAdvPostgres(db),
	}
}
