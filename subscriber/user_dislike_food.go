package subscriber

import (
	"context"
	"gitlab.com/genson1808/food-delivery/component/appctx"
	"gitlab.com/genson1808/food-delivery/component/pubsub"
	foodstorage "gitlab.com/genson1808/food-delivery/module/food/storage"
)

func UserDislikeFood(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "user like food",
		Handler: func(ctx context.Context, msg *pubsub.Message) error {
			store := foodstorage.NewStore(appCtx.GetMainDBConnection())
			data := msg.Data().(HasFoodId)
			return store.DecreaseLikeCount(ctx, data.GetFoodId())
		},
	}
}
