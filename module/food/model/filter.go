package foodmodel

type Filter struct {
	CategoryId   int    `json:"category_id" form:"category_id"`
	RestaurantId int    `json:"restaurant_id" gorm:"restaurant_id"`
	Name         string `json:"name" form:"name"`
	OrderBy      string `json:"order_by" form:"order_by"`
}
