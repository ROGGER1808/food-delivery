package restaurantlikemodel

import (
	"gitlab.com/genson1808/food-delivery/common"
	"time"
)

type Like struct {
	RestaurantId int       `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int       `json:"user_id" gorm:"column:user_id;"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at;"`
}

func (Like) TableName() string {
	return "restaurant_likes"
}

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"Cannot like this restaurant",
		"ErrCannotLikeRestaurant",
	)
}

func ErrCannotDislikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"Cannot dislike this restaurant",
		"ErrCannotDislikeRestaurant",
	)
}
