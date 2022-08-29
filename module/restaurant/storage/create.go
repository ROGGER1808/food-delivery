package restaurantstorage

import (
	"context"
	restaurantmodel "gitlab.com/genson1808/food-delivery/module/restaurant/model"
)

func (s *store) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
