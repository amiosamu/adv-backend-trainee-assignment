package app

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

const (
	defaultAttempts = 20
	defaultTimeout  = time.Second
)

var (
	attempts = defaultAttempts
	err      error
	m        *migrate.Migrate
)

func init() {
	databaseURI, ok := os.LookupEnv("PG_URI")
	if !ok || len(databaseURI) == 0 {
		log.Fatalf("migrate: environment variable not declared: PG_URI")
	}
	databaseURI += "?sslmode=disable"

	for attempts > 0 {
		m, err = migrate.New("file:///migrations", databaseURI)
		if err == nil {
			break
		}
		log.Printf("Migrate: pgdb is trying to connect, attempts left: %d", attempts)
		time.Sleep(defaultTimeout)
		attempts--
	}
	err = m.Up()
	defer func() {
		_, _ = m.Close()
	}()
	if err != nil && errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migrate: up error: %s", err)
	}
	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("Migrate: no change")
		return
	}
	log.Printf("Migrate: up success!")
}
