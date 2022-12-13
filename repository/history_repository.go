package repository

import (
	"promotions/model"

	"gorm.io/gorm"
)

type HistoryRepository struct {
	db gorm.DB
}

// TODO: url config
func GetHistoryRepository(db gorm.DB) PromotionRepository {
	return PromotionRepository{
		db: db,
	}
}

func (repo HistoryRepository) GetById(id uint) model.Promotion {
	var promotion model.Promotion
	repo.db.Find(&promotion, id)
	return promotion
}

func (repo HistoryRepository) Save(promotion model.Promotion) {
	repo.db.Create(&promotion)
}
