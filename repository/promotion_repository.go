package repository

import (
	"promotions/model"

	"gorm.io/gorm"
)

type PromotionRepository struct {
	db gorm.DB
}

// TODO: url config
func GetPromotionRepository(db gorm.DB) PromotionRepository {
	return PromotionRepository{
		db: db,
	}
}

func (repo PromotionRepository) GetById(id uint) model.Promotion {
	var promotion model.Promotion
	repo.db.Find(&promotion, id)
	return promotion
}

func (repo PromotionRepository) Save(promotion model.Promotion) {
	repo.db.Create(&promotion)
}
