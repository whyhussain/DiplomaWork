package repository

import (
	"DiplomaWork/internal/app/model"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type DiplomaServiceRepository struct {
	db *pgxpool.Pool
}

func NewDiplomaRepository(db *pgxpool.Pool) DiplomaRepository {
	return &DiplomaServiceRepository{db: db}
}

func (afr *DiplomaServiceRepository) FindAllRestaurants(ctx context.Context) ([]*model.RestaurantsModel, error) {
	restaurants := []*model.RestaurantsModel{}

	query := `select r.id, r.restaurant_name, r.category_id, r.partner_id, r.address, r.city, r.price_of_service, r.restaurant_uin, r.phone_number, r.rating, r.schedule from restaurants r
	join category c on c.id = r.category_id`
	rows, err := afr.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rest := model.RestaurantsModel{}
		rows.Scan(&rest.Id, &rest.RestaurantName, &rest.CategoryID, &rest.PartnerId, &rest.Address, &rest.City, &rest.PriceOfService,
			&rest.RestaurantUIN, &rest.PhoneNumber, &rest.Rating, &rest.Schedule)
		restaurants = append(restaurants, &rest)
	}
	return restaurants, nil
}

func (afr *DiplomaServiceRepository) FindRestaurantById(ctx context.Context, id int) (*model.RestaurantsModel, error) {
	restaurant := &model.RestaurantsModel{}

	query := `SELECT * FROM restaurants WHERE id = $1`
	rows, err := afr.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&restaurant.Id, &restaurant.RestaurantName, &restaurant.CategoryID, &restaurant.PartnerId, &restaurant.Address, &restaurant.City, &restaurant.PriceOfService,
			&restaurant.RestaurantUIN, &restaurant.PhoneNumber, &restaurant.Rating, &restaurant.Schedule)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("restaurant with id %d not found", id)
	}

	return restaurant, nil
}

func (afr *DiplomaServiceRepository) AddRestaurants(ctx context.Context, RestaurantName string, CategoryID int, PartnerId int, Address string,
	City string, PriceOfService int, RestaurantUIN int, PhoneNumber string, Rating float64, Schedule []model.Schedule) (string, error) {
	query := ` SELECT * FROM restaurants WHERE restaurant_name =$1 AND category_id =$2 AND partner_id =$3 AND address =$4 AND city =$5 
	AND price_of_service =$6 AND restaurant_uin =$7 AND phone_number=$8 AND rating=$9 AND schedule=$10`
	rows, err := afr.db.Query(ctx, query, RestaurantName, CategoryID, PartnerId, Address, City, PriceOfService, RestaurantUIN, PhoneNumber, Rating, Schedule)
	if err != nil {
		return err.Error(), err
	}
	defer rows.Close()
	if rows.Next() {
		return "we have this restaurant", err
	}
	query = `insert into restaurants(restaurant_name, category_id, partner_id, address, city, price_of_service, restaurant_uin, phone_number, rating, schedule)
	SELECT $1, $2, $3, $4, $5, $6, $7, $8, $9, $10 where
    NOT EXISTS (
        SELECT * FROM restaurants WHERE restaurant_name =$11 AND category_id =$12 AND partner_id =$13 AND address =$14 AND city =$15 
		AND price_of_service =$16 AND restaurant_uin =$17 AND phone_number=$18 AND rating=$19 AND schedule=$20
    )`
	_, err = afr.db.Exec(ctx, query, RestaurantName, CategoryID, PartnerId, Address, City, PriceOfService, RestaurantUIN, PhoneNumber, Rating, Schedule, RestaurantName, CategoryID, PartnerId, Address, City, PriceOfService, RestaurantUIN, PhoneNumber, Rating, Schedule)
	if err != nil {
		return err.Error(), err
	}

	return "restaurant created", nil
}

