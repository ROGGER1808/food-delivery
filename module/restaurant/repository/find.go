package restaurantrepo

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	restaurantmodel "gitlab.com/genson1808/food-delivery/module/restaurant/model"
)

type FindRestaurantStore interface {
	FindByCondition(
		ctx context.Context,
		condition map[string]any,
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type findRestaurantBiz struct {
	store FindRestaurantStore
}

func NewFindRestaurantBiz(store FindRestaurantStore) *findRestaurantBiz {
	return &findRestaurantBiz{store: store}
}

func (biz *findRestaurantBiz) GetById(
	ctx context.Context,
	id int,
) (*restaurantmodel.Restaurant, error) {
	// logic

	result, err := biz.store.FindByCondition(ctx, map[string]any{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	return result, nil
}
