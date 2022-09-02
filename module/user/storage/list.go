package userstorage

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	usermodel "gitlab.com/genson1808/food-delivery/module/user/model"
)

func (s *store) List(
	context context.Context,
	paging *common.Paging,
	filter *usermodel.Filter,
	moreKeys ...string) ([]usermodel.User, error) {
	var result []usermodel.User

	db := s.db.Table(usermodel.User{}.TableName()).Where("status in (1)")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
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
		last.Mask(false)
		paging.NextCursor = last.FakeId.String()
	}
	return result, nil
}
