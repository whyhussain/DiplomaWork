package service

import (
	"DiplomaWork/internal/app/model"
	"context"
)

type DiplomaService interface {
	GetAllRestaurant(ctx context.Context) ([]*model.RestaurantsModel, error)
	GetRestaurantById(ctx context.Context, id int) (*model.RestaurantsModel, error)
	AddRestaurant(ctx context.Context, Name string, Category int) (string, error)
	UpdateRestaurant(ctx context.Context, restaurant *model.RestaurantsModel) error
	DeleteRestaurantById(ctx context.Context, id int) error

	GetCategories(ctx context.Context) ([]*model.Category, error)
	AddCategory(ctx context.Context, Type string) (string, error)

	GetAllMenu(ctx context.Context) ([]*model.Menu, error)
	GetMenuById(ctx context.Context, id int) (*model.Menu, error)
	AddMenu(ctx context.Context, Name string, RestaurantId int, Price int) (string, error)
	DeleteMenuById(ctx context.Context, id int) error
	UpdateMenu(ctx context.Context, menu *model.Menu) error
}
