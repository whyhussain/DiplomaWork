package repository

import (
	"DiplomaWork/internal/app/model"
	"context"
)

type DiplomaRepository interface {
	FindAllRestaurants(ctx context.Context) (*model.Restaurant, error)
}
