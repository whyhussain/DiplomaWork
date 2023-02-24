package repository

import (
	"DiplomaWork/internal/app/model"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DiplomaServiceRepository struct {
	db *pgxpool.Pool
}

func NewDiplomaRepository(db *pgxpool.Pool) DiplomaRepository {
	return &DiplomaServiceRepository{db: db}
}

<<<<<<< Updated upstream
func (afr *DiplomaServiceRepository) FindAllRestaurants(ctx context.Context) (*model.Restaurant, error) {
	return nil, nil
=======
func (afr *DimplomaServiceRepository) FindAllRestaurants(ctx context.Context) ([]*model.RestaurantsModel, error) {
	restaurants := []*model.RestaurantsModel{}

	query := `select r.id, r.label,c.type from restaurants r
	join category c on c.id = r.category_type`
	rows, err := afr.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rest := model.RestaurantsModel{}
		rows.Scan(&rest.Id, &rest.RestaurantName, &rest.RestaurantCategory)
		restaurants = append(restaurants, &rest)
	}
	return restaurants, nil
}

func (afr *DimplomaServiceRepository) FindRestaurantById(ctx context.Context, id int) (*model.RestaurantsModel, error) {
	restaurant := &model.RestaurantsModel{}

	query := `SELECT id, label, category_type FROM restaurants WHERE id = $1`
	rows, err := afr.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&restaurant.Id, &restaurant.RestaurantName, &restaurant.CategoryID)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("restaurant with id %d not found", id)
	}

	return restaurant, nil
}

func (afr *DimplomaServiceRepository) AddRestaurants(ctx context.Context, Name string, Category int) (string, error) {
	query := ` SELECT label,category_type FROM restaurants WHERE label =$1 and category_type =$2`
	rows, err := afr.db.Query(ctx, query, Name, Category)
	if rows.Next() {
		return "we have this restaurant", err
	}
	query = `insert into restaurants(label, category_type)SELECT $1, $2 where
    NOT EXISTS (
        SELECT label,category_type FROM restaurants WHERE label =$3 and category_type =$4 
    );`
	_, err = afr.db.Query(ctx, query, Name, Category, Name, Category)
	if err != nil {
		return err.Error(), err
	}

	return "restaurant created", nil

}

func (afr *DimplomaServiceRepository) UpdateRestaurant(ctx context.Context, restaurant *model.RestaurantsModel) (*model.RestaurantsModel, error) {
	if afr.db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	query := `UPDATE restaurants SET label = $1, category_type = $2 WHERE id = $3`
	_, err := afr.db.Exec(ctx, query, restaurant.RestaurantName, restaurant.CategoryID, restaurant.Id)
	if err != nil {
		return nil, err
	}

	return restaurant, nil
}

func (afr *DimplomaServiceRepository) DeleteRestaurantById(ctx context.Context, id int) error {
	query := `DELETE FROM restaurants WHERE id=$1`
	_, err := afr.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (afr *DimplomaServiceRepository) AllCategories(ctx context.Context) ([]*model.Category, error) {
	categories := []*model.Category{}

	query := `select id,type from category`
	rows, err := afr.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		category := model.Category{}
		rows.Scan(&category.Id, &category.Type)
		categories = append(categories, &category)
	}
	return categories, nil
>>>>>>> Stashed changes
}

func (afr *DimplomaServiceRepository) AddCategory(ctx context.Context, Type string) (string, error) {
	query := ` SELECT type FROM category WHERE type =$1`
	rows, err := afr.db.Query(ctx, query, Type)
	if rows.Next() {
		return "we have this category", err
	}
	query = `insert into category(type)SELECT $1 where
    NOT EXISTS (
        SELECT type FROM category WHERE type =$2  
    );`
	_, err = afr.db.Query(ctx, query, Type, Type)
	if err != nil {
		return err.Error(), err
	}

	return "category created", nil

}

func (afr *DimplomaServiceRepository) FindAllMenu(ctx context.Context) ([]*model.Menu, error) {
	menus := []*model.Menu{}

	query := `select m.id, m.name, m.restaurant_id, m.price from menu m
	join restaurants r on r.id = m.restaurant_id`
	rows, err := afr.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		menu := model.Menu{}
		rows.Scan(&menu.Id, &menu.Name, &menu.RestaurantId, &menu.Price)
		menus = append(menus, &menu)
	}
	return menus, nil
}

func (afr *DimplomaServiceRepository) FindMenuById(ctx context.Context, id int) (*model.Menu, error) {
	menu := &model.Menu{}

	query := `SELECT id, name, restaurant_id, price FROM menu WHERE id = $1`
	rows, err := afr.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&menu.Id, &menu.Name, &menu.RestaurantId, &menu.Price)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("menu with id %d not found", id)
	}

	return menu, nil
}

func (afr *DimplomaServiceRepository) AddMenu(ctx context.Context, Name string, RestaurantId int, Price int) (string, error) {
	query := ` SELECT name, restaurant_id, price FROM menu WHERE name=$1 AND restaurant_id=$2 AND price=$3`
	rows, err := afr.db.Query(ctx, query, Name, RestaurantId, Price)
	if rows.Next() {
		return "we have this menu", err
	}
	query = `insert into menu(name, restaurant_id, price)
	SELECT $1, $2, $3 where
    NOT EXISTS (
        SELECT name, restaurant_id, price FROM menu WHERE name=$4 AND restaurant_id=$5 AND price=$6
    );`
	_, err = afr.db.Query(ctx, query, Name, RestaurantId, Price, Name, RestaurantId, Price)
	if err != nil {
		return err.Error(), err
	}

	return "menu created", nil

}

func (afr *DimplomaServiceRepository) DeleteMenuById(ctx context.Context, id int) error {
	query := `DELETE FROM menu WHERE id=$1`
	_, err := afr.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (afr *DimplomaServiceRepository) UpdateMenu(ctx context.Context, menu *model.Menu) (*model.Menu, error) {
	if afr.db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	query := `UPDATE menu SET name = $1, restaurant_id = $2, price = $3 WHERE id = $4`
	_, err := afr.db.Exec(ctx, query, menu.Name, menu.RestaurantId, menu.Price, menu.Id)
	if err != nil {
		return nil, err
	}

	return menu, nil
}
