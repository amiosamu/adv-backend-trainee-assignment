package pgdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/entity"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/repo/repoerrors"
	"github.com/amiosamu/adv-backend-trainee-assignment/pkg/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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
	sql, args, _ := a.Builder.Insert("advertisement").Columns("name", "description", "pictures", "price", "created_at").Values(advertisement.Name, advertisement.Description, advertisement.Pictures, advertisement.Price, advertisement.CreatedAt).Suffix("RETURNING id").ToSql()
	var id int
	err := a.Pool.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return 0, repoerrors.ErrAlreadyExists
			}
		}
		return 0, fmt.Errorf("AdvertisementRepo.CreateUser - r.Pool.QueryRow: %v", err)
	}

	return id, nil
}

func (a *AdvertisementRepo) GetAdvertisementById(ctx context.Context, id int) (entity.Advertisement, error) {
	sql, args, _ := a.Builder.Select("name, description, pictures, price, created_at").From("advertisement").Where("id = ?", id).ToSql()
	var advertisement entity.Advertisement
	err := a.Pool.QueryRow(ctx, sql, args...).Scan(&advertisement.Name, &advertisement.Description, &advertisement.Pictures, &advertisement.Price, &advertisement.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Advertisement{}, repoerrors.ErrNotFound
		}
		return entity.Advertisement{}, fmt.Errorf("AdvRepo.GetAdvertisementById - r.Pool.QueryRow: %v", err)
	}
	return advertisement, nil
}

func (a *AdvertisementRepo) GetAdvertisements(ctx context.Context) ([]entity.Advertisement, error) {
	sql, args, _ := a.Builder.Select("id, name, description, pictures, price, created_at").From("advertisement").ToSql()

	rows, err := a.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("UserRepo.GetAllUsers - r.Pool.Query: %v", err)
	}
	defer rows.Close()

	var advertisements []entity.Advertisement

	for rows.Next() {
		var advertisement entity.Advertisement
		err := rows.Scan(
			&advertisement.Id,
			&advertisement.Name,
			&advertisement.Description,
			&advertisement.Pictures,
			&advertisement.Price,
			&advertisement.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("AdvRepo.GetAdvertisements - rows.Scan: %v", err)
		}
		advertisements = append(advertisements, advertisement)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("AdvRepo.GetAdvertisements - rows.Err: %v", err)
	}
	return advertisements, nil
}

func NewAdvertisementRepo(pg *postgres.Postgres) *AdvertisementRepo {
	return &AdvertisementRepo{pg}
}
