package foodbiz

import (
	"context"
	"github.com/labstack/gommon/log"
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

type FindFoodLikedStore interface {
	GetFoodsLiked(
		ctx context.Context,
		ids []int,
		userId int,
	) (map[int]bool, error)
}

type findFoodBiz struct {
	foodStore     FindFoodStore
	foodLikeStore FindFoodLikedStore
}

func NewFindFoodBiz(foodStore FindFoodStore, foodLikeStore FindFoodLikedStore) *findFoodBiz {
	return &findFoodBiz{foodStore: foodStore, foodLikeStore: foodLikeStore}
}

func (biz *findFoodBiz) GetById(
	ctx context.Context,
	foodId int,
	userId int,
) (*foodmodel.Food, error) {
	result, err := biz.foodStore.FindByCondition(ctx, map[string]any{"id": foodId}, "Category")
	if err != nil {
		return nil, common.ErrCannotGetEntity(foodmodel.EntityName, err)
	}

	LikedMap, err := biz.foodLikeStore.GetFoodsLiked(ctx, []int{result.Id}, userId)
	if err != nil {
		log.Error(err)
	}

	result.IsLike = LikedMap[result.Id]

	return result, nil
}
