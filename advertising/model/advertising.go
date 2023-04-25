package model

import "github.com/google/uuid"

type Advertising struct {
	Id          uuid.UUID `db:"uuid" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Links       []string  `db:"links" json:"links"`
	Price       int       `db:"price" json:"price"`
}