func (afr *DiplomaServiceRepository) UpdateRestaurant(ctx context.Context, restaurant *model.RestaurantsModel) (*model.RestaurantsModel, error) {
	if afr.db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	query := `UPDATE restaurants SET restaurant_name=$1, category_id=$2, partner_id=$3, address=$4, city=$5, price_of_service=$6,
	 restaurant_uin=$7, phone_number=$8, rating=$9, schedule=$10 WHERE id = $11`
	_, err := afr.db.Exec(ctx, query, restaurant.RestaurantName, restaurant.CategoryID, restaurant.PartnerId, restaurant.Address, restaurant.City,
		restaurant.PriceOfService, restaurant.RestaurantUIN, restaurant.PhoneNumber, restaurant.Rating, restaurant.Schedule, restaurant.Id)
	if err != nil {
		return nil, err
	}

	return restaurant, nil
}

func (afr *DiplomaServiceRepository) DeleteRestaurantById(ctx context.Context, id int) error {
	query := `DELETE FROM restaurants WHERE id=$1`
	_, err := afr.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (afr *DiplomaServiceRepository) AllCategories(ctx context.Context) ([]*model.Category, error) {
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
}

func (afr *DiplomaServiceRepository) FindCategoryById(ctx context.Context, id int) (*model.Category, error) {
	category := &model.Category{}

	query := `SELECT * FROM category WHERE id = $1`
	rows, err := afr.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Type)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("category with id %d not found", id)
	}

	return category, nil
}

func (afr *DiplomaServiceRepository) AddCategory(ctx context.Context, Type string) (string, error) {
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

func (afr *DiplomaServiceRepository) UpdateCategory(ctx context.Context, category *model.Category) (*model.Category, error) {
	if afr.db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	query := `UPDATE category SET type=$1 WHERE id = $2`
	_, err := afr.db.Exec(ctx, query, category.Type, category.Id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (afr *DiplomaServiceRepository) DeleteCategory(ctx context.Context, id int) error {
	query := `DELETE FROM category WHERE id=$1`
	_, err := afr.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (afr *DiplomaServiceRepository) FindAllMenu(ctx context.Context) ([]*model.Menu, error) {
	menus := []*model.Menu{}

	query := `SELECT m.id, m.name, m.category_id, m.restaurant_id, m.description, m.price
	FROM menus m
	JOIN restaurants r ON r.id = m.restaurant_id`
	rows, err := afr.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		menu := model.Menu{}
		rows.Scan(&menu.Id, &menu.Name, &menu.CategoryId, &menu.RestaurantId, &menu.Description, &menu.Price)
		if err != nil {
			return nil, err
		}
		menus = append(menus, &menu)
	}
	return menus, nil
}

func (afr *DiplomaServiceRepository) FindMenuById(ctx context.Context, id int) (*model.Menu, error) {
	menu := &model.Menu{}

	query := `SELECT * FROM menus WHERE id = $1`
	rows, err := afr.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&menu.Id, &menu.Name, &menu.CategoryId, &menu.RestaurantId, &menu.Description, &menu.Price)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("menu with id %d not found", id)
	}

	return menu, nil
}

func (afr *DiplomaServiceRepository) AddMenu(ctx context.Context, Name string, CategoryId int, RestaurantId int, Description string, Price int32) (string, error) {
	query := `SELECT * FROM menus WHERE name=$1 AND category_id=$2 AND restaurant_id=$3 AND description=$4 AND price=$5`
	rows, err := afr.db.Query(ctx, query, Name, CategoryId, RestaurantId, Description, Price)
	if err != nil {
		return err.Error(), err
	}
	defer rows.Close()
	if rows.Next() {
		return "we have this menu", nil
	}
	query = `INSERT INTO menus(name, category_id, restaurant_id, description, price)
		SELECT $1, $2, $3, $4, $5 WHERE NOT EXISTS (
			SELECT * FROM menus WHERE name=$6 AND category_id=$7 AND restaurant_id=$8 AND description=$9 AND price=$10
		)`
	_, err = afr.db.Exec(ctx, query, Name, CategoryId, RestaurantId, Description, Price, Name, CategoryId, RestaurantId, Description, Price)
	if err != nil {
		return err.Error(), err
	}

	return "menu created", nil

}
func (afr *DiplomaServiceRepository) DeleteMenuById(ctx context.Context, id int) error {
	query := `DELETE FROM menus WHERE id=$1`
	_, err := afr.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (afr *DiplomaServiceRepository) UpdateMenu(ctx context.Context, menu *model.Menu) (*model.Menu, error) {
	if afr.db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	query := `UPDATE menus SET name = $1, restaurant_id = $2, category_id=$3, description=$4,price = $5 WHERE id = $6`
	_, err := afr.db.Exec(ctx, query, menu.Name, menu.RestaurantId, menu.CategoryId, menu.Description, menu.Price, menu.Id)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (afr *DiplomaServiceRepository) FindAllPartners(ctx context.Context) ([]*model.Partner, error) {
	partners := []*model.Partner{}

	query := `select id,name,email,password from partners`
	rows, err := afr.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		partner := model.Partner{}
		rows.Scan(&partner.Id, &partner.Name, &partner.Email, &partner.Password)
		partners = append(partners, &partner)
	}
	return partners, nil
}

func (afr *DiplomaServiceRepository) FindPartnerById(ctx context.Context, id int) (*model.Partner, error) {
	partner := &model.Partner{}

	query := `SELECT * FROM partners WHERE id = $1`
	rows, err := afr.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&partner.Id, &partner.Name, &partner.Email, &partner.Password)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("partner with id %d not found", id)
	}

	return partner, nil
}

