package restaurantlikemodel

import (
	"gitlab.com/genson1808/food-delivery/common"
	"gorm.io/gorm"
	"time"
)

const EntityName = "RestaurantLike"

type Like struct {
	RestaurantId int `json:"restaurant_id" gorm:"column:restaurant_id;"`
	//CreatedAt    time.Time          `json:"created_at" gorm:"column:created_at;"`
	UserId int                `json:"user_id" gorm:"column:user_id;"`
	User   *common.SimpleUser `json:"user" gorm:""`
	LikeAt int64              `json:"like_at" gorm:"column:like_at;"`
}

func (m *Like) GetRestaurantId() int {
	return m.RestaurantId
}

func (m *Like) BeforeCreate(tx *gorm.DB) (err error) {
	m.LikeAt = time.Now().UnixNano()
	return
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
