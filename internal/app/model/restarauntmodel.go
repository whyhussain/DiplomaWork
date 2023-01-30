package model

type RestarauntsModel struct {
	RestarauntName     string `json:"restaraunt_name"`
	RestarauntCategory string `json:"restaraunt_category"`
	RestarauntMenu     Menu   `json:"restaraunt_menu"`
}

type Menu struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}
