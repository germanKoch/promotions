package parser

import (
	"promotions/model"
	"strconv"
	"strings"
	"time"
)

func PromotionParser(csvString string) model.Promotion {
	fields := strings.Split(csvString, ",")
	//TODO: error handling
	//TODO: timezone
	price, _ := strconv.ParseFloat(fields[1], 64)
	expirationDate, _ := time.Parse("2006-01-02 15:04:05 +0200 CEST", fields[2])

	return model.Promotion{
		Id:             0,
		ExternalId:     fields[0],
		Price:          price,
		ExpirationDate: expirationDate,
	}
}
