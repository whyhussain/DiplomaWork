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

func (afr *DimplomaServiceRepository) FindAllRestaraunts(ctx context.Context) ([]*model.RestarauntsModel, error) {
	restaraunts := []*model.RestarauntsModel{}
	query := `select * from restaraunts`
	rows, err := afr.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&restaraunts)
	}
	return restaraunts, nil
}
func (afr *DimplomaServiceRepository) NewRestaraunts(ctx context.Context, Name string, Category int) (string, error) {
	query := `insert into restaraunts(label, category_type) VALUES ($1,$2)`
	tx, err := afr.db.Begin(ctx)
	if err != nil {
		return err.Error(), err
	}
	_, err = tx.Query(ctx, query, Name, Category)
	if err != nil {
		tx.Rollback(ctx)
		return err.Error(), err
	}
	tx.Commit(ctx)

	return "restaraunt is created", nil

}
