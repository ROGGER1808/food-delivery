package categorystorage

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	categorymodel "gitlab.com/genson1808/food-delivery/module/category/model"
)

func (s *store) Create(ctx context.Context, data *categorymodel.CategoryCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
