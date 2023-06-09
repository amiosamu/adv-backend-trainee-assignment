package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

type DataSource struct {
	DB *sqlx.DB
}

func InitDB() (*DataSource, error) {
	pgHost := os.Getenv("POSTGRES_HOST")
	pgPort := os.Getenv("POSTGRES_PORT")
	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	pgDB := os.Getenv("POSTGRES_DB")
	pgSSL := os.Getenv("PG_SSL")
	pgConnString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", pgHost, pgPort, pgUser, pgPassword, pgDB, pgSSL)
	db, err := sqlx.Open("postgres", pgConnString)
	if err != nil {
		return nil, fmt.Errorf("erorr connecting to db %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to db %w", err)
	}
	return &DataSource{
		DB: db,
	}, nil
}

func (d *DataSource) Close() error {
	if err := d.DB.Close(); err != nil {
		return fmt.Errorf("error closing Postgresql: %w", err)
	}
	return nil
}
