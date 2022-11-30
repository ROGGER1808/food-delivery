package categorystorage

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	categorymodel "gitlab.com/genson1808/food-delivery/module/category/model"
)

func (s *store) Delete(ctx context.Context, id int) error {
	if err := s.db.Table(categorymodel.Category{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]any{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