func (afr *DiplomaServiceRepository) AddPartner(ctx context.Context, Name string, Email string, Password string) (string, error) {
	query := ` SELECT name, email, password FROM partners WHERE name=$1 AND email=$2 AND password=$3`
	rows, err := afr.db.Query(ctx, query, Name, Email, Password)
	if rows.Next() {
		return "we have this partner", err
	}
	query = `insert into partners(name, email, password)SELECT $1,$2,$3 where
    NOT EXISTS (
        SELECT name, email, password FROM partners WHERE name=$4 AND email=$5 AND password=$6
    );`
	_, err = afr.db.Query(ctx, query, Name, Email, Password, Name, Email, Password)
	if err != nil {
		return err.Error(), err
	}

	return "partner created", nil

}
func (afr *DiplomaServiceRepository) UpdatePartnerById(ctx context.Context, partner *model.Partner) (*model.Partner, error) {
	if afr.db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	query := `UPDATE partners SET name=$1, email=$2, password=$3 WHERE id = $4`
	_, err := afr.db.Exec(ctx, query, partner.Name, partner.Email, partner.Password, partner.Id)
	if err != nil {
		return nil, err
	}

	return partner, nil
}
func (afr *DiplomaServiceRepository) DeletePartnerById(ctx context.Context, id int) error {
	query := `DELETE FROM partners WHERE id=$1`
	_, err := afr.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (afr *DiplomaServiceRepository) FindAllAdmins(ctx context.Context) ([]*model.Admin, error) {
	admins := []*model.Admin{}

	query := `select id,name,email,password from admins`
	rows, err := afr.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		admin := model.Admin{}
		rows.Scan(&admin.Id, &admin.Name, &admin.Email, &admin.Password)
		admins = append(admins, &admin)
	}
	return admins, nil
}
func (afr *DiplomaServiceRepository) FindAdminById(ctx context.Context, id int) (*model.Admin, error) {
	admin := &model.Admin{}

	query := `SELECT * FROM admins WHERE id = $1`
	rows, err := afr.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&admin.Id, &admin.Name, &admin.Email, &admin.Password)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("admin with id %d not found", id)
	}

	return admin, nil
}
func (afr *DiplomaServiceRepository) AddAdmin(ctx context.Context, Name string, Email string, Password string) (string, error) {
	query := ` SELECT name, email, password FROM admins WHERE name=$1 AND email=$2 AND password=$3`
	rows, err := afr.db.Query(ctx, query, Name, Email, Password)
	if rows.Next() {
		return "we have this admin", err
	}
	query = `insert into admins(name, email, password)SELECT $1,$2,$3 where
    NOT EXISTS (
        SELECT name, email, password FROM admins WHERE name=$4 AND email=$5 AND password=$6
    );`
	_, err = afr.db.Query(ctx, query, Name, Email, Password, Name, Email, Password)
	if err != nil {
		return err.Error(), err
	}

	return "admin created", nil

}
func (afr *DiplomaServiceRepository) UpdateAdminById(ctx context.Context, admin *model.Admin) (*model.Admin, error) {
	if afr.db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	query := `UPDATE admins SET name=$1, email=$2, password=$3 WHERE id = $4`
	_, err := afr.db.Exec(ctx, query, admin.Name, admin.Email, admin.Password, admin.Id)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (afr *DiplomaServiceRepository) DeleteAdminById(ctx context.Context, id int) error {
	query := `DELETE FROM admins WHERE id=$1`
	_, err := afr.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (afr *DiplomaServiceRepository) FindAllTechSupports(ctx context.Context) ([]*model.TechSupport, error) {
	techSupports := []*model.TechSupport{}

	query := `select id,name,email,password from tech_supports`
	rows, err := afr.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		techSupport := model.TechSupport{}
		rows.Scan(&techSupport.Id, &techSupport.Name, &techSupport.Email, &techSupport.Password)
		techSupports = append(techSupports, &techSupport)
	}
	return techSupports, nil
}

func (afr *DiplomaServiceRepository) FindTechSupportById(ctx context.Context, id int) (*model.TechSupport, error) {
	techSupport := &model.TechSupport{}

	query := `SELECT * FROM tech_supports WHERE id = $1`
	rows, err := afr.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&techSupport.Id, &techSupport.Name, &techSupport.Email, &techSupport.Password)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("tech support with id %d not found", id)
	}

	return techSupport, nil
}

