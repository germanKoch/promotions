package parser

import (
	"promotions/model"
	"strconv"
	"strings"
	"time"
)

type PromotionParser struct{}

func GetPromotionParser() PromotionParser {
	return PromotionParser{}
}

func (PromotionParser) Parse(csvString string) model.Promotion {
	fields := strings.Split(csvString, ",")
	//TODO: error handling
	//TODO: timezone
	price, _ := strconv.ParseFloat(fields[1], 64)
	expirationDate, _ := time.Parse("2006-01-02 15:04:05 +0200 CEST", fields[2])

	return model.Promotion{
		Id:             fields[0],
		Price:          price,
		ExpirationDate: expirationDate,
	}
}
