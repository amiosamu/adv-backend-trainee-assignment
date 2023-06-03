package service

import (
	"context"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/entity"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/repo"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/repo/repoerrors"
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
	advertisements, err := a.advertisementRepo.GetAdvertisements(ctx)
	if err != nil {
		return nil, err
	}
	return advertisements, nil
}

func (a *AdvertisementService) CreateAdvertisement(ctx context.Context, entity entity.Advertisement) (int, error) {
	id, err := a.advertisementRepo.CreateAdvertisement(ctx, entity)
	if err != nil {
		if err == repoerrors.ErrAlreadyExists {
			return 0, ErrAdvertisementAlreadyExists
		}
	}
	return id, nil

}

func (a *AdvertisementService) GetAdvertisementById(ctx context.Context, id int) (entity.Advertisement, error) {
	return a.advertisementRepo.GetAdvertisementById(ctx, id)
}
