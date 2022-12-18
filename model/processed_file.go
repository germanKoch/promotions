package model

import (
	"time"
)

type ProcessedFile struct {
	Id             uint `gorm:"primaryKey"`
	Path           string
	ProcessingDate time.Time
}

func (ProcessedFile) TableName() string {
	return "file_history"
}
