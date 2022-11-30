package foodLikeBiz

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	foodlikemodel "gitlab.com/genson1808/food-delivery/module/foodlike/model"
	usermodel "gitlab.com/genson1808/food-delivery/module/user/model"
)

type foodLikeStore interface {
	GetUsersLikedFood(
		ctx context.Context,
		filter *foodlikemodel.Filter,
		foodId int,
		paging *common.Paging,
	) ([]common.SimpleUser, error)

	GetFoodsLiked(
		ctx context.Context,
		ids []int,
		userId int,
	) (map[int]bool, error)
}

type likeFoodBiz struct {
	store foodLikeStore
}

func NewLikeFoodBiz(store foodLikeStore) *likeFoodBiz {
	return &likeFoodBiz{store: store}
}

func (biz *likeFoodBiz) GetUsersLikedFood(
	ctx context.Context,
	filter *foodlikemodel.Filter,
	foodId int,
	paging *common.Paging) ([]common.SimpleUser, error) {

	result, err := biz.store.GetUsersLikedFood(ctx, filter, foodId, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(usermodel.EntityName, err)
	}

	return result, err
}
