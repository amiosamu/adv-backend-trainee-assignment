package service

import (
	"context"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/entity"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/repo"
)

type AdvertisementService struct {
	advertisementRepo repo.Advertisement
}

func NewAdvertisementService(advertisementRepo repo.Advertisement) *AdvertisementService {
	return &AdvertisementService{
		advertisementRepo: advertisementRepo,
	}
}

func (a *AdvertisementService) GetAdvertisements(ctx context.Context) ([]entity.Advertisement, error) {
	panic("implement me")
}

func (a *AdvertisementService) CreateAdvertisement(ctx context.Context, entity entity.Advertisement) (int, error) {
	panic("implement me")

}

func (a *AdvertisementService) GetAdvertisementById(ctx context.Context, id int) (entity.Advertisement, error) {
	panic("implement me")
}
