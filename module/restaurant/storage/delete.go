package restaurantstorage

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	restaurantmodel "gitlab.com/genson1808/food-delivery/module/restaurant/model"
)

func (s *store) Delete(ctx context.Context, id int) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]any{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
