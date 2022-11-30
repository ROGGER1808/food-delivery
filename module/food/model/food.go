package foodmodel

import (
	"gitlab.com/genson1808/food-delivery/common"
	"gitlab.com/genson1808/food-delivery/component/fimage"
)

const EntityName = "Food"

type Food struct {
	common.SQLModel
	RestaurantId int                    `json:"restaurant_id" gorm:"column:restaurant_id;"`
	CategoryId   int                    `json:"category_id" gorm:"column:category_id;"`
	Category     *common.SimpleCategory `json:"category" gorm:"foreignKey:CategoryId;"`
	Name         string                 `json:"name" gorm:"column:name;"`
	Description  string                 `json:"description" gorm:"column:description;"`
	Price        float64                `json:"price" gorm:"column:price;"`
	Images       fimage.Images          `json:"images" gorm:"column:images;"`
	LikeCount    int                    `json:"like_count" gorm:"column:like_count;"`
	Rating       float64                `json:"rating" gorm:"column:rating;"`
	RatingCount  int                    `json:"rating_count" gorm:"column:rating_count;"`
	IsLike       bool                   `json:"is_like" gorm:"column:-"`
}

func (Food) TableName() string {
	return "foods"
}

func (f *Food) Mask(isAdminOrOwner bool) {
	f.GenUID(common.DbTypeFood)
	if f.Category != nil {
		f.Category.GenUID(common.DbTypeCategory)
	}
}

type FoodUpdate struct {
	CategoryId  *int     `json:"category_id" gorm:"column:category_id;"`
	Name        *string  `json:"name" gorm:"column:name;"`
	Description *string  `json:"description" gorm:"column:description;"`
	Price       *float64 `json:"price" gorm:"column:price;"`
	Status      *int     `json:"status" gorm:"column:status;"`
}

func (FoodUpdate) TableName() string {
	return Food{}.TableName()
}

type FoodCreate struct {
	common.SQLModel `json:",inline"`
	RestaurantId    int           `json:"restaurant_id" gorm:"column:restaurant_id;"`
	CategoryId      int           `json:"category_id" gorm:"column:category_id;"`
	Name            string        `json:"name" gorm:"column:name;"`
	Description     string        `json:"description" gorm:"column:description;"`
	Price           float64       `json:"price" gorm:"column:price;"`
	Images          fimage.Images `json:"images" gorm:"column:images;"`
}

func (FoodCreate) TableName() string {
	return Food{}.TableName()
}

func (f *FoodCreate) Mask(isAdminOrOwner bool) {
	f.GenUID(common.DbTypeFood)
}
