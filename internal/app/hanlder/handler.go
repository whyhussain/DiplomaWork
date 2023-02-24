package hanlder

import (
	"DiplomaWork/internal/app/service"
<<<<<<< Updated upstream
	"github.com/labstack/echo/v4"
=======
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"

	"net/http"
>>>>>>> Stashed changes
)

type DipHandler struct {
	service service.DiplomaService
}

func NewDiplomaHandler(srv service.DiplomaService) *DipHandler {
	return &DipHandler{
		service: srv}

}

func (h *DipHandler) GetAllRestaurants(c echo.Context) error {
<<<<<<< Updated upstream

	return nil
=======
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

	restaurants, err := h.service.AddRestaurant(ctx, rest.RestaurantName, rest.CategoryID)
	if err != nil {
		log.Println(err)
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
	// TODO CHECK FOR VALUES EXISTING IN UPDATEMENUFIELDS
	if updateRestaurantFields.RestaurantName == "" || updateRestaurantFields.CategoryID == 0 {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameters").Error())
	} else {
		existingRestaurant.RestaurantName = updateRestaurantFields.RestaurantName
		existingRestaurant.CategoryID = updateRestaurantFields.CategoryID
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

func (h *DipHandler) GetCategories(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	categories, err := h.service.GetCategories(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Println(json.Marshal(categories))

	return c.JSON(http.StatusOK, categories)
>>>>>>> Stashed changes

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

	if menu.Name == "" || menu.RestaurantId == 0 || menu.Price == 0 {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameters").Error())
	}

	menus, err := h.service.AddMenu(ctx, menu.Name, menu.RestaurantId, menu.Price)
	if err != nil {
		log.Println(err)
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
	// TODO CHECK FOR VALUES EXISTING IN UPDATEMENUFIELDS
	if updateMenuFields.Name == "" || updateMenuFields.RestaurantId == 0 || updateMenuFields.Price == 0 {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameters").Error())
	} else {
		existingMenu.Name = updateMenuFields.Name
		existingMenu.Price = updateMenuFields.Price
		existingMenu.RestaurantId = updateMenuFields.RestaurantId
	}

	// Save the updated menu to the database
	if err := h.service.UpdateMenu(ctx, existingMenu); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to update menu")
	}

	return c.JSON(http.StatusOK, existingMenu)
}
