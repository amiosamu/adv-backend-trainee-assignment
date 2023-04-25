package repository

import (
	"context"
	"fmt"
	"github.com/amiosamu/adv-backend-trainee-assignment/advertising/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const adversaryTable = "advertising"

type AdvPostgres struct {
	db *sqlx.DB
}

func NewAdvPostgres(db *sqlx.DB) *AdvPostgres {
	return &AdvPostgres{
		db: db,
	}
}

func (a *AdvPostgres) Create(ctx context.Context, adversary model.Advertising) (uuid.UUID, error) {
	tx, err := a.db.BeginTxx(ctx, nil)
	if err != nil {
		return uuid.Nil, fmt.Errorf("adv repo - create -%w", err)
	}
	defer func(tx *sqlx.Tx) {
		err := tx.Rollback()
		if err != nil {
			return
		}
	}(tx)
	query := fmt.Sprintf("INSERT INTO %s (id, name, description, links, price) VALUES ($1, $2, $3, $4, $5) RETURNING id", adversaryTable)
	var id uuid.UUID
	if err := tx.GetContext(ctx, &id, query, adversary.Id, adversary.Description, adversary.Links, adversary.Price); err != nil {
		return uuid.Nil, fmt.Errorf("adv repo - create -%w", err)
	}
	defer func(tx *sqlx.Tx) {
		err := tx.Rollback()
		if err != nil {
			return
		}
	}(tx)
	return id, tx.Commit()

}

func (a *AdvPostgres) GetAll(ctx context.Context, uuid uuid.UUID) ([]model.Advertising, error) {
	var adversaries []model.Advertising
	tx, err := a.db.BeginTxx(ctx, nil)
	if err != nil {
		return adversaries, fmt.Errorf("adv repo - create - %w", err)
	}
	defer func(tx *sqlx.Tx) {
		err := tx.Rollback()
		if err != nil {
			return
		}
	}(tx)
	query := fmt.Sprintf("")
	if err := a.db.Select(&adversaries, query, uuid); err != nil {
		return adversaries, fmt.Errorf("adv repo - create - %w", err)
	}
	return adversaries, nil
}

func (a *AdvPostgres) GetByID(ctx context.Context, id uuid.UUID) (model.Advertising, error) {
	tx, err := a.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.Advertising{}, fmt.Errorf("adv repo - get by id - %w", err)
	}
	defer func(tx *sqlx.Tx) {
		err := tx.Rollback()
		if err != nil {
			return
		}
	}(tx)
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", adversaryTable)
	var entity model.Advertising
	if err := tx.GetContext(ctx, &entity, query, id); err != nil {
		return model.Advertising{}, fmt.Errorf("adv repo - get by id - %w", err)
	}
	return entity, tx.Commit()
}
