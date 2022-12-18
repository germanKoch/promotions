package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"promotions/api"
	"promotions/config"
	"promotions/database"
	"promotions/repository"
	"promotions/service"
	"promotions/service/parser"
	"promotions/service/storage"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfgPath, configErr := parseFlags()
	if configErr != nil {
		log.Fatal(configErr)
		panic("Config params undefined")
	}
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

func parseFlags() (string, error) {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")

	flag.Parse()
	if err := validateConfigPath(configPath); err != nil {
		return "", err
	}

	return configPath, nil
}

func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}
