package categorystorage

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	categorymodel "gitlab.com/genson1808/food-delivery/module/category/model"
)

func (s *store) List(
	ctx context.Context,
	filter *categorymodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]categorymodel.Category, error) {
	var result []categorymodel.Category

	db := s.db.Table(categorymodel.Category{}.TableName()).Where("status in (1)")

	if f := filter; f != nil {
		if len(f.Name) > 0 {
			db = db.Where("name LIKE ?", f.Name)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)
		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(result) > 0 {
		last := result[len(result)-1]
		last.Mask(true)
		paging.NextCursor = last.FakeId.String()
	}

	return result, nil
}
