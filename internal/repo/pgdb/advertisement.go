package pgdb

import (
	"context"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/entity"
	"github.com/amiosamu/adv-backend-trainee-assignment/pkg/postgres"
	_ "github.com/lib/pq"
)

const (
	maxPaginationLimit            = 10
	defaultPaginationLimit        = 10
	PriceSortType          string = "amount"
	DateSortType           string = "date"
)

type AdvertisementRepo struct {
	*postgres.Postgres
}

func (a *AdvertisementRepo) CreateAdvertisement(ctx context.Context, advertisement entity.Advertisement) (int, error) {
	panic("implement me")
}

func (a *AdvertisementRepo) GetAdvertisementById(ctx context.Context, id int) (entity.Advertisement, error) {
	panic("implement me")
}

func (a *AdvertisementRepo) GetAdvertisements(ctx context.Context) ([]entity.Advertisement, error) {
	panic("implement me")
}

func NewAdvertisementRepo(pg *postgres.Postgres) *AdvertisementRepo {
	return &AdvertisementRepo{pg}
}
