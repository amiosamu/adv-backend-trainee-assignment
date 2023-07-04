package pgdb

import (
	"context"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"time"
)

const (
	maxPaginationLimit            = 10
	defaultPaginationLimit        = 10
	PriceSortType          string = "amount"
	DateSortType           string = "date"
)

type AdvertisementRepo struct {
	DB *sqlx.DB
}

func (a *AdvertisementRepo) CreateAdvertisement(ctx context.Context, advertisement *entity.Advertisement) (int, error) {
	q := `INSERT INTO advertisement (name, description, pictures, price) VALUES ($1, $2, $3, $4) RETURNING id`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	args := []interface{}{advertisement.Name, advertisement.Description, pq.Array(&advertisement.Pictures), advertisement.Price}
	err := a.DB.QueryRowContext(ctx, q, args...).Scan(&advertisement.Id)
	if err != nil {
		return 0, err
	}
	return advertisement.Id, nil
}

func (a *AdvertisementRepo) GetAdvertisementById(ctx context.Context, id int) (entity.Advertisement, error) {
	panic("implement me")
}

func (a *AdvertisementRepo) GetAdvertisements(ctx context.Context) ([]entity.Advertisement, error) {
	panic("implement me")
}

func NewAdvertisementRepo(pg *sqlx.DB) *AdvertisementRepo {
	return &AdvertisementRepo{pg}
}
