package foodlikestorage

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	foodlikemodel "gitlab.com/genson1808/food-delivery/module/foodlike/model"
)

func (s *store) GetUsersLikedFood(
	ctx context.Context,
	filter *foodlikemodel.Filter,
	foodId int,
	paging *common.Paging,
) ([]common.SimpleUser, error) {

	db := s.db.Where("food_id = ?", foodId)

	if f := filter; f != nil {
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	db.Preload("User")
	if v := paging.FakeCursor; v != "" {
		likeAt, _ := common.Base58ToUnixInt(paging.FakeCursor)
		db = db.Where("like_at < ?", likeAt)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	var listFoodLike []foodlikemodel.FoodLike

	if err := db.
		Limit(paging.Limit).
		Order("like_at desc").
		Find(&listFoodLike).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	var result = make([]common.SimpleUser, len(listFoodLike))
	for i, item := range listFoodLike {
		result[i] = *item.User
		result[i].CreatedAt = nil
		result[i].UpdatedAt = nil
	}

	if len(listFoodLike) > 0 {
		last := listFoodLike[len(listFoodLike)-1]
		paging.NextCursor = common.UnixToBase58(last.LikeAt)
	}

	return result, nil
}

func (s *store) GetFoodsLiked(
	ctx context.Context,
	ids []int,
	userId int,
) (map[int]bool, error) {

	result := make(map[int]bool)

	type data struct {
		FoodId int `gorm:"column:food_id;"`
	}

	var ListLiked []data

	if err := s.db.Table(foodlikemodel.FoodLike{}.TableName()).
		Select("food_id").
		Where(map[string]any{"food_id": ids, "user_id": userId}).
		Find(&ListLiked).Error; err != nil {
		return nil, err
	}

	for _, item := range ListLiked {
		result[item.FoodId] = true
	}

	return result, nil
}
