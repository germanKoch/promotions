package main

import (
	"promotions/api"
	"promotions/config"
	"promotions/repository"
	"promotions/service"
	"promotions/service/parser"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := config.GetDb()

	repo := repository.GetPromotionRepository(db)
	promotionService := service.GetPromotionRepoService(repo)
	historyRepository := repository.GetHistoryRepository(db)
	promotionParser := parser.GetPromotionParser()

	scheduler := service.GetScheduledReader(promotionService, historyRepository, promotionParser)
	promotionController := api.GetPromotionController(promotionService)

	promotionController.GetRouts(app)
	scheduler.ScheduleJob()

	app.Listen(":3000")
}
