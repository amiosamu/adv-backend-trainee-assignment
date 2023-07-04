package repo

import (
	"context"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/entity"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/repo/pgdb"
	"github.com/jmoiron/sqlx"
)

type Advertisement interface {
	CreateAdvertisement(ctx context.Context, advertisement *entity.Advertisement) (int, error)
	GetAdvertisementById(ctx context.Context, id int) (entity.Advertisement, error)
	GetAdvertisements(ctx context.Context) ([]entity.Advertisement, error)
}

type Repos struct {
	Advertisement
}

func NewRepos(pg *sqlx.DB) *Repos {
	return &Repos{
		Advertisement: pgdb.NewAdvertisementRepo(pg),
	}
}
