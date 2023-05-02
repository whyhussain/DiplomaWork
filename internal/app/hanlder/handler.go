package hanlder

import (
	"DiplomaWork/internal/app/model"
	"DiplomaWork/internal/app/service"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"

	"net/http"
)

type DipHandler struct {
	service service.DiplomaService
}

func NewDiplomaHandler(srv service.DiplomaService) *DipHandler {
	return &DipHandler{
		service: srv}

}

//  -----------------------------  RESTAURANT CRUD

func (h *DipHandler) GetAllRestaurants(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	restaurants, err := h.service.GetAllRestaurant(ctx)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(http.StatusOK, restaurants)

}

func (h *DipHandler) GetRestaurantById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	restaurantIdStr := c.Param("id")
	restaurantId, err := strconv.Atoi(restaurantIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid restaurant ID")
	}

	restaurant, err := h.service.GetRestaurantById(ctx, restaurantId)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(http.StatusOK, restaurant)
}

func (h *DipHandler) AddRestaurant(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	body := c.Request().Body
	req, _ := io.ReadAll(body)
	defer body.Close()

	rest := &model.RestaurantsModel{}
	err := json.Unmarshal(req, rest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if rest.RestaurantName == "" || rest.CategoryID == 0 {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameters").Error())
	}

	restaurants, err := h.service.AddRestaurant(ctx, rest.RestaurantName, rest.CategoryID, rest.PartnerId, rest.Address, rest.City, rest.PriceOfService, rest.RestaurantUIN, rest.PhoneNumber, rest.Rating, rest.Schedule)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, restaurants)

}
func (h *DipHandler) UpdateRestaurantById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Get menu ID from URL param
	restaurantIdStr := c.Param("id")
	restaurantId, err := strconv.Atoi(restaurantIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid restaurant ID")
	}
	log.Println("working ID")
	// Get the existing menu
	existingRestaurant, err := h.service.GetRestaurantById(ctx, restaurantId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to update restaurant")
	}
	log.Println("found existing id restaurant")

	// Parse request body to get updated fields
	var updateRestaurantFields = model.RestaurantsModel{Id: restaurantId}
	if err := c.Bind(&updateRestaurantFields); err != nil {
		log.Println("not binded update fields")
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	if updateRestaurantFields.RestaurantName == "" || updateRestaurantFields.CategoryID == 0 {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameters").Error())
	} else {
		existingRestaurant.RestaurantName = updateRestaurantFields.RestaurantName
		existingRestaurant.CategoryID = updateRestaurantFields.CategoryID
		existingRestaurant.PartnerId = updateRestaurantFields.PartnerId
		existingRestaurant.Address = updateRestaurantFields.Address
		existingRestaurant.City = updateRestaurantFields.City
		existingRestaurant.PriceOfService = updateRestaurantFields.PriceOfService
		existingRestaurant.RestaurantUIN = updateRestaurantFields.RestaurantUIN
		existingRestaurant.PhoneNumber = updateRestaurantFields.PhoneNumber
		existingRestaurant.Rating = updateRestaurantFields.Rating
		existingRestaurant.Schedule = updateRestaurantFields.Schedule
	}

	// Save the updated menu to the database
	if err := h.service.UpdateRestaurant(ctx, existingRestaurant); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to update restaurant")
	}

	return c.JSON(http.StatusOK, existingRestaurant)
}
func (h *DipHandler) DeleteRestaurantById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	restaurantIdStr := c.Param("id")
	restaurantId, err := strconv.Atoi(restaurantIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid restaurant ID")
	}

	err = h.service.DeleteRestaurantById(ctx, restaurantId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to delete restaurant")
	}

	return c.NoContent(http.StatusOK)
}

//  -----------------------------  CATEGORY CRUD

func (h *DipHandler) GetCategories(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	categories, err := h.service.GetCategories(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Println(json.Marshal(categories))

	return c.JSON(http.StatusOK, categories)

}
func (h *DipHandler) GetCategoryById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	categoryIdStr := c.Param("id")
	categoryId, err := strconv.Atoi(categoryIdStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid category ID")
	}

	category, err := h.service.GetCategoryById(ctx, categoryId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Error fetching category")
	}

	return c.JSON(http.StatusOK, category)
}

