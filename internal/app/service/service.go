package service

import (
	"DiplomaWork/internal/app/model"
	"context"
)

type DiplomaService interface {
	GetAllRestaraunt(ctx context.Context) ([]*model.RestarauntsModel, error)
	NewRestaraunt(ctx context.Context, Name string, Category int) (string, error)
	GetCattegories(ctx context.Context) ([]*model.Category, error)
}
