package model

import (
	"time"

	"gorm.io/gorm"
)

type Promotion struct {
	gorm.Model
	Id             uint
	ExternalId     string
	Price          float32
	ExpirationDate time.Time
}
