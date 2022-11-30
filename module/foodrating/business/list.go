package foodbiz

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	foodmodel "gitlab.com/genson1808/food-delivery/module/food/model"
)

type ListFoodRepo interface {
	List(
		ctx context.Context,
		filter *foodmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]foodmodel.Food, error)
}

type listFoodBiz struct {
	repo ListFoodRepo
}

func NewListFoodBiz(repo ListFoodRepo) *listFoodBiz {
	return &listFoodBiz{repo: repo}
}

func (biz *listFoodBiz) List(
	ctx context.Context,
	filter *foodmodel.Filter,
	paging *common.Paging,
) ([]foodmodel.Food, error) {
	// logic

	result, err := biz.repo.List(ctx, filter, paging, "Category")
	if err != nil {
		return nil, common.ErrCannotListEntity(foodmodel.EntityName, err)
	}

	return result, nil
}
