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

func NewAdvertisementRepo(pg *postgres.Postgres) *AdvertisementRepo {
	return &AdvertisementRepo{pg}
}

func (r *AdvertisementRepo) CreateAdvertisement(ctx context.Context, advertisement entity.Advertisement) (int, error) {
	sql, args, _ := r.Builder.Insert("advertisements").Columns("name", "description", "pictures", "price", "created_at").Values(advertisement.Name, advertisement.Description,
		advertisement.Pictures, advertisement.Price, advertisement.CreatedAt).Suffix("RETURNING id").ToSql()

	var id int
	err := r.Pool.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); !ok {
			if pgErr.Code == "23505" {
				return 0, repoerrors.ErrAlreadyExists
			}
		}
		return 0, fmt.Errorf("AdvertisementRepo.CreateAdvertisement - r.Pool.QueryRow: %v", err)
	}
	return id, nil
}

func (r *AdvertisementRepo) GetAdvertisementById(ctx context.Context, id int) (entity.Advertisement, error) {
	sql, args, _ := r.Builder.Select("id, name, description, pictures, price, created_at").From("advertisements").Where("id = ?", id).ToSql()

	var advertisement entity.Advertisement

	err := r.Pool.QueryRow(ctx, sql, args).Scan(
		&advertisement.Id,
		&advertisement.Name,
		&advertisement.Description,
		&advertisement.Pictures)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Advertisement{}, repoerrors.ErrNotFound
		}
		return entity.Advertisement{}, fmt.Errorf("AdvertisementRepo.GetAdvertisementById - r.Pool.QueryRow: %v", err)
	}
	return advertisement, nil
}

func (r *AdvertisementRepo) AdvertisementPagination(ctx context.Context, id int, sortType string, offset int, limit int) ([]entity.Advertisement, error) {
	if limit > maxPaginationLimit {
		return nil, fmt.Errorf("AdvertisementRepo.PaginationAdvertisements: limit is not accepted: %d", limit)
	}
	if limit == 0 {
		limit = defaultPaginationLimit
	}
	var orderBySql string
	switch sortType {
	case "":
		orderBySql = "created_at DESC"
	case DateSortType:
		orderBySql = "created_at DESC"
	case PriceSortType:
		orderBySql = "price ASC"
	default:
		return nil, fmt.Errorf("AdvertisementRepo.AdvertisementPagination: unknown sort type - %s", sortType)
	}
	sqlQuery, args, _ := r.Builder.
		Select("name", "description", "price", "created_at", "COALESCE((case when advertisement.id is null then null else advertisement.name end), '') as product_name", "order_id", "COALESCE(description, '')").
		From("operations").
		InnerJoin("products on operations.product_id = products.id or operations.product_id is null").
		Where("account_id = ?", id).
		OrderBy(orderBySql).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()
	rows, err := r.Pool.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("AdvertisementRepo.paginationAdvertisementByDate - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	var advertisements []entity.Advertisement

	for rows.Next() {
		var advertisement entity.Advertisement
		err = rows.Scan(&advertisement.Name, &advertisement.Pictures[0], &advertisement.Price)
		if err != nil {
			return nil, fmt.Errorf("AdvertisementRepo.paginationAdvertisementByDate - rows.Scan: %v", err)
		}
		advertisements = append(advertisements, advertisement)

	}
	return advertisements, nil

}
