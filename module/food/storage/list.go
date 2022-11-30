package foodstorage

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	foodmodel "gitlab.com/genson1808/food-delivery/module/food/model"
)

func (s *store) List(
	ctx context.Context,
	filter *foodmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]foodmodel.Food, error) {
	var result []foodmodel.Food

	db := s.db.Table(foodmodel.Food{}.TableName()).Where("status in (1)")

	if f := filter; f != nil {
		if f.CategoryId > 0 {
			db = db.Where("category_id = ?", f.CategoryId)
		}
		if f.RestaurantId > 0 {
			db = db.Where("restaurant_id = ?", f.RestaurantId)
		}
		if f.OrderBy != "" {
			switch f.OrderBy {
			case "rating_desc":
				db = db.Order("rating desc")
			case "rating_asc":
				db = db.Order("rating_asc")
			case "like_desc":
				db = db.Order("like_desc")
			case "like_asc":
				db = db.Order("like_asc")
			}
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
