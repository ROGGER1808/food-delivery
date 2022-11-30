package categorystorage

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	categorymodel "gitlab.com/genson1808/food-delivery/module/category/model"
	"gorm.io/gorm"
)

func (s *store) FindByCondition(
	ctx context.Context,
	condition map[string]any,
	moreKeys ...string,
) (*categorymodel.Category, error) {
	db := s.db

	db = db.Where(condition)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var result categorymodel.Category

	if err := db.First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
