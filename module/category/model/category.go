package categorymodel

import (
	"gitlab.com/genson1808/food-delivery/common"
	"gitlab.com/genson1808/food-delivery/component/fimage"
)

const EntityName = "Category"

type Category struct {
	common.SQLModel
	Name        string        `json:"name" gorm:"column:name;"`
	Description string        `json:"description" gorm:"column:description;"`
	Icon        *fimage.Image `json:"icon" gorm:"column:icon;"`
}

func (Category) TableName() string {
	return "categories"
}

func (c *Category) Mask(isAdminOrOwner bool) {
	c.GenUID(common.DbTypeRestaurant)
}

type CategoryCreate struct {
	common.SQLModel
	Name        string `json:"name" gorm:"column:name;"`
	Description string `json:"description" gorm:"column:description;"`
	Icon        string `json:"icon" gorm:"column:icon;"`
}

func (CategoryCreate) TableName() string {
	return Category{}.TableName()
}

func (c *CategoryCreate) Mask(isAdminOrOwner bool) {
	c.GenUID(common.DbTypeRestaurant)
}

type CategoryUpdate struct {
	Name        *string `json:"name" gorm:"column:name;"`
	Description *string `json:"description" gorm:"column:description;"`
	Icon        *string `json:"icon" gorm:"column:icon;"`
	Status      *int    `json:"status" gorm:"column:status;"`
}

func (CategoryUpdate) TableName() string {
	return Category{}.TableName()
}
