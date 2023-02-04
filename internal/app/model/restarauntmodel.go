package model

type RestarauntsModel struct {
	//Id                 string `json:"id"`
	RestarauntName     string `json:"restaraunt_name"`
	RestarauntCategory string `json:"restaraunt_category"`
}

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Menu struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}
