package service

import (
	"DiplomaWork/internal/app/model"
	"DiplomaWork/internal/app/repository"
	"context"
	"time"
)

type DiplomaServiceImpl struct {
	dipRepository repository.DiplomaRepository
}

func NewDiplomaService(repo repository.DiplomaRepository) DiplomaService {
	return &DiplomaServiceImpl{dipRepository: repo}
}

func (as *DiplomaServiceImpl) GetAllRestaurant(ctx context.Context) ([]*model.RestaurantsModel, error) {
	rests, err := as.dipRepository.FindAllRestaurants(ctx)
	if err != nil {
		return nil, err
	}
	return rests, nil
}
func (as *DiplomaServiceImpl) GetRestaurantById(ctx context.Context, id int) (*model.RestaurantsModel, error) {
	restaurant, err := as.dipRepository.FindRestaurantById(ctx, id)
	if err != nil {
		return nil, err
	}
	return restaurant, nil
}
func (as *DiplomaServiceImpl) AddRestaurant(ctx context.Context, RestaurantName string, CategoryID int, PartnerId int, Address string, City string, PriceOfService int, RestaurantUIN int, PhoneNumber string, Rating float64, Schedule []model.Schedule) (string, error) {
	msg, err := as.dipRepository.AddRestaurants(ctx, RestaurantName, CategoryID, PartnerId, Address,
		City, PriceOfService, RestaurantUIN, PhoneNumber, Rating, Schedule)
	if err != nil {
		return msg, err
	}
	return msg, nil
}
func (as *DiplomaServiceImpl) UpdateRestaurant(ctx context.Context, restaurant *model.RestaurantsModel) error {
	_, err := as.dipRepository.UpdateRestaurant(ctx, restaurant)
	if err != nil {
		return err
	}
	return nil
}
func (as *DiplomaServiceImpl) DeleteRestaurantById(ctx context.Context, id int) error {
	err := as.dipRepository.DeleteRestaurantById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func (as *DiplomaServiceImpl) GetCategories(ctx context.Context) ([]*model.Category, error) {
	categories, err := as.dipRepository.AllCategories(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (as *DiplomaServiceImpl) GetCategoryById(ctx context.Context, id int) (*model.Category, error) {
	category, err := as.dipRepository.FindCategoryById(ctx, id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (as *DiplomaServiceImpl) AddCategory(ctx context.Context, Type string) (string, error) {
	msg, err := as.dipRepository.AddCategory(ctx, Type)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func (as *DiplomaServiceImpl) UpdateCategory(ctx context.Context, category *model.Category) error {
	_, err := as.dipRepository.UpdateCategory(ctx, category)
	if err != nil {
		return err
	}
	return nil
}
func (as *DiplomaServiceImpl) DeleteCategory(ctx context.Context, id int) error {
	err := as.dipRepository.DeleteCategory(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (as *DiplomaServiceImpl) GetAllMenu(ctx context.Context) ([]*model.Menu, error) {
	menus, err := as.dipRepository.FindAllMenu(ctx)
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (as *DiplomaServiceImpl) GetMenuById(ctx context.Context, id int) (*model.Menu, error) {
	menu, err := as.dipRepository.FindMenuById(ctx, id)
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (as *DiplomaServiceImpl) AddMenu(ctx context.Context, Name string, CategoryId int, RestaurantId int, Description string, Price int32) (string, error) {
	msg, err := as.dipRepository.AddMenu(ctx, Name, CategoryId, RestaurantId, Description, Price)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func (as *DiplomaServiceImpl) DeleteMenuById(ctx context.Context, id int) error {
	err := as.dipRepository.DeleteMenuById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (as *DiplomaServiceImpl) UpdateMenu(ctx context.Context, menu *model.Menu) error {
	_, err := as.dipRepository.UpdateMenu(ctx, menu)
	if err != nil {
		return err
	}
	return nil
}

func (as *DiplomaServiceImpl) GetPartners(ctx context.Context) ([]*model.Partner, error) {
	partners, err := as.dipRepository.FindAllPartners(ctx)
	if err != nil {
		return nil, err
	}
	return partners, nil
}

func (as *DiplomaServiceImpl) GetPartnerById(ctx context.Context, id int) (*model.Partner, error) {
	partner, err := as.dipRepository.FindPartnerById(ctx, id)
	if err != nil {
		return nil, err
	}
	return partner, nil
}

func (as *DiplomaServiceImpl) AddPartner(ctx context.Context, Name string, Email string, Password string) (string, error) {
	msg, err := as.dipRepository.AddPartner(ctx, Name, Email, Password)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func (as *DiplomaServiceImpl) UpdatePartnerById(ctx context.Context, partner *model.Partner) error {
	_, err := as.dipRepository.UpdatePartnerById(ctx, partner)
	if err != nil {
		return err
	}
	return nil
}
func (as *DiplomaServiceImpl) DeletePartnerById(ctx context.Context, id int) error {
	err := as.dipRepository.DeletePartnerById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (as *DiplomaServiceImpl) GetAdmins(ctx context.Context) ([]*model.Admin, error) {
	admins, err := as.dipRepository.FindAllAdmins(ctx)
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func (as *DiplomaServiceImpl) GetAdminById(ctx context.Context, id int) (*model.Admin, error) {
	admin, err := as.dipRepository.FindAdminById(ctx, id)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (as *DiplomaServiceImpl) AddAdmin(ctx context.Context, Name string, Email string, Password string) (string, error) {
	msg, err := as.dipRepository.AddAdmin(ctx, Name, Email, Password)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func (as *DiplomaServiceImpl) UpdateAdminById(ctx context.Context, admin *model.Admin) error {
	_, err := as.dipRepository.UpdateAdminById(ctx, admin)
	if err != nil {
		return err
	}
	return nil
}
func (as *DiplomaServiceImpl) DeleteAdminById(ctx context.Context, id int) error {
	err := as.dipRepository.DeleteAdminById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (as *DiplomaServiceImpl) GetTechSupports(ctx context.Context) ([]*model.TechSupport, error) {
	techSupports, err := as.dipRepository.FindAllTechSupports(ctx)
	if err != nil {
		return nil, err
	}
	return techSupports, nil
}

func (as *DiplomaServiceImpl) GetTechSupportById(ctx context.Context, id int) (*model.TechSupport, error) {
	techSupport, err := as.dipRepository.FindTechSupportById(ctx, id)
	if err != nil {
		return nil, err
	}
	return techSupport, nil
}

func (as *DiplomaServiceImpl) AddTechSupport(ctx context.Context, Name string, Email string, Password string) (string, error) {
	msg, err := as.dipRepository.AddTechSupport(ctx, Name, Email, Password)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func (as *DiplomaServiceImpl) UpdateTechSupportById(ctx context.Context, techSupport *model.TechSupport) error {
	_, err := as.dipRepository.UpdateTechSupportById(ctx, techSupport)
	if err != nil {
		return err
	}
	return nil
}
func (as *DiplomaServiceImpl) DeleteTechSupportById(ctx context.Context, id int) error {
	err := as.dipRepository.DeleteTechSupportById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (as *DiplomaServiceImpl) GetCustomers(ctx context.Context) ([]*model.Customer, error) {
	customers, err := as.dipRepository.FindAllCustomers(ctx)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (as *DiplomaServiceImpl) GetCustomerById(ctx context.Context, id int) (*model.Customer, error) {
	customer, err := as.dipRepository.FindCustomerById(ctx, id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (as *DiplomaServiceImpl) AddCustomer(ctx context.Context, Name string, Email string, Password string, DeliveryAddress string, City string, Birthdate time.Time) (string, error) {
	msg, err := as.dipRepository.AddCustomer(ctx, Name, Email, Password, DeliveryAddress, City, Birthdate)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func (as *DiplomaServiceImpl) UpdateCustomerById(ctx context.Context, customer *model.Customer) error {
	_, err := as.dipRepository.UpdateCustomerById(ctx, customer)
	if err != nil {
		return err
	}
	return nil
}
func (as *DiplomaServiceImpl) DeleteCustomerById(ctx context.Context, id int) error {
	err := as.dipRepository.DeleteCustomerById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (as *DiplomaServiceImpl) GetReviews(ctx context.Context) ([]*model.Review, error) {
	reviews, err := as.dipRepository.FindAllReviews(ctx)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (as *DiplomaServiceImpl) GetReviewById(ctx context.Context, id int) (*model.Review, error) {
	review, err := as.dipRepository.FindReviewById(ctx, id)
	if err != nil {
		return nil, err
	}
	return review, nil
}

func (as *DiplomaServiceImpl) AddReview(ctx context.Context, CustomerId int, RestaurantId int, MenuId int, Point int, Review string, Date time.Time) (string, error) {
	msg, err := as.dipRepository.AddReview(ctx, CustomerId, RestaurantId, MenuId, Point, Review, Date)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func (as *DiplomaServiceImpl) UpdateReviewById(ctx context.Context, review *model.Review) error {
	_, err := as.dipRepository.UpdateReviewById(ctx, review)
	if err != nil {
		return err
	}
	return nil
}
func (as *DiplomaServiceImpl) DeleteReviewById(ctx context.Context, id int) error {
	err := as.dipRepository.DeleteReviewById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (as *DiplomaServiceImpl) GetSchedules(ctx context.Context) ([]*model.Schedule, error) {
	schedules, err := as.dipRepository.FindAllSchedules(ctx)
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

func (as *DiplomaServiceImpl) GetScheduleById(ctx context.Context, id int) (*model.Schedule, error) {
	schedule, err := as.dipRepository.FindScheduleById(ctx, id)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

func (as *DiplomaServiceImpl) AddSchedule(ctx context.Context, DayOfWeek string, OpeningTime time.Time, ClosingTime time.Time) (string, error) {
	msg, err := as.dipRepository.AddSchedule(ctx, DayOfWeek, OpeningTime, ClosingTime)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func (as *DiplomaServiceImpl) UpdateScheduleById(ctx context.Context, schedule *model.Schedule) error {
	_, err := as.dipRepository.UpdateScheduleById(ctx, schedule)
	if err != nil {
		return err
	}
	return nil
}
func (as *DiplomaServiceImpl) DeleteScheduleById(ctx context.Context, id int) error {
	err := as.dipRepository.DeleteScheduleById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (as *DiplomaServiceImpl) GetDeliveryPersonnels(ctx context.Context) ([]*model.DeliveryPersonnel, error) {
	deliveryPersonnels, err := as.dipRepository.FindAllDeliveryPersonnels(ctx)
	if err != nil {
		return nil, err
	}
	return deliveryPersonnels, nil
}

func (as *DiplomaServiceImpl) GetDeliveryPersonnelById(ctx context.Context, id int) (*model.DeliveryPersonnel, error) {
	deliveryPersonnel, err := as.dipRepository.FindDeliveryPersonnelById(ctx, id)
	if err != nil {
		return nil, err
	}
	return deliveryPersonnel, nil
}

func (as *DiplomaServiceImpl) AddDeliveryPersonnel(ctx context.Context, Name string, Email string, Password string, AvailabilityStatus model.DeliveryPersonnelAvailability) (string, error) {
	msg, err := as.dipRepository.AddDeliveryPersonnel(ctx, Name, Email, Password, AvailabilityStatus)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func (as *DiplomaServiceImpl) UpdateDeliveryPersonnelById(ctx context.Context, deliveryPersonnel *model.DeliveryPersonnel) error {
	_, err := as.dipRepository.UpdateDeliveryPersonnelById(ctx, deliveryPersonnel)
	if err != nil {
		return err
	}
	return nil
}
func (as *DiplomaServiceImpl) DeleteDeliveryPersonnelById(ctx context.Context, id int) error {
	err := as.dipRepository.DeleteDeliveryPersonnelById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
