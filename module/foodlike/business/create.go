package foodLikeBiz

import (
	"context"
	"github.com/labstack/gommon/log"
	"gitlab.com/genson1808/food-delivery/common"
	"gitlab.com/genson1808/food-delivery/component/pubsub"
	foodlikemodel "gitlab.com/genson1808/food-delivery/module/foodlike/model"
)

type createFoodLikeStore interface {
	FindByCondition(
		ctx context.Context,
		conditions map[string]any,
		moreKeys ...string,
	) (*foodlikemodel.FoodLike, error)

	Create(ctx context.Context, data *foodlikemodel.FoodLikeCreate) error
}

type UserLikeFoodBiz struct {
	store  createFoodLikeStore
	pubsub pubsub.PubSub
}

func NewUserLikeFoodBiz(store createFoodLikeStore, pubsub pubsub.PubSub) *UserLikeFoodBiz {
	return &UserLikeFoodBiz{store: store, pubsub: pubsub}
}

func (biz *UserLikeFoodBiz) UserLikeFood(
	ctx context.Context,
	data *foodlikemodel.FoodLikeCreate,
) error {
	exist, err := biz.store.FindByCondition(ctx, map[string]any{"user_id": data.UserId, "food_id": data.FoodId})
	if err != nil {
		return common.ErrInvalidRequest(err)
	}

	if exist != nil {
		return foodlikemodel.ErrUserAlreadyLikeFood
	}

	err = biz.store.Create(ctx, data)
	if err != nil {
		return common.ErrCannotCreateEntity("FoodLike", err)
	}

	err = biz.pubsub.Publish(ctx, common.TopicUserLikeFood, pubsub.NewMessage(data))
	if err != nil {
		log.Error(err)
	}

	return nil
}