func (afr *DiplomaServiceRepository) AddTechSupport(ctx context.Context, Name string, Email string, Password string) (string, error) {
	query := ` SELECT name, email, password FROM tech_supports WHERE name=$1 AND email=$2 AND password=$3`
	rows, err := afr.db.Query(ctx, query, Name, Email, Password)
	if rows.Next() {
		return "we have this tech support", err
	}
	query = `insert into tech_supports(name, email, password)SELECT $1,$2,$3 where
    NOT EXISTS (
        SELECT name, email, password FROM tech_supports WHERE name=$4 AND email=$5 AND password=$6
    );`
	_, err = afr.db.Query(ctx, query, Name, Email, Password, Name, Email, Password)
	if err != nil {
		return err.Error(), err
	}

	return "tech support created", nil

}
func (afr *DiplomaServiceRepository) UpdateTechSupportById(ctx context.Context, techSupport *model.TechSupport) (*model.TechSupport, error) {
	if afr.db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	query := `UPDATE tech_supports SET name=$1, email=$2, password=$3 WHERE id = $4`
	_, err := afr.db.Exec(ctx, query, techSupport.Name, techSupport.Email, techSupport.Password, techSupport.Id)
	if err != nil {
		return nil, err
	}

	return techSupport, nil
}

func (afr *DiplomaServiceRepository) DeleteTechSupportById(ctx context.Context, id int) error {
	query := `DELETE FROM tech_supports WHERE id=$1`
	_, err := afr.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (afr *DiplomaServiceRepository) FindAllCustomers(ctx context.Context) ([]*model.Customer, error) {
	customers := []*model.Customer{}

	query := `select id,name,email,password, delivery_address, city, birthdate from customers`
	rows, err := afr.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		customer := model.Customer{}
		rows.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Password, &customer.DeliveryAddress, &customer.City, &customer.Birthdate)
		customers = append(customers, &customer)
	}
	return customers, nil
}

