package userstorage

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	usermodel "gitlab.com/genson1808/food-delivery/module/user/model"
)

func (s *store) Create(context context.Context, data *usermodel.UserCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
