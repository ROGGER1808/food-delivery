package restaurantbiz

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	restaurantmodel "gitlab.com/genson1808/food-delivery/module/restaurant/model"
)

type UpdateRestaurantStore interface {
	Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	// logic

	if err := biz.store.Update(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
