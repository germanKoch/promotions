package repository

import (
	"promotions/model"
	"promotions/model/bussiness_error"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PromotionRepository struct {
	db gorm.DB
}

func GetPromotionRepository(db gorm.DB) PromotionRepository {
	return PromotionRepository{
		db: db,
	}
}

func (repo PromotionRepository) GetById(id string) (model.Promotion, error) {
	var promotion model.Promotion
	err := repo.db.First(&promotion, "id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return promotion, bussiness_error.ErrNotFound
	}
	return promotion, nil
}

func (repo PromotionRepository) UpsertAll(promotions []model.Promotion) {
	repo.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"price", "expiration_date"}),
	}).Create(&promotions)
}
