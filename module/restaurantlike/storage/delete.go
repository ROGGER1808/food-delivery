package restaurantlikestore

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	restaurantlikemodel "gitlab.com/genson1808/food-delivery/module/restaurantlike/model"
)

func (s *store) Delete(ctx context.Context, data *restaurantlikemodel.Like) error {
	if err := s.db.Table(restaurantlikemodel.Like{}.TableName()).
		Where("restaurant_id = ? and user_id = ?", data.RestaurantId, data.UserId).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
