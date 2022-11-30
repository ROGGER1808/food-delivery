package foodLikeBiz

import (
	"context"
	"github.com/labstack/gommon/log"
	"gitlab.com/genson1808/food-delivery/common"
	"gitlab.com/genson1808/food-delivery/component/pubsub"
	foodlikemodel "gitlab.com/genson1808/food-delivery/module/foodlike/model"
)

type deleteFoodLikeStore interface {
	FindByCondition(
		ctx context.Context,
		conditions map[string]any,
		moreKeys ...string,
	) (*foodlikemodel.FoodLike, error)

	Delete(ctx context.Context, data *foodlikemodel.FoodLike) error
}

type UserDislikeFoodBiz struct {
	store  deleteFoodLikeStore
	pubsub pubsub.PubSub
}

func NewUserDislikeFoodBiz(store deleteFoodLikeStore, pubsub pubsub.PubSub) *UserDislikeFoodBiz {
	return &UserDislikeFoodBiz{store: store, pubsub: pubsub}
}

func (biz *UserDislikeFoodBiz) UserDislikeFood(
	ctx context.Context,
	data *foodlikemodel.FoodLike,
) error {
	exist, err := biz.store.FindByCondition(ctx, map[string]any{"user_id": data.UserId, "food_id": data.FoodId})
	if err != nil {
		return common.ErrInvalidRequest(err)
	}

	if exist == nil {
		return foodlikemodel.ErrUserNotLikeFoodYet
	}

	err = biz.store.Delete(ctx, data)
	if err != nil {
		return common.ErrCannotCreateEntity("FoodLike", err)
	}

	err = biz.pubsub.Publish(ctx, common.TopicUserDislikeFood, pubsub.NewMessage(data))
	if err != nil {
		log.Error(err)
	}

	return nil
}
