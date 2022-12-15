package repository

import (
	"promotions/model"
	"time"

	"gorm.io/gorm"
)

type HistoryRepository struct {
	db gorm.DB
}

func GetHistoryRepository(db gorm.DB) HistoryRepository {
	return HistoryRepository{
		db: db,
	}
}

func (repo HistoryRepository) GetAfter(processedAfter time.Time) []model.ProcessedFile {
	var processedFiles []model.ProcessedFile
	repo.db.Where("processing_date > ?", processedAfter).Find(&processedFiles)
	return processedFiles
}

func (repo HistoryRepository) Save(processedFile model.ProcessedFile) {
	repo.db.Create(&processedFile)
}
