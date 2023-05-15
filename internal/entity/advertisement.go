package entity

import "time"

type Advertisement struct {
	Id          int       `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Pictures    []string  `db:"pictures"`
	Price       int       `db:"price"`
	CreatedAt   time.Time `db:"created_at"`
}
