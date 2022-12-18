package config

import (
	"promotions/model"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDb() gorm.DB {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=promotion port=5432"), &gorm.Config{})
	if err != nil {
		// control error
	}

	db.AutoMigrate(&model.Promotion{}, &model.ProcessedFile{})
	dbExec, err := db.DB()
	if err != nil {
		// control error
	}

	dbExec.SetMaxIdleConns(10)
	dbExec.SetMaxOpenConns(100)
	dbExec.SetConnMaxLifetime(time.Hour)

	return *db
}
