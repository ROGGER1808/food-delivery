package foodlikemodel

type Filter struct {
	FoodId int `json:"food_id" form:"food_id"`
	UserId int `json:"user_id" form:"user_id"`
}
