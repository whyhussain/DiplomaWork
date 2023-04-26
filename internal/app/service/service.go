package service

import (
	"DiplomaWork/internal/app/model"
	"context"
	"time"
)

type DiplomaService interface {
	GetAllRestaurant(ctx context.Context) ([]*model.RestaurantsModel, error)
	GetRestaurantById(ctx context.Context, id int) (*model.RestaurantsModel, error)
	AddRestaurant(ctx context.Context, RestaurantName string, CategoryID int, PartnerId int, Address string,
		City string, PriceOfService int, RestaurantUIN int, PhoneNumber string, Rating float64, Schedule []model.Schedule) (string, error)
	UpdateRestaurant(ctx context.Context, restaurant *model.RestaurantsModel) error
	DeleteRestaurantById(ctx context.Context, id int) error

	GetCategories(ctx context.Context) ([]*model.Category, error)
	GetCategoryById(ctx context.Context, id int) (*model.Category, error)
	AddCategory(ctx context.Context, Type string) (string, error)
	UpdateCategory(ctx context.Context, category *model.Category) error
	DeleteCategory(ctx context.Context, id int) error

	GetAllMenu(ctx context.Context) ([]*model.Menu, error)
	GetMenuById(ctx context.Context, id int) (*model.Menu, error)
	AddMenu(ctx context.Context, Name string, RestaurantId int, CategoryId int, Description string, Price int32) (string, error)
	DeleteMenuById(ctx context.Context, id int) error
	UpdateMenu(ctx context.Context, menu *model.Menu) error

	GetPartners(ctx context.Context) ([]*model.Partner, error)
	GetPartnerById(ctx context.Context, id int) (*model.Partner, error)
	AddPartner(ctx context.Context, Name string, Email string, Password string) (string, error)
	UpdatePartnerById(ctx context.Context, partner *model.Partner) error
	DeletePartnerById(ctx context.Context, id int) error

	GetAdmins(ctx context.Context) ([]*model.Admin, error)
	GetAdminById(ctx context.Context, id int) (*model.Admin, error)
	AddAdmin(ctx context.Context, Name string, Email string, Password string) (string, error)
	UpdateAdminById(ctx context.Context, admin *model.Admin) error
	DeleteAdminById(ctx context.Context, id int) error

	GetTechSupports(ctx context.Context) ([]*model.TechSupport, error)
	GetTechSupportById(ctx context.Context, id int) (*model.TechSupport, error)
	AddTechSupport(ctx context.Context, Name string, Email string, Password string) (string, error)
	UpdateTechSupportById(ctx context.Context, techSupport *model.TechSupport) error
	DeleteTechSupportById(ctx context.Context, id int) error

	GetCustomers(ctx context.Context) ([]*model.Customer, error)
	GetCustomerById(ctx context.Context, id int) (*model.Customer, error)
	AddCustomer(ctx context.Context, Name string, Email string, Password string, DeliveryAddress string, City string, Birthdate time.Time) (string, error)
	UpdateCustomerById(ctx context.Context, customer *model.Customer) error
	DeleteCustomerById(ctx context.Context, id int) error

	GetReviews(ctx context.Context) ([]*model.Review, error)
	GetReviewById(ctx context.Context, id int) (*model.Review, error)
	AddReview(ctx context.Context, CustomerId int, RestaurantId int, MenuId int, Point int, Review string, Date time.Time) (string, error)
	UpdateReviewById(ctx context.Context, review *model.Review) error
	DeleteReviewById(ctx context.Context, id int) error

	GetSchedules(ctx context.Context) ([]*model.Schedule, error)
	GetScheduleById(ctx context.Context, id int) (*model.Schedule, error)
	AddSchedule(ctx context.Context, DayOfWeek string, OpeningTime time.Time, ClosingTime time.Time) (string, error)
	UpdateScheduleById(ctx context.Context, schedule *model.Schedule) error
	DeleteScheduleById(ctx context.Context, id int) error

	GetDeliveryPersonnels(ctx context.Context) ([]*model.DeliveryPersonnel, error)
	GetDeliveryPersonnelById(ctx context.Context, id int) (*model.DeliveryPersonnel, error)
	AddDeliveryPersonnel(ctx context.Context, personnel *model.DeliveryPersonnel, AvailabilityStatus model.DeliveryPersonnelAvailability) (string, error)
	UpdateDeliveryPersonnelById(ctx context.Context, deliveryPersonnel *model.DeliveryPersonnel) error
	DeleteDeliveryPersonnelById(ctx context.Context, id int) error
}