func (h *DipHandler) UpdateCategory(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	categoryIdStr := c.Param("id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid category ID")
	}
	existingCategory, err := h.service.GetCategoryById(ctx, categoryId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to update category")
	}

	var updateCategoryFields = model.Category{Id: categoryId}
	if err := c.Bind(&updateCategoryFields); err != nil {
		log.Println("not binded update fields")
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	if updateCategoryFields.Type == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameters").Error())
	} else {
		existingCategory.Type = updateCategoryFields.Type
	}
	if err := h.service.UpdateCategory(ctx, existingCategory); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to update category")
	}

	return c.JSON(http.StatusOK, existingCategory)
}

func (h *DipHandler) DeleteCategory(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	categoryIdStr := c.Param("id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid category ID")
	}

	err = h.service.DeleteCategory(ctx, categoryId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to delete category")
	}

	return c.NoContent(http.StatusOK)
}

func (h *DipHandler) AddCategory(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	body := c.Request().Body
	req, _ := io.ReadAll(body)
	defer body.Close()

	category := &model.Category{}
	err := json.Unmarshal(req, category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if category.Type == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameter").Error())
	}

	categories, err := h.service.AddCategory(ctx, category.Type)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(http.StatusCreated, categories)

}

//  -----------------------------  MENU CRUD

func (h *DipHandler) GetAllMenu(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	menus, err := h.service.GetAllMenu(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Println(json.Marshal(menus))

	return c.JSON(http.StatusOK, menus)

}

func (h *DipHandler) GetMenuById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	menuIdStr := c.Param("id")
	menuId, err := strconv.Atoi(menuIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid menu ID")
	}

	menu, err := h.service.GetMenuById(ctx, menuId)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(http.StatusOK, menu)
}

func (h *DipHandler) AddMenu(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	body := c.Request().Body
	req, _ := io.ReadAll(body)
	defer body.Close()

	menu := &model.Menu{}
	err := json.Unmarshal(req, menu)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if menu.Name == "" || menu.RestaurantId == 0 || menu.Price == 0 || menu.CategoryId == 0 {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameters").Error())
	}

	menus, err := h.service.AddMenu(ctx, menu.Name, menu.CategoryId, menu.RestaurantId, menu.Description, menu.Price)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, menus)

}

func (h *DipHandler) DeleteMenuById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	menuIdStr := c.Param("id")
	menuId, err := strconv.Atoi(menuIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid menu ID")
	}

	err = h.service.DeleteMenuById(ctx, menuId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to delete menu")
	}

	return c.NoContent(http.StatusOK)
}

func (h *DipHandler) UpdateMenuById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Get menu ID from URL param
	menuIdStr := c.Param("id")
	menuId, err := strconv.Atoi(menuIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid menu ID")
	}
	log.Println("working ID")
	// Get the existing menu
	existingMenu, err := h.service.GetMenuById(ctx, menuId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to update menu")
	}
	log.Println("found existing id menu")

	// Parse request body to get updated fields
	var updateMenuFields = model.Menu{Id: menuId}
	if err := c.Bind(&updateMenuFields); err != nil {
		log.Println("not binded update fields")
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	if updateMenuFields.Name == "" || updateMenuFields.RestaurantId == 0 || updateMenuFields.Price == 0 || updateMenuFields.CategoryId == 0 {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameters").Error())
	} else {
		existingMenu.Name = updateMenuFields.Name
		existingMenu.Price = updateMenuFields.Price
		existingMenu.RestaurantId = updateMenuFields.RestaurantId
		existingMenu.CategoryId = updateMenuFields.CategoryId
		existingMenu.Description = updateMenuFields.Description
	}

	// Save the updated menu to the database
	if err := h.service.UpdateMenu(ctx, existingMenu); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to update menu")
	}

	return c.JSON(http.StatusOK, existingMenu)
}

//  -----------------------------  PARTNER CRUD

