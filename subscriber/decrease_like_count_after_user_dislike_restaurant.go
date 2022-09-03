package subscriber

import (
	"context"
	"gitlab.com/genson1808/food-delivery/component/appctx"
	"gitlab.com/genson1808/food-delivery/component/pubsub"
	restaurantstorage "gitlab.com/genson1808/food-delivery/module/restaurant/storage"
)

/*
//DecreaseLikeAfterUserLikeRestaurant decrease like restaurant only using  pubsub
func DecreaseLikeAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, common.TopicUserDislikeRestaurant)

	store := restaurantstorage.NewStore(appCtx.GetMainDBConnection())

	go func() {
		defer common.AppRecover()
		for {
			msg := <-c
			likeData := msg.Data().(HasRestaurantId)
			_ = store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		}
	}()
}
*/

//DecreaseLikeAfterUserLikeRestaurant decrease like restaurant using  pubsub integrate with asyncJob
func DecreaseLikeAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "User liked restaurant",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
