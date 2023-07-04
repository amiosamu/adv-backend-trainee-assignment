package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/entity"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/repo"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/repo/repoerrors"
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
	advertisements, err := a.advertisementRepo.GetAdvertisements(ctx)
	if err != nil {
		log.Printf("Failed to retrieve advertisements: %v", err)
		return nil, ErrCannotGetAdvertisement
	}
	return advertisements, nil
}

func (a *AdvertisementService) GetAdvertisementById(ctx context.Context, id int) (entity.Advertisement, error) {
	advertisement, err := a.advertisementRepo.GetAdvertisementById(ctx, id)
	if err != nil {

		fmt.Printf("Failed to retrieve advertisement with ID %d: %s\n", id, err.Error())

		if errors.Is(err, repoerrors.ErrNotFound) {
			return advertisement, ErrAdvertisementNotFound
		}

		return advertisement, ErrCannotGetAdvertisement
	}

	return advertisement, nil
}
