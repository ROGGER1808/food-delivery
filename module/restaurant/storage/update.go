package restaurantstorage

import (
	"context"
	restaurantmodel "gitlab.com/genson1808/food-delivery/module/restaurant/model"
)

func (s *store) Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {

	if err := s.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}
