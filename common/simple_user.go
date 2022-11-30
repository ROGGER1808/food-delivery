package common

import "gitlab.com/genson1808/food-delivery/component/fimage"

type SimpleUser struct {
	SQLModel
	LastName  string        `json:"last_name" gorm:"column:last_name;"`
	FirstName string        `json:"first_name" gorm:"column:first_name;"`
	Role      string        `json:"role" gorm:"column:role;"`
	Avatar    *fimage.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (SimpleUser) TableName() string {
	return "users"
}

func (u *SimpleUser) Mask(isAdmin bool) {
	u.GenUID(DbTypeUser)
}
