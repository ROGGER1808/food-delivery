package restaurantlikebusiness

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	restaurantlikemodel "gitlab.com/genson1808/food-delivery/module/restaurantlike/model"
)

type UserLikedRestaurantStore interface {
	GetUserLikedRestaurant(
		ctx context.Context,
		conditions map[string]any,
		filter *restaurantlikemodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]common.SimpleUser, error)
}

type restaurantLikedBiz struct {
	store UserLikedRestaurantStore
}

func NewRestaurantLikedBiz(store UserLikedRestaurantStore) *restaurantLikedBiz {
	return &restaurantLikedBiz{store: store}
}

func (biz *restaurantLikedBiz) GetUserLikedRestaurant(
	ctx context.Context,
	filter *restaurantlikemodel.Filter,
	paging *common.Paging) ([]common.SimpleUser, error) {

	result, err := biz.store.GetUserLikedRestaurant(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantlikemodel.EntityName, err)
	}

	return result, nil
}
