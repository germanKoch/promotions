package database

import (
	"fmt"
	"math"
	"promotions/config"
	"promotions/model"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDb(dbConfig config.DbConfig) gorm.DB {
	dbConfigStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBname,
		dbConfig.Port,
	)
	db, err := gorm.Open(postgres.Open(dbConfigStr), &gorm.Config{})
	if err != nil {
		panic("Could not construct db obj")
	}

	db.AutoMigrate(&model.Promotion{}, &model.ProcessedFile{})
	dbExec, err := db.DB()
	if err != nil {
		panic("Could not construct db obj")
	}

	dbExec.SetMaxIdleConns(dbConfig.ConnectionPool.MaxIdleConnections)
	dbExec.SetMaxOpenConns(dbConfig.ConnectionPool.MaxOpenConnections)
	dbExec.SetConnMaxLifetime(time.Duration(dbConfig.ConnectionPool.ConnectionLifetime * int64(math.Pow(10, 6))))
	return *db
}
