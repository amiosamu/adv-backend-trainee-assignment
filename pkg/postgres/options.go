package postgres

import "time"

type Option func(*Postgres)

func MaxPoolSize(poolSize int) Option {
	return func(c *Postgres) {
		c.maxPoolSize = poolSize
	}
}

func MaxConnAttempts(connAttempts int) Option {
	return func(c *Postgres) {
		c.connAttempts = connAttempts
	}
}

func ConnTimeout(connTimeout time.Duration) Option {
	return func(c *Postgres) {
		c.connTimeout = connTimeout
	}
}
