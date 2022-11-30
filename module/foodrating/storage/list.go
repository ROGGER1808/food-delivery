package foodratingstorage

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	foodratingmodel "gitlab.com/genson1808/food-delivery/module/foodrating/model"
)

func (s *store) List(
	ctx context.Context,
	filter *foodratingmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]foodratingmodel.FoodRating, error) {
	var result []foodratingmodel.FoodRating

	db := s.db.Table(foodratingmodel.FoodRating{}.TableName()).Where("status in (1)")

	if f := filter; f != nil {
		if f.FoodId != 0 {
			db = db.Where("food_id = ?", f.FoodId)
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

func (s *store) CalculateAVGPoint(ctx context.Context, conditions map[string]any) (float64, error) {
	db := s.db.Table(foodratingmodel.FoodRating{}.TableName()).Where(conditions)

	var result float64

	if err := db.Select("AVG(point) result").Find(&result).Error; err != nil {
		return 0, common.ErrDB(err)
	}

	return result, nil
}