func (afr *DiplomaServiceRepository) FindCustomerById(ctx context.Context, id int) (*model.Customer, error) {
	customer := &model.Customer{}

	query := `SELECT * FROM customers WHERE id = $1`
	rows, err := afr.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Password, &customer.DeliveryAddress, &customer.City, &customer.Birthdate)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("customer with id %d not found", id)
	}

	return customer, nil
}

func (afr *DiplomaServiceRepository) AddCustomer(ctx context.Context, Name string, Email string, Password string, DeliveryAddress string, City string, Birthdate time.Time) (string, error) {
	query := ` SELECT name, email, password, delivery_address, city, birthdate FROM customers WHERE name=$1 AND email=$2 AND password=$3 AND delivery_address=$4 AND city=$5 AND birthdate=$6`
	rows, err := afr.db.Query(ctx, query, Name, Email, Password, DeliveryAddress, City, Birthdate)
	if rows.Next() {
		return "we have this customer", err
	}
	query = `insert into customers(name, email, password, delivery_address, city, birthdate)SELECT $1,$2,$3,$4,$5,$6 where
    NOT EXISTS (
        SELECT name, email, password, delivery_address, city, birthdate FROM customers WHERE name=$7 AND email=$8 AND password=$9 AND delivery_address=$10 AND city=$11 AND birthdate=$12
    );`
	_, err = afr.db.Query(ctx, query, Name, Email, Password, DeliveryAddress, City, Birthdate, Name, Email, Password, DeliveryAddress, City, Birthdate)
	if err != nil {
		return err.Error(), err
	}

	return "customer created", nil

}
func (afr *DiplomaServiceRepository) UpdateCustomerById(ctx context.Context, customer *model.Customer) (*model.Customer, error) {
	if afr.db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	query := `UPDATE customers SET name=$1, email=$2, password=$3, delivery_address=$4, city=$5, birthdate=$6 WHERE id = $7`
	_, err := afr.db.Exec(ctx, query, customer.Name, customer.Email, customer.Password, customer.DeliveryAddress, customer.City, customer.Birthdate, customer.Id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (afr *DiplomaServiceRepository) DeleteCustomerById(ctx context.Context, id int) error {
	query := `DELETE FROM customers WHERE id=$1`
	_, err := afr.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (afr *DiplomaServiceRepository) FindAllReviews(ctx context.Context) ([]*model.Review, error) {
	reviews := []*model.Review{}

	query := `select id, customer_id, restaurant_id, menu_id, point, review, date from reviews`
	rows, err := afr.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		review := model.Review{}
		rows.Scan(&review.Id, &review.CustomerId, &review.RestaurantId, &review.MenuId, &review.Point, &review.Review, &review.Date)
		reviews = append(reviews, &review)
	}
	return reviews, nil
}

func (afr *DiplomaServiceRepository) FindReviewById(ctx context.Context, id int) (*model.Review, error) {
	review := &model.Review{}

	query := `SELECT * FROM reviews WHERE id = $1`
	rows, err := afr.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&review.Id, &review.CustomerId, &review.RestaurantId, &review.MenuId, &review.Point, &review.Review, &review.Date)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("review with id %d not found", id)
	}

	return review, nil
}

func (afr *DiplomaServiceRepository) AddReview(ctx context.Context, CustomerId int, RestaurantId int, MenuId int, Point int, Review string, Date time.Time) (string, error) {
	query := ` SELECT customer_id, restaurant_id, menu_id, point, review, date FROM reviews WHERE customer_id=$1 AND restaurant_id=$2 AND menu_id=$3 AND point=$4 AND review=$5 AND date=$6`
	rows, err := afr.db.Query(ctx, query, CustomerId, RestaurantId, MenuId, Point, Review, Date)
	if rows.Next() {
		return "we have this review", err
	}
	query = `insert into reviews(customer_id, restaurant_id, menu_id, point, review, date)SELECT $1,$2,$3,$4,$5,$6 where
    NOT EXISTS (
        SELECT customer_id, restaurant_id, menu_id, point, review, date FROM reviews WHERE customer_id=$7 AND restaurant_id=$8 AND menu_id=$9 AND point=$10 AND review=$11 AND date=$12
    );`
	_, err = afr.db.Query(ctx, query, CustomerId, RestaurantId, MenuId, Point, Review, Date, CustomerId, RestaurantId, MenuId, Point, Review, Date)
	if err != nil {
		return err.Error(), err
	}

	return "review created", nil

}

