package repository

import (
	"DiplomaWork/internal/app/model"
	"context"
)

type DiplomaRepository interface {
	FindAllRestaraunts(ctx context.Context) ([]*model.RestarauntsModel, error)
	NewRestaraunts(ctx context.Context, Name string, Category int) (string, error)
}
