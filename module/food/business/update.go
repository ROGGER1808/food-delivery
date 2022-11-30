package foodbiz

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	foodmodel "gitlab.com/genson1808/food-delivery/module/food/model"
)

type UpdateFoodStore interface {
	Update(ctx context.Context, id int, data *foodmodel.FoodUpdate) error
}

type updateFoodBiz struct {
	store UpdateFoodStore
}

func NewUpdateFoodBiz(store UpdateFoodStore) *updateFoodBiz {
	return &updateFoodBiz{store: store}
}

func (biz *updateFoodBiz) Update(ctx context.Context, id int, data *foodmodel.FoodUpdate) error {
	// logic

	if err := biz.store.Update(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(foodmodel.EntityName, err)
	}
	return nil
}
