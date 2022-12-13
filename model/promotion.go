package model

import (
	"time"
)

type Promotion struct {
	Id             uint
	ExternalId     string
	Price          float64
	ExpirationDate time.Time
}

func (Promotion) TableName() string {
	return "promotion"
}
