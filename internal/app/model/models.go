package model

import "time"

type RestaurantsModel struct {
	Id             int    `json:"id"`
	RestaurantName string `json:"restaurant_name"`
	//RestaurantCategory string `json:"restaurant_category"`
	CategoryID     int        `json:"category_id"`
	PartnerId      int        `json:"partner_id"`
	Address        string     `json:"address"`
	City           string     `json:"city"`
	PriceOfService int        `json:"price_of_service"`
	RestaurantUIN  int        `json:"restaurant_uin"`
	PhoneNumber    string     `json:"phone_number"`
	Rating         float64    `json:"rating"`
	Schedule       []Schedule `json:"schedule"`
	//Photo
}

type Schedule struct {
	Id          int                 `json:"id"`
	DayOfWeek   string              `json:"day_of_week"`
	OpeningTime time.Time           `json:"opening_time"`
	ClosingTime time.Time           `json:"closing_time"`
	Restaurants []*RestaurantsModel `json:"restaurants"`
}

type Category struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
}

type Menu struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	CategoryId int    `json:"category_id"`
	//Photo        uuid.UUID `json:"photo"`
	RestaurantId int    `json:"restaurant_id"`
	Description  string `json:"description"`
	Price        int32  `json:"price"`
}

// type OrderStatus int
// const (
//     Placed OrderStatus = iota
//     Accepted
//     Shipped
//     Canceled
// 	Completed
// )

// type Order struct {
// 	Id   int    `json:"id"`
// 	RestaurantId int `json:"restaurant_id"`
// 	CustomerId int `json:"customer_id"`
// 	DeliveryPersonnelId int `json:"delivery_personnel_id"`
// 	MenuId int `json:"menu_id"`
// 	DeliveryAddress string `json:"delivery_address"`
// 	DeliveryStatusId int `json:"delivery_status_id"`
// 	TotalPrice int `json:"total_price"`
// }

// type Role int
// const (
//     Admin Role = iota
//     DeliveryPersonnel
//     Customer
//     Partner
// 	TechSupport
// )

type Partner struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Customer struct {
	Id              int       `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	DeliveryAddress string    `json:"delivery_address"`
	City            string    `json:"city"`
	Birthdate       time.Time `json:"birthdate"`
}

// type DeliveryStatus struct {
// 	Id   int    `json:"id"`
// 	OrderId int `json:"order_id"`
// 	DeliveryPersonnelId int `json:"delivery_personnel_id"`
// 	OrderStatus OrderStatus `json:"delivery_status"`
// 	TimeOfDelivery int `json:"time_of_delivery"`
// }

type DeliveryPersonnelAvailability int

const (
	Available DeliveryPersonnelAvailability = iota
	Busy
	Offline
)

type DeliveryPersonnel struct {
	Id                 int                           `json:"id"`
	Name               string                        `json:"name"`
	Email              string                        `json:"email"`
	Password           string                        `json:"password"`
	AvailabilityStatus DeliveryPersonnelAvailability `json:"availability_status"`
}

type Review struct {
	Id           int       `json:"id"`
	CustomerId   int       `json:"customer_id"`
	RestaurantId int       `json:"restaurant_id"`
	MenuId       int       `json:"menu_id"`
	Point        int       `json:"point"`
	Review       string    `json:"review"`
	Date         time.Time `json:"date"`
}

type Admin struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TechSupport struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
