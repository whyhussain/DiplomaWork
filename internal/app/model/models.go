package model

type RestaurantsModel struct {
	Id                 int    `json:"id"`
	RestaurantName     string `json:"restaurant_name"`
	RestaurantCategory string `json:"restaurant_category"`
	CategoryID         int    `json:"category_id"`
}

type Category struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
}

type Menu struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	//Photo        uuid.UUID `json:"photo"`
	RestaurantId int `json:"restaurant_id"`
	Price        int `json:"price"`
}
