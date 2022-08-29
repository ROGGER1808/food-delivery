package restaurantbiz

import (
	"context"
	restaurantmodel "gitlab.com/genson1808/food-delivery/module/restaurant/model"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type CreateRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *CreateRestaurantBiz {
	return &CreateRestaurantBiz{store: store}
}

func (biz *CreateRestaurantBiz) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	// logic

	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
