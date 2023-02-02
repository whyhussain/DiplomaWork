package repository

import (
	"DiplomaWork/internal/app/model"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DiplomaServiceRepository struct {
	db *pgxpool.Pool
}

func NewDiplomaRepository(db *pgxpool.Pool) DiplomaRepository {
	return &DiplomaServiceRepository{db: db}
}

func (afr *DiplomaServiceRepository) FindAllRestaurants(ctx context.Context) (*model.Restaurant, error) {
	return nil, nil
}
