package foodbiz

import (
	"context"
	"github.com/labstack/gommon/log"
	"gitlab.com/genson1808/food-delivery/common"
	foodmodel "gitlab.com/genson1808/food-delivery/module/food/model"
)

type ListFoodStore interface {
	List(
		ctx context.Context,
		filter *foodmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]foodmodel.Food, error)
}

type listFoodBiz struct {
	foodStore     ListFoodStore
	foodLikeStore FindFoodLikedStore
}

func NewListFoodBiz(foodStore ListFoodStore, foodLikeStore FindFoodLikedStore) *listFoodBiz {
	return &listFoodBiz{foodStore: foodStore, foodLikeStore: foodLikeStore}
}

func (biz *listFoodBiz) ListFoodOfRestaurant(
	ctx context.Context,
	restaurantId int,
	userId int,
	filter *foodmodel.Filter,
	paging *common.Paging,
) ([]foodmodel.Food, error) {
	//TODO: check restaurant existed

	if filter == nil {
		filter = &foodmodel.Filter{}
	}
	filter.RestaurantId = restaurantId

	result, err := biz.foodStore.List(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(foodmodel.EntityName, err)
	}

	ids := make([]int, len(result))
	for i := range result {
		ids[i] = result[i].Id
	}

	likedMap, err := biz.foodLikeStore.GetFoodsLiked(ctx, ids, userId)
	if err != nil {
		log.Error(err)
	}

	if len(likedMap) > 0 {
		for i, item := range result {
			if v, exist := likedMap[item.Id]; exist {
				result[i].IsLike = v
			}
		}
	}

	return result, nil
}

func (biz *listFoodBiz) ListAllFood(
	ctx context.Context,
	userId int,
	paging *common.Paging,
	filter *foodmodel.Filter,
) ([]foodmodel.Food, error) {

	result, err := biz.foodStore.List(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(foodmodel.EntityName, err)
	}

	ids := make([]int, len(result))
	for i := range result {
		ids[i] = result[i].Id
	}

	likedMap, err := biz.foodLikeStore.GetFoodsLiked(ctx, ids, userId)
	if err != nil {
		log.Error(err)
	}

	if len(likedMap) > 0 {
		for i, item := range result {
			if v, exist := likedMap[item.Id]; exist {
				result[i].IsLike = v
			}
		}
	}

	return result, nil
}
