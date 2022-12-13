package main

import (
	"promotions/api"
	"promotions/config"
	"promotions/repository"
	"promotions/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := config.GetDb()

	repo := repository.GetPromotionRepository(db)
	promotionService := service.GetPromotionRepoService(repo)
	scheduler := service.GetScheduledReader(promotionService)
	promotionController := api.GetPromotionController(promotionService)

	promotionController.GetRouts(app)
	scheduler.ScheduleJob()

	app.Listen(":3000")
}
