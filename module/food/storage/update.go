package foodstorage

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	foodmodel "gitlab.com/genson1808/food-delivery/module/food/model"
	"gorm.io/gorm"
)

func (s *store) Update(ctx context.Context, id int, data *foodmodel.FoodUpdate) error {

	if err := s.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *store) IncreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(foodmodel.Food{}.TableName()).
		Where("id = ?", id).
		Update("like_count", gorm.Expr("like_count + ?", 1)).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *store) DecreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(foodmodel.Food{}.TableName()).
		Where("id = ?", id).
		Update("like_count", gorm.Expr("like_count - ?", 1)).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *store) UpdateRating(ctx context.Context, id int, rating float64) error {
	db := s.db

	if err := db.Table(foodmodel.Food{}.TableName()).
		Where("id = ?", id).
		Update("rating", rating).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *store) IncreaseRatingCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(foodmodel.Food{}.TableName()).
		Where("id = ?", id).
		Update("rating_count", gorm.Expr("rating_count + ?", 1)).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
