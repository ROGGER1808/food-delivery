package restaurantrepo

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	restaurantmodel "gitlab.com/genson1808/food-delivery/module/restaurant/model"
	"log"
)

type ListRestaurantStore interface {
	List(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type GetRestaurantLikeStore interface {
	GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantRepo struct {
	store     ListRestaurantStore
	likeStore GetRestaurantLikeStore
}

func NewListRestaurantRepo(store ListRestaurantStore, likeStore GetRestaurantLikeStore) *listRestaurantRepo {
	return &listRestaurantRepo{store: store, likeStore: likeStore}
}

func (repo *listRestaurantRepo) List(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	// logic

	result, err := repo.store.List(ctx, filter, paging, "User")
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	ids := make([]int, len(result))
	for i := range result {
		ids[i] = result[i].Id
	}

	likeMap, err := repo.likeStore.GetRestaurantLike(ctx, ids)
	if err != nil {
		log.Println(err)
		return result, nil
	}

	for i, item := range result {
		result[i].LikeCount = likeMap[item.Id]
	}

	return result, nil
}
