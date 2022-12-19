package parser

import (
	"promotions/model"
	"promotions/model/bussiness_error"
	"strconv"
	"strings"
	"time"
)

type PromotionParser struct{}

func GetPromotionParser() PromotionParser {
	return PromotionParser{}
}

func (PromotionParser) Parse(csvString string) (model.Promotion, error) {
	fields := strings.Split(csvString, ",")

	if len(fields) != 3 {
		return model.Promotion{}, bussiness_error.ErrInvalidLineFormat
	}

	price, err := strconv.ParseFloat(fields[1], 64)

	if err != nil {
		return model.Promotion{}, err
	}

	expirationDate, err := time.Parse("2006-01-02 15:04:05 -0700 MST", fields[2])

	if err != nil {
		return model.Promotion{}, err
	}

	return model.Promotion{
		Id:             fields[0],
		Price:          price,
		ExpirationDate: expirationDate,
	}, nil
}
