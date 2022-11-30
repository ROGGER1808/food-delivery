package foodbiz

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	foodmodel "gitlab.com/genson1808/food-delivery/module/food/model"
)

type FindFoodStore interface {
	FindByCondition(
		ctx context.Context,
		condition map[string]any,
		moreKeys ...string,
	) (*foodmodel.Food, error)
}

type findFoodBiz struct {
	store FindFoodStore
}

func NewFindFoodBiz(store FindFoodStore) *findFoodBiz {
	return &findFoodBiz{store: store}
}

func (biz *findFoodBiz) GetById(
	ctx context.Context,
	id int,
) (*foodmodel.Food, error) {
	// logic

	result, err := biz.store.FindByCondition(ctx, map[string]any{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(foodmodel.EntityName, err)
	}

	return result, nil
}
