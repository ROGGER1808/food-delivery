package categorybiz

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	categorymodel "gitlab.com/genson1808/food-delivery/module/category/model"
)

type FindCategoryStore interface {
	FindByCondition(
		ctx context.Context,
		condition map[string]any,
		moreKeys ...string,
	) (*categorymodel.Category, error)
}

type findCategoryBiz struct {
	store FindCategoryStore
}

func NewFindCategoryBiz(store FindCategoryStore) *findCategoryBiz {
	return &findCategoryBiz{store: store}
}

func (biz *findCategoryBiz) GetById(
	ctx context.Context,
	id int,
) (*categorymodel.Category, error) {
	// logic

	result, err := biz.store.FindByCondition(ctx, map[string]any{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(categorymodel.EntityName, err)
	}

	return result, nil
}
