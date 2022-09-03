package restaurantlikebusiness

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	"gitlab.com/genson1808/food-delivery/component/asyncjob"
	restaurantmodel "gitlab.com/genson1808/food-delivery/module/restaurant/model"
	restaurantlikemodel "gitlab.com/genson1808/food-delivery/module/restaurantlike/model"
	"log"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type RestaurantLikeStore interface {
	FindByCondition(
		ctx context.Context,
		condition map[string]any,
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store           UserLikeRestaurantStore
	restaurantStore RestaurantLikeStore
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore, restaurantStore RestaurantLikeStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store, restaurantStore: restaurantStore}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error {
	restaurantFound, err := biz.restaurantStore.FindByCondition(ctx, map[string]any{"id": data.RestaurantId})
	if err != nil {
		return common.ErrEntityNotFound(restaurantmodel.EntityName, err)
	}

	if restaurantFound.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	err = biz.store.Create(ctx, data)
	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	// Side effect
	j := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.restaurantStore.IncreaseLikeCount(ctx, data.RestaurantId)
	})

	if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
		log.Println(err)
	}

	return nil
}
