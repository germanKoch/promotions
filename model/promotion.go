package model

import (
	"time"
)

type Promotion struct {
	Id             uint
	ExternalId     string
	Price          float32
	ExpirationDate time.Time
}
