package main

import (
	"fmt"
	"log"
	"promotions/api"
	"promotions/config"
	"promotions/database"
	"promotions/repository"
	"promotions/service"
	"promotions/service/parser"
	"promotions/storage"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfgPath := "./config.yml"
	cfg, configErr := config.NewConfig(cfgPath)
	if configErr != nil {
		log.Fatal(configErr)
		panic("Config could not be parsed")
	}

	app := fiber.New()
	db := database.GetDb(cfg.DbConfig)

	storage := storage.GetLocalStorage(cfg.LocalStorageConfig)
	repo := repository.GetPromotionRepository(db)
	promotionService := service.GetPromotionRepoService(repo)
	historyRepository := repository.GetHistoryRepository(db)
	promotionParser := parser.GetPromotionParser()

	scheduler := service.GetScheduledReader(cfg.SchedulerConfig, promotionService, historyRepository, promotionParser, storage)
	promotionController := api.GetPromotionController(promotionService)

	promotionController.GetRouts(app)
	scheduler.ScheduleJob()

	serverErr := app.Listen(fmt.Sprintf("%s:%s", cfg.ServerConfig.Host, cfg.ServerConfig.Port))
	if serverErr != nil {
		log.Fatal(serverErr)
		panic("Could not lanch server")
	}

}
