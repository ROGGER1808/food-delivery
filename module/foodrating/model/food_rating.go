package foodratingmodel

import (
	"gitlab.com/genson1808/food-delivery/common"
	foodmodel "gitlab.com/genson1808/food-delivery/module/food/model"
)

const EntityName = "FoodRating"

type FoodRating struct {
	common.SQLModel
	UserId  int                `json:"user_id" gorm:"column:user_id;"`
	User    *common.SimpleUser `json:"user" gorm:""`
	FoodId  int                `json:"food_id" gorm:"column:food_id;"`
	Food    *foodmodel.Food    `json:"food" gorm:""`
	Point   float64            `json:"point" gorm:"column:point;"`
	Comment string             `json:"comment" gorm:"column:comment;"`
}

func (FoodRating) TableName() string {
	return "food_ratings"
}

func (f *FoodRating) Mask(isAdmin bool) {
	f.GenUID(common.DbTypeFoodRating)
}

type FoodRatingUpdate struct {
	UserId  int     `json:"user_id" gorm:"column:user_id;"`
	FoodId  int     `json:"food_id" gorm:"column:food_id;"`
	Point   float64 `json:"point" gorm:"column:point;"`
	Comment string  `json:"comment" gorm:"column:comment;"`
	Status  int     `json:"status" gorm:"column:status;"`
}

func (FoodRatingUpdate) TableName() string {
	return FoodRating{}.TableName()
}

type FoodRatingCreate struct {
	UserId  int     `json:"user_id" gorm:"column:user_id;"`
	FoodId  int     `json:"food_id" gorm:"column:food_id;"`
	Point   float64 `json:"point" gorm:"column:point;"`
	Comment string  `json:"comment" gorm:"column:comment;"`
}

func (FoodRatingCreate) TableName() string {
	return FoodRating{}.TableName()
}
