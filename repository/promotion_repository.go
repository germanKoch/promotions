package repository

import (
	"promotions/model"

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

func (repo PromotionRepository) GetById(id string) model.Promotion {
	var promotion model.Promotion
	repo.db.Find(&promotion, id)
	return promotion
}

func (repo PromotionRepository) UpsertAll(promotions []model.Promotion) {
	repo.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"price", "expiration_date"}),
	}).Create(&promotions)
}
