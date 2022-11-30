package restaurantlikestore

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	restaurantlikemodel "gitlab.com/genson1808/food-delivery/module/restaurantlike/model"
)

func (s *store) GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)

	var listLike []struct {
		RestaurantId int `gorm:"column:restaurant_id;"`
		LikeCount    int `gorm:"column:count;"`
	}

	if err := s.db.Table(restaurantlikemodel.Like{}.TableName()).
		Select("restaurant_id, count(restaurant_id) as count").
		Where("restaurant_id in (?)", ids).
		Group("restaurant_id").
		Find(&listLike).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listLike {
		result[item.RestaurantId] = item.LikeCount
	}

	return result, nil
}

func (s *store) GetUserLikedRestaurant(
	ctx context.Context,
	conditions map[string]any,
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]common.SimpleUser, error) {

	db := s.db
	db = db.Table(restaurantlikemodel.Like{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.RestaurantId > 0 {
			db.Where("restaurant_id = ?", v.RestaurantId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db.Preload("User")
	if v := paging.FakeCursor; v != "" {
		likeAt, _ := common.Base58ToUnixInt(paging.FakeCursor)
		db = db.Where("like_at < ?", likeAt)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	var listRestaurantLikeResult []restaurantlikemodel.Like

	if err := db.Limit(paging.Limit).
		Order("like_at desc").
		Find(&listRestaurantLikeResult).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	var result = make([]common.SimpleUser, len(listRestaurantLikeResult))
	for i, item := range listRestaurantLikeResult {
		result[i] = *item.User
		result[i].CreatedAt = nil
		result[i].UpdatedAt = nil
	}

	if len(listRestaurantLikeResult) > 0 {
		last := listRestaurantLikeResult[len(listRestaurantLikeResult)-1]
		paging.NextCursor = common.UnixToBase58(last.LikeAt)
	}

	return result, nil
}