func (afr *DiplomaServiceRepository) UpdateReviewById(ctx context.Context, review *model.Review) (*model.Review, error) {
	if afr.db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	query := `UPDATE reviews SET customer_id=$1, restaurant_id=$2, menu_id=$3, point=$4, review=$5, date=$6 WHERE id = $7`
	_, err := afr.db.Exec(ctx, query, review.CustomerId, review.RestaurantId, review.MenuId, review.Point, review.Review, review.Date, review.Id)
	if err != nil {
		return nil, err
	}

	return review, nil
}

func (afr *DiplomaServiceRepository) DeleteReviewById(ctx context.Context, id int) error {
	query := `DELETE FROM reviews WHERE id=$1`
	_, err := afr.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (afr *DiplomaServiceRepository) FindAllSchedules(ctx context.Context) ([]*model.Schedule, error) {
	schedules := []*model.Schedule{}

	query := `select id, day_of_week, opening_time, closing_time from schedules`
	rows, err := afr.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		schedule := model.Schedule{}
		rows.Scan(&schedule.Id, &schedule.DayOfWeek, &schedule.OpeningTime, &schedule.ClosingTime)
		schedules = append(schedules, &schedule)
	}
	return schedules, nil
}

func (afr *DiplomaServiceRepository) FindScheduleById(ctx context.Context, id int) (*model.Schedule, error) {
	schedule := &model.Schedule{}

	query := `SELECT * FROM schedules WHERE id = $1`
	rows, err := afr.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&schedule.Id, &schedule.DayOfWeek, &schedule.OpeningTime, &schedule.ClosingTime)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("schedule with id %d not found", id)
	}

	return schedule, nil
}

func (afr *DiplomaServiceRepository) AddSchedule(ctx context.Context, DayOfWeek string, OpeningTime time.Time, ClosingTime time.Time) (string, error) {
	query := ` SELECT day_of_week, opening_time, closing_time FROM schedules WHERE day_of_week=$1 AND opening_time=$2 AND closing_time=$3`
	rows, err := afr.db.Query(ctx, query, DayOfWeek, OpeningTime, ClosingTime)
	if rows.Next() {
		return "we have this schedule", err
	}
	query = `insert into schedules(day_of_week, opening_time, closing_time)SELECT $1,$2,$3 where
    NOT EXISTS (
        SELECT day_of_week, opening_time, closing_time FROM schedules WHERE day_of_week=$4 AND opening_time=$5 AND closing_time=$6
    );`
	_, err = afr.db.Query(ctx, query, DayOfWeek, OpeningTime, ClosingTime, DayOfWeek, OpeningTime, ClosingTime)
	if err != nil {
		return err.Error(), err
	}

	return "schedule created", nil

}

func (afr *DiplomaServiceRepository) UpdateScheduleById(ctx context.Context, schedule *model.Schedule) (*model.Schedule, error) {
	if afr.db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	query := `UPDATE schedules SET day_of_week=$1, opening_time=$2, closing_time=$3 WHERE id = $4`
	_, err := afr.db.Exec(ctx, query, schedule.DayOfWeek, schedule.OpeningTime, schedule.ClosingTime, schedule.Id)
	if err != nil {
		return nil, err
	}

	return schedule, nil
}

