package model

type Restaurant struct {
	RestaurantName     string `json:"restaurant_name"`
	RestaurantCategory string `json:"restaurant_category"`
	RestaurantMenu     Menu   `json:"restaurant_menu"`
}

type Menu struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}
