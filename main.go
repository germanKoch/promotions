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
	promotionController := api.GetPromotionController(promotionService)

	app.Listen(":3000")
}
