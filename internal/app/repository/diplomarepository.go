package repository

import (
	"DiplomaWork/internal/app/model"
	"context"
	"time"
)

type DiplomaRepository interface {
	FindAllRestaurants(ctx context.Context) ([]*model.RestaurantsModel, error)
	FindRestaurantById(ctx context.Context, id int) (*model.RestaurantsModel, error)
	AddRestaurants(ctx context.Context, RestaurantName string, CategoryID int, PartnerId int, Address string,
		City string, PriceOfService int, RestaurantUIN int, PhoneNumber string, Rating float64, Schedule []model.Schedule) (string, error)
	UpdateRestaurant(ctx context.Context, restaurant *model.RestaurantsModel) (*model.RestaurantsModel, error)
	DeleteRestaurantById(ctx context.Context, id int) error

	AllCategories(ctx context.Context) ([]*model.Category, error)
	FindCategoryById(ctx context.Context, id int) (*model.Category, error)
	AddCategory(ctx context.Context, Type string) (string, error)
	UpdateCategory(ctx context.Context, category *model.Category) (*model.Category, error)
	DeleteCategory(ctx context.Context, id int) error

	FindAllMenu(ctx context.Context) ([]*model.Menu, error)
	FindMenuById(ctx context.Context, id int) (*model.Menu, error)
	AddMenu(ctx context.Context, Name string, CategoryId int, RestaurantId int, Description string, Price int32) (string, error)
	DeleteMenuById(ctx context.Context, id int) error
	UpdateMenu(ctx context.Context, menu *model.Menu) (*model.Menu, error)

	FindAllPartners(ctx context.Context) ([]*model.Partner, error)
	FindPartnerById(ctx context.Context, id int) (*model.Partner, error)
	AddPartner(ctx context.Context, Name string, Email string, Password string) (string, error)
	UpdatePartnerById(ctx context.Context, partner *model.Partner) (*model.Partner, error)
	DeletePartnerById(ctx context.Context, id int) error

	FindAllAdmins(ctx context.Context) ([]*model.Admin, error)
	FindAdminById(ctx context.Context, id int) (*model.Admin, error)
	AddAdmin(ctx context.Context, Name string, Email string, Password string) (string, error)
	UpdateAdminById(ctx context.Context, admin *model.Admin) (*model.Admin, error)
	DeleteAdminById(ctx context.Context, id int) error

	FindAllTechSupports(ctx context.Context) ([]*model.TechSupport, error)
	FindTechSupportById(ctx context.Context, id int) (*model.TechSupport, error)
	AddTechSupport(ctx context.Context, Name string, Email string, Password string) (string, error)
	UpdateTechSupportById(ctx context.Context, techSupport *model.TechSupport) (*model.TechSupport, error)
	DeleteTechSupportById(ctx context.Context, id int) error

	FindAllCustomers(ctx context.Context) ([]*model.Customer, error)
	FindCustomerById(ctx context.Context, id int) (*model.Customer, error)
	AddCustomer(ctx context.Context, Name string, Email string, Password string, DeliveryAddress string, City string, Birthdate time.Time) (string, error)
	UpdateCustomerById(ctx context.Context, customer *model.Customer) (*model.Customer, error)
	DeleteCustomerById(ctx context.Context, id int) error

	FindAllReviews(ctx context.Context) ([]*model.Review, error)
	FindReviewById(ctx context.Context, id int) (*model.Review, error)
	AddReview(ctx context.Context, CustomerId int, RestaurantId int, MenuId int, Point int, Review string, Date time.Time) (string, error)
	UpdateReviewById(ctx context.Context, review *model.Review) (*model.Review, error)
	DeleteReviewById(ctx context.Context, id int) error

	FindAllSchedules(ctx context.Context) ([]*model.Schedule, error)
	FindScheduleById(ctx context.Context, id int) (*model.Schedule, error)
	AddSchedule(ctx context.Context, DayOfWeek string, OpeningTime time.Time, ClosingTime time.Time) (string, error)
	UpdateScheduleById(ctx context.Context, schedule *model.Schedule) (*model.Schedule, error)
	DeleteScheduleById(ctx context.Context, id int) error

	FindAllDeliveryPersonnels(ctx context.Context) ([]*model.DeliveryPersonnel, error)
	FindDeliveryPersonnelById(ctx context.Context, id int) (*model.DeliveryPersonnel, error)
	AddDeliveryPersonnel(ctx context.Context, Name string, Email string, Password string, AvailabilityStatus model.DeliveryPersonnelAvailability) (string, error)
	UpdateDeliveryPersonnelById(ctx context.Context, deliveryPersonnel *model.DeliveryPersonnel) (*model.DeliveryPersonnel, error)
	DeleteDeliveryPersonnelById(ctx context.Context, id int) error
}
