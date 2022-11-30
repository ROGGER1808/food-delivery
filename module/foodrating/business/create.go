package foodbiz

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	foodmodel "gitlab.com/genson1808/food-delivery/module/food/model"
)

type CreateFoodStore interface {
	Create(ctx context.Context, data *foodmodel.FoodCreate) error
}

type createFoodBiz struct {
	store CreateFoodStore
}

func NewCreateFoodBiz(store CreateFoodStore) *createFoodBiz {
	return &createFoodBiz{store: store}
}

func (biz *createFoodBiz) Create(ctx context.Context, data *foodmodel.FoodCreate) error {
	// logic

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(foodmodel.EntityName, err)
	}

	return nil
}
