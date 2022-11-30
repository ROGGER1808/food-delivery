package foodratingstorage

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	foodratingmodel "gitlab.com/genson1808/food-delivery/module/foodrating/model"
)

func (s *store) Delete(ctx context.Context, id int) error {
	if err := s.db.Table(foodratingmodel.FoodRating{}.TableName()).
		Where("id = ?", id).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
