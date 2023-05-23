package pgdb

import (
	"context"
	"fmt"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/entity"
	"github.com/amiosamu/adv-backend-trainee-assignment/pkg/postgres"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const (
	maxPaginationLimit            = 10
	defaultPaginationLimit        = 10
	PriceSortType          string = "amount"
	DateSortType           string = "date"
)

type AdvertisementRepo struct {
	*postgres.DataSources
}

func NewAdvertisementRepo(pg *postgres.DataSources) *AdvertisementRepo {
	return &AdvertisementRepo{pg}
}

func (r *AdvertisementRepo) CreateAdvertisement(ctx context.Context, advertisement entity.Advertisement) (int, error) {
	var id int
	q := `INSERT INTO advertisement(name,description, pictures, price) values ($1, $2, $3, $4) RETURNING id`
	if err := r.DB.GetContext(ctx, advertisement, q, advertisement.Name, advertisement.Description, advertisement.Pictures, advertisement.Price); err != nil {
		return 0, err
	}
	err := r.DB.QueryRowxContext(ctx, q, advertisement.Name, advertisement.Description, advertisement.Pictures, advertisement.Price).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("AdvertisementRepo.CreateAdvertisement - r.DB.QueryRowContext: %v", err)
	}
	return id, nil
}

func (r *AdvertisementRepo) GetAdvertisementById(ctx context.Context, id int) (*entity.Advertisement, error) {
	advertisement := &entity.Advertisement{}
	q := `SELECT * FROM advertisement where id=$1`
	if err := r.DB.GetContext(ctx, advertisement, q, id); err != nil {
		log.Printf("unable to get user with email address: %v. Err: %v\n", id, advertisement)
		return nil, err
	}
	return advertisement, nil
}

func (r *AdvertisementRepo) GetAdvertisements(ctx context.Context, id int, sortType string, offset int, limit int) ([]entity.Advertisement, error) {
	var advertisements []entity.Advertisement
	query := `SELECT * FROM advertisement`
	err := r.DB.SelectContext(ctx, &advertisements, fmt.Sprintf(query, sortType), limit, offset)
	if err != nil {
		return nil, fmt.Errorf("AdvertisementRepo.GetAdvertisements - r.DB.SelectContext: %v", err)
	}
	return advertisements, nil
}
