package repository

import (
	"DiplomaWork/internal/app/model"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DimplomaServiceRepository struct {
	db *pgxpool.Pool
}

func NewDiplomaRepository(db *pgxpool.Pool) DiplomaRepository {
	return &DimplomaServiceRepository{db: db}
}

func (afr *DimplomaServiceRepository) FindAllRestaraunts(ctx context.Context) (*model.RestarauntsModel, error) {
	return nil, nil
}
