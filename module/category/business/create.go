package categorybiz

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	categorymodel "gitlab.com/genson1808/food-delivery/module/category/model"
	restaurantmodel "gitlab.com/genson1808/food-delivery/module/restaurant/model"
)

type CreateCategoryStore interface {
	Create(ctx context.Context, data *categorymodel.CategoryCreate) error
}

type createCategoryBiz struct {
	store CreateCategoryStore
}

func NewCreateCategoryBiz(store CreateCategoryStore) *createCategoryBiz {
	return &createCategoryBiz{store: store}
}

func (biz *createCategoryBiz) Create(ctx context.Context, data *categorymodel.CategoryCreate) error {
	// logic

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
