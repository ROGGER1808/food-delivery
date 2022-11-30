package categorystorage

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	categorymodel "gitlab.com/genson1808/food-delivery/module/category/model"
)

func (s *store) Update(ctx context.Context, id int, data *categorymodel.CategoryUpdate) error {

	if err := s.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