func (h *DipHandler) GetPartners(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	partners, err := h.service.GetPartners(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Println(json.Marshal(partners))

	return c.JSON(http.StatusOK, partners)
}
func (h *DipHandler) GetPartnerById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	partnerIdStr := c.Param("id")
	partnerId, err := strconv.Atoi(partnerIdStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid partner ID")
	}

	partner, err := h.service.GetPartnerById(ctx, partnerId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Error fetching partner")
	}

	return c.JSON(http.StatusOK, partner)
}
func (h *DipHandler) UpdatePartnerById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	partnerIdStr := c.Param("id")
	partnerId, err := strconv.Atoi(partnerIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid partner ID")
	}
	existingPartner, err := h.service.GetPartnerById(ctx, partnerId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to update partner")
	}

	var updatePartnerFields = model.Partner{Id: partnerId}
	if err := c.Bind(&updatePartnerFields); err != nil {
		log.Println("not binded update fields")
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	if updatePartnerFields.Email == "" || updatePartnerFields.Name == "" || updatePartnerFields.Password == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameters").Error())
	} else {
		existingPartner.Name = updatePartnerFields.Name
		existingPartner.Email = updatePartnerFields.Email
		existingPartner.Password = updatePartnerFields.Password
	}
	if err := h.service.UpdatePartnerById(ctx, existingPartner); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to update partner")
	}

	return c.JSON(http.StatusOK, existingPartner)
}
func (h *DipHandler) AddPartner(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	body := c.Request().Body
	req, _ := io.ReadAll(body)
	defer body.Close()

	partner := &model.Partner{}
	err := json.Unmarshal(req, partner)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if partner.Email == "" || partner.Password == "" || partner.Name == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameter").Error())
	}

	partners, err := h.service.AddPartner(ctx, partner.Name, partner.Email, partner.Password)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(http.StatusCreated, partners)
}
func (h *DipHandler) DeletePartnerById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	partnerIdStr := c.Param("id")
	partnerId, err := strconv.Atoi(partnerIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid partner ID")
	}
	err = h.service.DeletePartnerById(ctx, partnerId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to delete partner")
	}
	return c.NoContent(http.StatusOK)
}

//  -----------------------------  ADMIN CRUD

func (h *DipHandler) GetAdmins(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	admins, err := h.service.GetAdmins(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Println(json.Marshal(admins))

	return c.JSON(http.StatusOK, admins)
}
func (h *DipHandler) GetAdminById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	adminIdStr := c.Param("id")
	adminId, err := strconv.Atoi(adminIdStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid admin ID")
	}

	admin, err := h.service.GetAdminById(ctx, adminId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Error fetching admin")
	}

	return c.JSON(http.StatusOK, admin)
}
func (h *DipHandler) UpdateAdminById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	adminIdStr := c.Param("id")
	adminId, err := strconv.Atoi(adminIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid admin ID")
	}
	existingAdmin, err := h.service.GetAdminById(ctx, adminId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to update admin")
	}

	var updateAdminFields = model.Admin{Id: adminId}
	if err := c.Bind(&updateAdminFields); err != nil {
		log.Println("not binded update fields")
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	if updateAdminFields.Email == "" || updateAdminFields.Name == "" || updateAdminFields.Password == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameters").Error())
	} else {
		existingAdmin.Name = updateAdminFields.Name
		existingAdmin.Email = updateAdminFields.Email
		existingAdmin.Password = updateAdminFields.Password
	}
	if err := h.service.UpdateAdminById(ctx, existingAdmin); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to update admin")
	}

	return c.JSON(http.StatusOK, existingAdmin)
}
func (h *DipHandler) AddAdmin(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	body := c.Request().Body
	req, _ := io.ReadAll(body)
	defer body.Close()

	admin := &model.Admin{}
	err := json.Unmarshal(req, admin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if admin.Email == "" || admin.Password == "" || admin.Name == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameter").Error())
	}

	admins, err := h.service.AddAdmin(ctx, admin.Name, admin.Email, admin.Password)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(http.StatusCreated, admins)
}
func (h *DipHandler) DeleteAdminById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	adminIdStr := c.Param("id")
	adminId, err := strconv.Atoi(adminIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid admin ID")
	}
	err = h.service.DeleteAdminById(ctx, adminId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to delete admin")
	}
	return c.NoContent(http.StatusOK)
}

