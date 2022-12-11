package main

import (
	"promotions/api"
	"promotions/repository"
	"promotions/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	repo := repository.GetPromotionRepositoryPostgres()
	promotionService := service.GetPromotionRepoService(repo)
	scheduler := service.GetScheduledReader(promotionService)
	promotionController := api.GetPromotionController(promotionService)

	promotionController.GetRouts(app)
	scheduler.ScheduleJob()

	app.Listen(":3000")
}
