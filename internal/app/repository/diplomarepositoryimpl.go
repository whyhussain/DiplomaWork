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

	query := `select r.label,c.name from restaraunts r
	join category c on c.id = r.category_type`
	rows, err := afr.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rest := model.RestarauntsModel{}
		rows.Scan(&rest.RestarauntName, &rest.RestarauntCategory)
		restaraunts = append(restaraunts, &rest)
	}
	return restaraunts, nil
}
func (afr *DimplomaServiceRepository) NewRestaraunts(ctx context.Context, Name string, Category int) (string, error) {
	query := ` SELECT label,category_type FROM restaraunts WHERE label =$1 and category_type =$2`
	rows, err := afr.db.Query(ctx, query, Name, Category)
	if rows.Next() != false {
		return "we have this restaraunt", nil
	}
	query = `insert into restaraunts(label, category_type)SELECT $1, $2 where
    NOT EXISTS (
        SELECT label,category_type FROM restaraunts WHERE label =$3 and category_type =$4 
    );`
	_, err = afr.db.Query(ctx, query, Name, Category, Name, Category)
	if err != nil {
		return err.Error(), err
	}

	return "restaraunt created", nil

}
func (afr *DimplomaServiceRepository) AllCategories(ctx context.Context) ([]*model.Category, error) {
	categories := []*model.Category{}

	query := `select id,name from category`
	rows, err := afr.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		category := model.Category{}
		rows.Scan(&category.Id, &category.Name)
		categories = append(categories, &category)
	}
	return categories, nil
}