// ----------------------- TECH SUPPORT CRUD
func (h *DipHandler) GetTechSupports(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	techSupports, err := h.service.GetTechSupports(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Println(json.Marshal(techSupports))

	return c.JSON(http.StatusOK, techSupports)
}
func (h *DipHandler) GetTechSupportById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	techSupportIdStr := c.Param("id")
	techSupportId, err := strconv.Atoi(techSupportIdStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid tech support ID")
	}

	techSupport, err := h.service.GetTechSupportById(ctx, techSupportId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Error fetching tech support")
	}

	return c.JSON(http.StatusOK, techSupport)
}
func (h *DipHandler) UpdateTechSupportById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	techSupportIdStr := c.Param("id")
	techSupportId, err := strconv.Atoi(techSupportIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid tech support ID")
	}
	existingTechSupport, err := h.service.GetTechSupportById(ctx, techSupportId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to update tech support")
	}

	var updateTechSupportFields = model.TechSupport{Id: techSupportId}
	if err := c.Bind(&updateTechSupportFields); err != nil {
		log.Println("not binded update fields")
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	if updateTechSupportFields.Email == "" || updateTechSupportFields.Name == "" || updateTechSupportFields.Password == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameters").Error())
	} else {
		existingTechSupport.Name = updateTechSupportFields.Name
		existingTechSupport.Email = updateTechSupportFields.Email
		existingTechSupport.Password = updateTechSupportFields.Password
	}
	if err := h.service.UpdateTechSupportById(ctx, existingTechSupport); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to update tech support")
	}

	return c.JSON(http.StatusOK, existingTechSupport)
}
func (h *DipHandler) AddTechSupport(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	body := c.Request().Body
	req, _ := io.ReadAll(body)
	defer body.Close()

	techSupport := &model.TechSupport{}
	err := json.Unmarshal(req, techSupport)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if techSupport.Email == "" || techSupport.Password == "" || techSupport.Name == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameter").Error())
	}

	techSupports, err := h.service.AddTechSupport(ctx, techSupport.Name, techSupport.Email, techSupport.Password)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(http.StatusCreated, techSupports)
}
func (h *DipHandler) DeleteTechSupportById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	techSupportIdStr := c.Param("id")
	techSupportId, err := strconv.Atoi(techSupportIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid tech support ID")
	}
	err = h.service.DeleteTechSupportById(ctx, techSupportId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to delete tech support")
	}
	return c.NoContent(http.StatusOK)
}

// ----------------------- CUSTOMER CRUD
func (h *DipHandler) GetCustomers(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	customers, err := h.service.GetCustomers(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Println(json.Marshal(customers))

	return c.JSON(http.StatusOK, customers)
}
func (h *DipHandler) GetCustomerById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	customerIdStr := c.Param("id")
	customerId, err := strconv.Atoi(customerIdStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid customer ID")
	}

	customer, err := h.service.GetCustomerById(ctx, customerId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Error fetching customer")
	}

	return c.JSON(http.StatusOK, customer)
}
func (h *DipHandler) UpdateCustomerById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	customerIdStr := c.Param("id")
	customerId, err := strconv.Atoi(customerIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid customer ID")
	}
	existingCustomer, err := h.service.GetCustomerById(ctx, customerId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to update customer")
	}

	var updateCustomerFields = model.Customer{Id: customerId}
	if err := c.Bind(&updateCustomerFields); err != nil {
		log.Println("not binded update fields")
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	if updateCustomerFields.Email == "" || updateCustomerFields.Name == "" || updateCustomerFields.Password == "" || updateCustomerFields.City == "" || updateCustomerFields.DeliveryAddress == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameters").Error())
	} else {
		existingCustomer.Name = updateCustomerFields.Name
		existingCustomer.Email = updateCustomerFields.Email
		existingCustomer.Password = updateCustomerFields.Password
		existingCustomer.DeliveryAddress = updateCustomerFields.DeliveryAddress
		existingCustomer.City = updateCustomerFields.City
		existingCustomer.Birthdate = updateCustomerFields.Birthdate
	}
	if err := h.service.UpdateCustomerById(ctx, existingCustomer); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to update customer")
	}

	return c.JSON(http.StatusOK, existingCustomer)
}
func (h *DipHandler) AddCustomer(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	body := c.Request().Body
	req, _ := io.ReadAll(body)
	defer body.Close()

	customer := &model.Customer{}
	err := json.Unmarshal(req, customer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if customer.Email == "" || customer.Password == "" || customer.Name == "" || customer.City == "" || customer.DeliveryAddress == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameter").Error())
	}

	customers, err := h.service.AddCustomer(ctx, customer.Name, customer.Email, customer.Password, customer.DeliveryAddress, customer.City, customer.Birthdate)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(http.StatusCreated, customers)
}
func (h *DipHandler) DeleteCustomerById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	customerIdStr := c.Param("id")
	customerId, err := strconv.Atoi(customerIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid customer ID")
	}
	err = h.service.DeleteCustomerById(ctx, customerId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to delete customer")
	}
	return c.NoContent(http.StatusOK)
}

