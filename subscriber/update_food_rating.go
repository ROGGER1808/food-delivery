package subscriber

import (
	"context"
	"gitlab.com/genson1808/food-delivery/component/appctx"
	"gitlab.com/genson1808/food-delivery/component/pubsub"
	foodstorage "gitlab.com/genson1808/food-delivery/module/food/storage"
	foodratingstorage "gitlab.com/genson1808/food-delivery/module/foodrating/storage"
)

func UpdateCalculateRatingFood(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Upadte Calculate rating food",
		Handler: func(ctx context.Context, msg *pubsub.Message) error {
			db := appCtx.GetMainDBConnection()

			foodRatingStore := foodratingstorage.NewStore(db)
			foodStore := foodstorage.NewStore(db)

			data := msg.Data().(HasFoodId)

			rating, err := foodRatingStore.CalculateAVGPoint(ctx, map[string]any{"food_id": data.GetFoodId()})
			if err != nil {
				appCtx.Logger().Errorw("Subscriber.CalculateRatingFood.CalculateAVGPoint", "ERROR", err)
			}

			err = foodStore.UpdateRating(ctx, data.GetFoodId(), rating)
			if err != nil {
				appCtx.Logger().Errorw("Subscriber.CalculateRatingFood.UpdateRating", "ERROR", err)
			}

			return nil
		},
	}
}
