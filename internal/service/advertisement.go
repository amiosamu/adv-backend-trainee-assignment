package service

import (
	"context"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/entity"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/repo"
	log "github.com/sirupsen/logrus"
)

type AdvertisementService struct {
	advertisementRepo repo.Advertisement
}

func NewAdvertisementService(advertisementRepo repo.Advertisement) *AdvertisementService {
	return &AdvertisementService{
		advertisementRepo: advertisementRepo,
	}
}

func (a *AdvertisementService) CreateAdvertisement(ctx context.Context, advertisement *entity.Advertisement) (int, error) {
	id, err := a.advertisementRepo.CreateAdvertisement(ctx, advertisement)
	if err != nil {
		log.Fatalf("error: cannot create advertisement: %v\n", err.Error())
		return 0, err
	}
	return id, nil
}

func (a *AdvertisementService) GetAdvertisements(ctx context.Context) ([]entity.Advertisement, error) {
	panic("implement me")
}

func (a *AdvertisementService) GetAdvertisementById(ctx context.Context, id int) (entity.Advertisement, error) {
	panic("implement me")
}