// ------------------ Review
func (h *DipHandler) GetReviews(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reviews, err := h.service.GetReviews(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Println(json.Marshal(reviews))

	return c.JSON(http.StatusOK, reviews)
}
func (h *DipHandler) GetReviewById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reviewIdStr := c.Param("id")
	reviewId, err := strconv.Atoi(reviewIdStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid review ID")
	}

	review, err := h.service.GetReviewById(ctx, reviewId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Error fetching review")
	}

	return c.JSON(http.StatusOK, review)
}
func (h *DipHandler) UpdateReviewById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reviewIdStr := c.Param("id")
	reviewId, err := strconv.Atoi(reviewIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid review ID")
	}
	existingReview, err := h.service.GetReviewById(ctx, reviewId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to update review")
	}

	var updateReviewFields = model.Review{Id: reviewId}
	if err := c.Bind(&updateReviewFields); err != nil {
		log.Println("not binded update fields")
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	if updateReviewFields.Review == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameters").Error())
	} else {
		existingReview.CustomerId = updateReviewFields.CustomerId
		existingReview.RestaurantId = updateReviewFields.RestaurantId
		existingReview.MenuId = updateReviewFields.MenuId
		existingReview.Point = updateReviewFields.Point
		existingReview.Review = updateReviewFields.Review
		existingReview.Date = updateReviewFields.Date
	}
	if err := h.service.UpdateReviewById(ctx, existingReview); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to update review")
	}

	return c.JSON(http.StatusOK, existingReview)
}
func (h *DipHandler) AddReview(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	body := c.Request().Body
	req, _ := io.ReadAll(body)
	defer body.Close()

	review := &model.Review{}
	err := json.Unmarshal(req, review)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if review.Review == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameter").Error())
	}

	reviews, err := h.service.AddReview(ctx, review.CustomerId, review.RestaurantId, review.MenuId, review.Point, review.Review, review.Date)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(http.StatusCreated, reviews)
}
func (h *DipHandler) DeleteReviewById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	reviewIdStr := c.Param("id")
	reviewId, err := strconv.Atoi(reviewIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid review ID")
	}
	err = h.service.DeleteReviewById(ctx, reviewId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to delete review")
	}
	return c.NoContent(http.StatusOK)
}

func (h *DipHandler) GetSchedules(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	schedules, err := h.service.GetSchedules(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Println(json.Marshal(schedules))

	return c.JSON(http.StatusOK, schedules)
}
func (h *DipHandler) GetScheduleById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	scheduleIdStr := c.Param("id")
	scheduleId, err := strconv.Atoi(scheduleIdStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid schedule ID")
	}

	schedule, err := h.service.GetScheduleById(ctx, scheduleId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Error fetching schedule")
	}

	return c.JSON(http.StatusOK, schedule)
}
func (h *DipHandler) UpdateScheduleById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	scheduleIdStr := c.Param("id")
	scheduleId, err := strconv.Atoi(scheduleIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid schedule ID")
	}
	existingSchedule, err := h.service.GetScheduleById(ctx, scheduleId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to update schedule")
	}

	var updateScheduleFields = model.Schedule{Id: scheduleId}
	if err := c.Bind(&updateScheduleFields); err != nil {
		log.Println("not binded update fields")
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	if updateScheduleFields.DayOfWeek == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameters").Error())
	} else {
		existingSchedule.DayOfWeek = updateScheduleFields.DayOfWeek
		existingSchedule.OpeningTime = updateScheduleFields.OpeningTime
		existingSchedule.ClosingTime = updateScheduleFields.ClosingTime
	}
	if err := h.service.UpdateScheduleById(ctx, existingSchedule); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to update schedule")
	}

	return c.JSON(http.StatusOK, existingSchedule)
}
func (h *DipHandler) AddSchedule(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	body := c.Request().Body
	req, _ := io.ReadAll(body)
	defer body.Close()

	schedule := &model.Schedule{}
	err := json.Unmarshal(req, schedule)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if schedule.DayOfWeek == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameter").Error())
	}

	schedules, err := h.service.AddSchedule(ctx, schedule.DayOfWeek, schedule.OpeningTime, schedule.ClosingTime)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(http.StatusCreated, schedules)
}
func (h *DipHandler) DeleteScheduleById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	scheduleIdStr := c.Param("id")
	scheduleId, err := strconv.Atoi(scheduleIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid schedule ID")
	}
	err = h.service.DeleteScheduleById(ctx, scheduleId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to delete schedule")
	}
	return c.NoContent(http.StatusOK)
}

