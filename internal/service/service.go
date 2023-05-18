package service

import (
	"context"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/entity"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/repo"
)

type Dependencies struct {
	Repos *repo.Repos
}

type AdvertisementInput struct {
	Id       int
	SortType string
	Offset   int
	Limit    int
}
type AdvertisementOutput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Pictures    string `json:"pictures"`
	Price       int    `json:"price"`
	Order       *int   `json:"order"`
}

type Advertisement interface {
	CreateAdvertisement(ctx context.Context, entity entity.Advertisement) (int, error)
	GetAdvertisementById(ctx context.Context, id int) (entity.Advertisement, error)
	GetAdvertisements(ctx context.Context) ([]entity.Advertisement, error)
}

type Services struct {
	Advertisement Advertisement
}

func NewServices(deps Dependencies) *Services {
	return &Services{
		Advertisement: NewAdvertisementService(deps.Repos.Advertisement),
	}
}
