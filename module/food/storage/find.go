package foodstorage

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	foodmodel "gitlab.com/genson1808/food-delivery/module/food/model"
	"gorm.io/gorm"
)

func (s *store) FindByCondition(
	ctx context.Context,
	conditions map[string]any,
	moreKeys ...string,
) (*foodmodel.Food, error) {
	db := s.db

	db = db.Where(conditions)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var result foodmodel.Food

	if err := db.First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
