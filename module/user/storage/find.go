package userstorage

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	usermodel "gitlab.com/genson1808/food-delivery/module/user/model"
	"gorm.io/gorm"
)

func (s *store) FindWithCondition(context context.Context, conditions map[string]any, moreKeys ...string) (*usermodel.User, error) {
	var data usermodel.User
	err := s.db.Where(conditions).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil

}

func (s *store) Find(context context.Context, conditions map[string]any, moreKeys ...string) (*usermodel.User, error) {
	return s.FindWithCondition(context, conditions, moreKeys...)
}
