package foodratingstorage

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	foodratingmodel "gitlab.com/genson1808/food-delivery/module/foodrating/model"
)

func (s *store) Create(ctx context.Context, data *foodratingmodel.FoodRatingCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