func (afr *DiplomaServiceRepository) DeleteScheduleById(ctx context.Context, id int) error {
	query := `DELETE FROM schedules WHERE id=$1`
	_, err := afr.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (afr *DiplomaServiceRepository) FindAllDeliveryPersonnels(ctx context.Context) ([]*model.DeliveryPersonnel, error) {
	deliveryPersonnels := []*model.DeliveryPersonnel{}

	query := `select * from delivery_personnel`
	rows, err := afr.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		deliveryPersonnel := model.DeliveryPersonnel{}
		rows.Scan(&deliveryPersonnel.Id, &deliveryPersonnel.Name, &deliveryPersonnel.Email, &deliveryPersonnel.Password, &deliveryPersonnel.AvailabilityStatus)
		deliveryPersonnels = append(deliveryPersonnels, &deliveryPersonnel)
	}
	return deliveryPersonnels, nil
}

func (afr *DiplomaServiceRepository) FindDeliveryPersonnelById(ctx context.Context, id int) (*model.DeliveryPersonnel, error) {
	deliveryPersonnel := &model.DeliveryPersonnel{}

	query := `SELECT * FROM delivery_personnel WHERE id = $1`
	rows, err := afr.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var availabilityStatus string
		err := rows.Scan(&deliveryPersonnel.Id, &deliveryPersonnel.Name, &deliveryPersonnel.Email, &deliveryPersonnel.Password, &availabilityStatus)
		switch availabilityStatus {
		case "Available":
			deliveryPersonnel.AvailabilityStatus = model.Available
		case "Busy":
			deliveryPersonnel.AvailabilityStatus = model.Busy
		case "Offline":
			deliveryPersonnel.AvailabilityStatus = model.Offline
		default:
			fmt.Println("unknown availability status:", availabilityStatus)
		}
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("delivery personnel with id %d not found", id)
	}

	return deliveryPersonnel, nil
}

func (afr *DiplomaServiceRepository) AddDeliveryPersonnel(ctx context.Context, Personel *model.DeliveryPersonnel, AvailabilityStatus model.DeliveryPersonnelAvailability) (string, error) {
	query := ` SELECT name, email, password, availability_status from delivery_personnel WHERE name=$1 AND email=$2 AND password=$3 AND availability_status=$4`

	tx, txErr := afr.db.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	})
	if txErr != nil {
		errMsg := fmt.Sprintf("Cannot start transaction")
		return errMsg, txErr
	}
	rows, err := tx.Query(ctx, query, Personel.Name, Personel.Email, Personel.Password, AvailabilityStatus)
	if err != nil {
		tx.Rollback(ctx)
	}
	if rows.Next() {
		return "we have this delivery_personnel", err
	}
	query = `insert into delivery_personnel(name, email, password, availability_status)SELECT $1,$2,$3,$4 where
    NOT EXISTS (
        SELECT name, email, password, availability_status FROM delivery_personnel WHERE name=$5 AND email=$6 AND password=$7 AND availability_status=$8
    );`
	_, err = tx.Query(ctx, query, Personel.Name, Personel.Email, Personel.Password, AvailabilityStatus, Personel.Name, Personel.Email, Personel.Password, AvailabilityStatus)
	if err != nil {
		return err.Error(), err
		tx.Rollback(ctx)
	}
	tx.Commit(ctx)

	return "delivery_personnel created", nil
}
func (afr *DiplomaServiceRepository) UpdateDeliveryPersonnelById(ctx context.Context, deliveryPersonnel *model.DeliveryPersonnel) (*model.DeliveryPersonnel, error) {
	if afr.db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	query := `UPDATE delivery_personnel SET name=$1, email=$2, password=$3, availability_status=$4 WHERE id = $5`
	_, err := afr.db.Exec(ctx, query, deliveryPersonnel.Name, deliveryPersonnel.Email, deliveryPersonnel.Password, deliveryPersonnel.AvailabilityStatus, deliveryPersonnel.Id)
	if err != nil {
		return nil, err
	}

	return deliveryPersonnel, nil
}

func (afr *DiplomaServiceRepository) DeleteDeliveryPersonnelById(ctx context.Context, id int) error {
	query := `DELETE FROM delivery_personnel WHERE id=$1`
	_, err := afr.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
