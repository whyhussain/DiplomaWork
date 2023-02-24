package service

import (
	"DiplomaWork/internal/app/model"
	"DiplomaWork/internal/app/repository"
	"context"
)

type DiplomaServiceImpl struct {
	dipRepository repository.DiplomaRepository
}

func NewDiplomaService(repo repository.DiplomaRepository) DiplomaService {
	return &DiplomaServiceImpl{dipRepository: repo}
}

<<<<<<< Updated upstream
func (as *DiplomaServiceImpl) GetAllRestaurants(ctx context.Context) (*model.Restaurant, error) {
	restaurants, err := as.dipRepository.FindAllRestaurants(ctx)
	if err != nil {
		return nil, err
	}
	return restaurants, nil
=======
func (as *DiplomserviceImpl) GetAllRestaurant(ctx context.Context) ([]*model.RestaurantsModel, error) {
	rests, err := as.dipRepository.FindAllRestaurants(ctx)
	if err != nil {
		return nil, err
	}
	return rests, nil
}
func (as *DiplomserviceImpl) GetRestaurantById(ctx context.Context, id int) (*model.RestaurantsModel, error) {
	restaurant, err := as.dipRepository.FindRestaurantById(ctx, id)
	if err != nil {
		return nil, err
	}
	return restaurant, nil
}
func (as *DiplomserviceImpl) AddRestaurant(ctx context.Context, Name string, Category int) (string, error) {
	msg, err := as.dipRepository.AddRestaurants(ctx, Name, Category)
	if err != nil {
		return msg, err
	}
	return msg, nil
}
func (as *DiplomserviceImpl) UpdateRestaurant(ctx context.Context, restaurant *model.RestaurantsModel) error {
	_, err := as.dipRepository.UpdateRestaurant(ctx, restaurant)
	if err != nil {
		return err
	}
	return nil
}
func (as *DiplomserviceImpl) DeleteRestaurantById(ctx context.Context, id int) error {
	err := as.dipRepository.DeleteRestaurantById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func (as *DiplomserviceImpl) GetCategories(ctx context.Context) ([]*model.Category, error) {
	categories, err := as.dipRepository.AllCategories(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
>>>>>>> Stashed changes
}

func (as *DiplomserviceImpl) AddCategory(ctx context.Context, Type string) (string, error) {
	msg, err := as.dipRepository.AddCategory(ctx, Type)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func (as *DiplomserviceImpl) GetAllMenu(ctx context.Context) ([]*model.Menu, error) {
	menus, err := as.dipRepository.FindAllMenu(ctx)
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (as *DiplomserviceImpl) GetMenuById(ctx context.Context, id int) (*model.Menu, error) {
	menu, err := as.dipRepository.FindMenuById(ctx, id)
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (as *DiplomserviceImpl) AddMenu(ctx context.Context, Name string, RestaurantId int, Price int) (string, error) {
	msg, err := as.dipRepository.AddMenu(ctx, Name, RestaurantId, Price)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func (as *DiplomserviceImpl) DeleteMenuById(ctx context.Context, id int) error {
	err := as.dipRepository.DeleteMenuById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (as *DiplomserviceImpl) UpdateMenu(ctx context.Context, menu *model.Menu) error {
	_, err := as.dipRepository.UpdateMenu(ctx, menu)
	if err != nil {
		return err
	}
	return nil
}
