package service

import "fmt"

var (
	ErrAdvertisementNotFound  = fmt.Errorf("advertisement not found")
	ErrCannotGetAdvertisement = fmt.Errorf("cannot get advertisement")
)
