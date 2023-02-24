package repository

import (
	"DiplomaWork/internal/app/model"
	"context"
)

type DiplomaRepository interface {
	FindAllRestaurants(ctx context.Context) ([]*model.RestaurantsModel, error)
	FindRestaurantById(ctx context.Context, id int) (*model.RestaurantsModel, error)
	AddRestaurants(ctx context.Context, Name string, Category int) (string, error)
	UpdateRestaurant(ctx context.Context, restaurant *model.RestaurantsModel) (*model.RestaurantsModel, error)
	DeleteRestaurantById(ctx context.Context, id int) error

	AllCategories(ctx context.Context) ([]*model.Category, error)
	AddCategory(ctx context.Context, Type string) (string, error)

	FindAllMenu(ctx context.Context) ([]*model.Menu, error)
	FindMenuById(ctx context.Context, id int) (*model.Menu, error)
	AddMenu(ctx context.Context, Name string, RestaurantId int, Price int) (string, error)
	DeleteMenuById(ctx context.Context, id int) error
	UpdateMenu(ctx context.Context, menu *model.Menu) (*model.Menu, error)
}
