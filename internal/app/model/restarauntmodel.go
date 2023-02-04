package model

type RestarauntsModel struct {
	Id                 string `json:"id"`
	RestarauntName     string `json:"restaraunt_name"`
	RestarauntCategory int    `json:"restaraunt_category"`
}

type Menu struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}
