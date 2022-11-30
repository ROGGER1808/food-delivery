package categorybiz

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	categorymodel "gitlab.com/genson1808/food-delivery/module/category/model"
)

type DeleteCategoryStore interface {
	FindByCondition(
		ctx context.Context,
		condition map[string]any,
		moreKeys ...string,
	) (*categorymodel.Category, error)
	Delete(ctx context.Context, id int) error
}

type deleteCategoryBiz struct {
	store     DeleteCategoryStore
	requester common.Requester
}

func NewDeleteCategoryBiz(store DeleteCategoryStore) *deleteCategoryBiz {
	return &deleteCategoryBiz{store: store}
}

func (biz *deleteCategoryBiz) Delete(ctx context.Context, id int) error {
	oldData, err := biz.store.FindByCondition(ctx, map[string]any{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(categorymodel.EntityName, err)
	}
	if oldData.Status == 0 {
		return common.ErrEntityDeleted(categorymodel.EntityName, nil)
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(categorymodel.EntityName, err)
	}

	return nil
}
