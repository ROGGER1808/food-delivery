package foodlikestorage

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	foodlikemodel "gitlab.com/genson1808/food-delivery/module/foodlike/model"
	"gorm.io/gorm"
)

func (s *store) FindByCondition(
	ctx context.Context,
	conditions map[string]any,
	moreKeys ...string,
) (*foodlikemodel.FoodLike, error) {
	db := s.db

	db = db.Where(conditions)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var result foodlikemodel.FoodLike

	if err := db.First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
