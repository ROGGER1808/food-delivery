package foodratingstorage

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	foodratingmodel "gitlab.com/genson1808/food-delivery/module/foodrating/model"
	"gorm.io/gorm"
)

func (s *store) FindByCondition(
	ctx context.Context,
	condition map[string]any,
	moreKeys ...string,
) (*foodratingmodel.FoodRating, error) {
	var result foodratingmodel.FoodRating

	if err := s.db.Where(condition).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
