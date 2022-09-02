package restaurantlikebusiness

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	restaurantmodel "gitlab.com/genson1808/food-delivery/module/restaurant/model"
	restaurantlikemodel "gitlab.com/genson1808/food-delivery/module/restaurantlike/model"
)

type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, data *restaurantlikemodel.Like) error
}

type userDislikeRestaurantBiz struct {
	store           UserDislikeRestaurantStore
	restaurantStore RestaurantStore
}

func NewUserDislikeRestaurantBiz(store UserDislikeRestaurantStore, restaurantStore RestaurantStore) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{store: store, restaurantStore: restaurantStore}
}

func (biz *userDislikeRestaurantBiz) DislikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error {
	restaurantFound, err := biz.restaurantStore.FindByCondition(ctx, map[string]any{"id": data.RestaurantId})
	if err != nil {
		return common.ErrEntityNotFound(restaurantmodel.EntityName, err)
	}

	if restaurantFound.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	err = biz.store.Delete(ctx, data)
	if err != nil {
		return restaurantlikemodel.ErrCannotDislikeRestaurant(err)
	}

	return nil
}
