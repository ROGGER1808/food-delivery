package foodlikemodel

import (
	"gitlab.com/genson1808/food-delivery/common"
	foodmodel "gitlab.com/genson1808/food-delivery/module/food/model"
	"gorm.io/gorm"
	"time"
)

type FoodLike struct {
	FoodId    int                `json:"food_id" gorm:"column:food_id;"`
	Food      *foodmodel.Food    `json:"food" gorm:""`
	UserId    int                `json:"user_id" gorm:"column:user_id;"`
	User      *common.SimpleUser `json:"user" gorm:""`
	CreatedAt time.Time          `json:"created_at" gorm:"column:created_at;"`
	LikeAt    int64              `json:"like_at" gorm:"column:like_at;"`
}

func (m *FoodLike) BeforeCreate(tx *gorm.DB) (err error) {
	m.LikeAt = time.Now().UnixNano()
	return
}

func (FoodLike) TableName() string {
	return "food_likes"
}

func (f *FoodLike) GetFoodId() int {
	return f.FoodId
}

type FoodLikeCreate struct {
	FoodId int `json:"food_id" gorm:"column:food_id;"`
	UserId int `json:"user_id" gorm:"column:user_id;"`
}

func (FoodLikeCreate) TableName() string {
	return "food_likes"
}

func (f *FoodLikeCreate) GetFoodId() int {
	return f.FoodId
}

var (
	ErrUserAlreadyLikeFood = common.NewCustomError(nil, "user already like food", "UserAlreadyLikeFood")
	ErrUserNotLikeFoodYet  = common.NewCustomError(nil, "user not like food yet", "ErrUserNotLikeFoodYet")
)
