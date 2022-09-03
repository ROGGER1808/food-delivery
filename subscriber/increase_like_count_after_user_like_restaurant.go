package subscriber

import (
	"context"
	"gitlab.com/genson1808/food-delivery/component/appctx"
	"gitlab.com/genson1808/food-delivery/component/pubsub"
	restaurantstorage "gitlab.com/genson1808/food-delivery/module/restaurant/storage"
)

/*
func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, common.TopicUserLikeRestaurant)

	store := restaurantstorage.NewStore(appCtx.GetMainDBConnection())

	go func() {
		defer common.AppRecover()
		for {
			msg := <-c
			likeData := msg.Data().(HasRestaurantId)
			_ = store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		}
	}()
}
*/

//IncreaseLikeAfterUserLikeRestaurant increase like restaurant using  pubsub integrate with asyncJob
func IncreaseLikeAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "User liked restaurant",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
