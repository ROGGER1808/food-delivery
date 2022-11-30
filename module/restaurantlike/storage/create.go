package restaurantlikestore

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	restaurantlikemodel "gitlab.com/genson1808/food-delivery/module/restaurantlike/model"
)

func (s *store) Create(ctx context.Context, data *restaurantlikemodel.Like) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
