package userstorage

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	usermodel "gitlab.com/genson1808/food-delivery/module/user/model"
)

func (s *store) Delete(context context.Context, id int) error {
	if err := s.db.Table(usermodel.User{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]int{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
