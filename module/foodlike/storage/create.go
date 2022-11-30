package foodlikestorage

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	foodlikemodel "gitlab.com/genson1808/food-delivery/module/foodlike/model"
)

func (s *store) Create(ctx context.Context, data *foodlikemodel.FoodLikeCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
