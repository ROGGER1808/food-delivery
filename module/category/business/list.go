package categorybiz

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	categorymodel "gitlab.com/genson1808/food-delivery/module/category/model"
	restaurantmodel "gitlab.com/genson1808/food-delivery/module/restaurant/model"
)

type ListCategoryRepo interface {
	List(
		ctx context.Context,
		filter *categorymodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]categorymodel.Category, error)
}

type listCategoryBiz struct {
	repo ListCategoryRepo
}

func NewListCategoryBiz(repo ListCategoryRepo) *listCategoryBiz {
	return &listCategoryBiz{repo: repo}
}

func (biz *listCategoryBiz) List(
	ctx context.Context,
	filter *categorymodel.Filter,
	paging *common.Paging,
) ([]categorymodel.Category, error) {
	// logic

	result, err := biz.repo.List(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	return result, nil
}
