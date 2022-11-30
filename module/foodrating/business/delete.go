package foodbiz

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	foodmodel "gitlab.com/genson1808/food-delivery/module/food/model"
)

type DeleteFoodStore interface {
	FindByCondition(
		ctx context.Context,
		condition map[string]any,
		moreKeys ...string,
	) (*foodmodel.Food, error)
	Delete(ctx context.Context, id int) error
}

type deleteFoodBiz struct {
	store DeleteFoodStore
}

func NewDeleteFoodBiz(store DeleteFoodStore) *deleteFoodBiz {
	return &deleteFoodBiz{store: store}
}

func (biz *deleteFoodBiz) Delete(ctx context.Context, id int) error {
	oldData, err := biz.store.FindByCondition(ctx, map[string]any{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(foodmodel.EntityName, err)
	}
	if oldData.Status == 0 {
		return common.ErrEntityDeleted(foodmodel.EntityName, nil)
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(foodmodel.EntityName, err)
	}

	return nil
}
