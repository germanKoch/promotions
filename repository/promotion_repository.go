package repository

import (
	"promotions/model"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PromotionRepository struct {
	db gorm.DB
}

func PromotionRepositoryPostgres() PromotionRepository {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=promotion port=5432"), &gorm.Config{})
	if err != nil {
		// control error
	}

	dbExec, err := db.DB()
	if err != nil {
		// control error
	}

	dbExec.SetMaxIdleConns(10)
	dbExec.SetMaxOpenConns(100)
	dbExec.SetConnMaxLifetime(time.Hour)
	return PromotionRepository{
		db: *db,
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
