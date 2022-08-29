package restaurantmodel

import "time"

type Restaurant struct {
	Id        int       `json:"id" gorm:"column:id;"`
	Name      string    `json:"name" gorm:"column:name;"`
	Addr      string    `json:"addr" gorm:"column:addr;"`
	Lat       float32   `json:"lat" gorm:"column:lat;"`
	Lng       float32   `json:"lng" gorm:"column:Lng;"`
	Status    int       `json:"status" gorm:"column:status;"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}
