package restaurantmodel

import (
	"gitlab.com/genson1808/food-delivery/common"
	"gitlab.com/genson1808/food-delivery/component/fimage"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string             `json:"name" gorm:"column:name;"`
	Addr            string             `json:"addr" gorm:"column:addr;"`
	Lat             float32            `json:"lat" gorm:"column:lat;"`
	Lng             float32            `json:"lng" gorm:"column:Lng;"`
	Logo            *fimage.Image      `json:"logo" gorm:"column:logo;"`
	Cover           *fimage.Images     `json:"cover" gorm:"column:cover;"`
	OwnerId         int                `json:"-" gorm:"column:owner_id;"`
	User            *common.SimpleUser `json:"user" gorm:"foreignKey:OwnerId;"`
	LikeCount       int                `json:"like_count" gorm:"column:like_count;"`
}

func (Restaurant) TableName() string { return "restaurants" }

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
	if r.User != nil {
		r.User.Mask(false)
	}
}

type RestaurantUpdate struct {
	Name  *string        `json:"name" gorm:"column:name;"`
	Addr  *string        `json:"addr" gorm:"column:addr;"`
	Logo  *fimage.Image  `json:"logo" gorm:"column:logo;"`
	Cover *fimage.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            string         `json:"addr" gorm:"column:addr;"`
	OwnerId         int            `json:"-" gorm:"column:owner_id;"`
	Logo            *fimage.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *fimage.Images `json:"cover" gorm:"column:cover;"`
}

func (r *RestaurantCreate) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}
