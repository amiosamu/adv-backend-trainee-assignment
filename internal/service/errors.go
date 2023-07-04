package service

import "fmt"

var (
	ErrAdvertisementAlreadyExists = fmt.Errorf("advertisement already exists")
	ErrCannotCreateAdvertisement  = fmt.Errorf("cannot create advertisement")
	ErrAdvertisementNotFound      = fmt.Errorf("advertisement not found")
	ErrCannotGetAdvertisement     = fmt.Errorf("cannot get advertisement")
	ErrCannotGetAdvertisements    = fmt.Errorf("cannot get advertisements")
)
