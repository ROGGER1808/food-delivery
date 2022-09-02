package restaurantlikebusiness

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	restaurantmodel "gitlab.com/genson1808/food-delivery/module/restaurant/model"
	restaurantlikemodel "gitlab.com/genson1808/food-delivery/module/restaurantlike/model"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type RestaurantStore interface {
	FindByCondition(
		ctx context.Context,
		condition map[string]any,
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type userLikeRestaurantBiz struct {
	store           UserLikeRestaurantStore
	restaurantStore RestaurantStore
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore, restaurantStore RestaurantStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store, restaurantStore: restaurantStore}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error {
	restaurantFound, err := biz.restaurantStore.FindByCondition(ctx, map[string]any{"id": data.RestaurantId})
	if err != nil {
		return common.ErrEntityNotFound(restaurantmodel.EntityName, err)
	}

	if restaurantFound.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	err = biz.store.Create(ctx, data)
	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	return nil
}
