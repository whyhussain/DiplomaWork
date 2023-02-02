package service

import (
	"DiplomaWork/internal/app/model"
	"context"
)

type DiplomaService interface {
	GetAllRestaurants(ctx context.Context) (*model.Restaurant, error)
}
