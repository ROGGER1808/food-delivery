package foodlikestorage

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	foodlikemodel "gitlab.com/genson1808/food-delivery/module/foodlike/model"
)

func (s *store) Delete(ctx context.Context, data *foodlikemodel.FoodLike) error {
	if err := s.db.Table(foodlikemodel.FoodLike{}.TableName()).
		Where("user_id = ? and food_id = ?", data.UserId, data.FoodId).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
