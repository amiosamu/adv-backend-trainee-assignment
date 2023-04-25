package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

type DataSources struct {
	DB *sqlx.DB
}

func InitDB() (*DataSources, error) {
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	db := os.Getenv("PG_DB")
	ssl := os.Getenv("PG_SSL")
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, db, ssl)
	database, err := sqlx.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database %w", err)
	}
	if err := database.Ping(); err != nil {
		return nil, database.Ping()
	}
	return &DataSources{
		DB: database,
	}, nil
}

func Close(d *DataSources) error {
	if err := d.DB.Close(); err != nil {
		return fmt.Errorf("error closing Database: %w", err)
	}
	return nil
}