func (h *DipHandler) GetDeliveryPersonnels(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	deliveryPersonnels, err := h.service.GetDeliveryPersonnels(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Println(json.Marshal(deliveryPersonnels))

	return c.JSON(http.StatusOK, deliveryPersonnels)
}

func (h *DipHandler) GetDeliveryPersonnelById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	deliveryPersonnelIdStr := c.Param("id")
	deliveryPersonnelId, err := strconv.Atoi(deliveryPersonnelIdStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid delivery personnel ID")
	}

	deliveryPersonnel, err := h.service.GetDeliveryPersonnelById(ctx, deliveryPersonnelId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Error fetching delivery personnel")
	}

	return c.JSON(http.StatusOK, deliveryPersonnel)
}
func (h *DipHandler) UpdateDeliveryPersonnelById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	deliveryPersonnelIdStr := c.Param("id")
	deliveryPersonnelId, err := strconv.Atoi(deliveryPersonnelIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid delivery personnel ID")
	}
	existingDeliveryPersonnel, err := h.service.GetDeliveryPersonnelById(ctx, deliveryPersonnelId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to update delivery personnel")
	}

	var updateDeliveryPersonnelFields = model.DeliveryPersonnel{Id: deliveryPersonnelId}
	if err := c.Bind(&updateDeliveryPersonnelFields); err != nil {
		log.Println("not binded update fields")
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	if updateDeliveryPersonnelFields.Name == "" || updateDeliveryPersonnelFields.Email == "" || updateDeliveryPersonnelFields.Password == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameters").Error())
	} else {
		existingDeliveryPersonnel.Name = updateDeliveryPersonnelFields.Name
		existingDeliveryPersonnel.Email = updateDeliveryPersonnelFields.Email
		existingDeliveryPersonnel.Password = updateDeliveryPersonnelFields.Password
		existingDeliveryPersonnel.AvailabilityStatus = updateDeliveryPersonnelFields.AvailabilityStatus
	}
	if err := h.service.UpdateDeliveryPersonnelById(ctx, existingDeliveryPersonnel); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to update delivery personnel")
	}

	return c.JSON(http.StatusOK, existingDeliveryPersonnel)
}
func (h *DipHandler) AddDeliveryPersonnel(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	body := c.Request().Body
	req, _ := io.ReadAll(body)
	defer body.Close()

	deliveryPersonnel := &model.DeliveryPersonnel{}
	err := json.Unmarshal(req, deliveryPersonnel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if deliveryPersonnel.Name == "" || deliveryPersonnel.Email == "" || deliveryPersonnel.Password == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameter").Error())
	}

	//availabilityStatusNum, err := strconv.Atoi(c.FormValue("availability_status"))
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid availability_status").Error())
	//}
	//availabilityStatus := model.DeliveryPersonnelAvailability(availabilityStatusNum)

	deliveryPersonnels, err := h.service.AddDeliveryPersonnel(ctx, deliveryPersonnel.Name, deliveryPersonnel.Email, deliveryPersonnel.Password, deliveryPersonnel.AvailabilityStatus)

	//deliveryPersonnels, err := h.service.AddDeliveryPersonnel(ctx, deliveryPersonnel.Name, deliveryPersonnel.Email, deliveryPersonnel.Password, deliveryPersonnel.AvailabilityStatus)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(http.StatusCreated, deliveryPersonnels)
}
func (h *DipHandler) DeleteDeliveryPersonnelById(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	deliveryPersonnelIdStr := c.Param("id")
	deliveryPersonnelId, err := strconv.Atoi(deliveryPersonnelIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid delivery Personnel ID")
	}
	err = h.service.DeleteDeliveryPersonnelById(ctx, deliveryPersonnelId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to delete delivery personnel")
	}
	return c.NoContent(http.StatusOK)
}
