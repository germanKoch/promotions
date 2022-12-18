package model

import (
	"time"
)

type Promotion struct {
	Id             string `gorm:"primaryKey"`
	Price          float64
	ExpirationDate time.Time
}

func (Promotion) TableName() string {
	return "promotion"
}
