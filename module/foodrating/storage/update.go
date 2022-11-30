package foodratingstorage

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	foodratingmodel "gitlab.com/genson1808/food-delivery/module/foodrating/model"
)

func (s *store) Update(ctx context.Context, id int, data *foodratingmodel.FoodRatingUpdate) error {

	if err := s.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
