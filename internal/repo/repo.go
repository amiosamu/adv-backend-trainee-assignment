package repo

import (
	"context"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/entity"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/repo/pgdb"
	postgres "github.com/amiosamu/adv-backend-trainee-assignment/pkg/postgres"
)

type Advertisement interface {
	CreateAdvertisement(ctx context.Context, advertisement entity.Advertisement) (int, error)
	GetAdvertisementById(ctx context.Context, id int) (*entity.Advertisement, error)
	GetAdvertisements(ctx context.Context, id int, sortType string, offset int, limit int) ([]entity.Advertisement, error)
}

type Repos struct {
	Advertisement
}

func NewRepos(pg *postgres.Postgres) *Repos {
	return &Repos{
		Advertisement: pgdb.NewAdvertisementRepo(pg),
	}
}
