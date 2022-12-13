package main

import (
	"fmt"
	"promotions/service/parser"
	// "promotions/api"
	// "promotions/config"
	// "promotions/repository"
	// "promotions/service"
	// "github.com/gofiber/fiber/v2"
)

func main() {
	p := parser.PromotionParser("d018ef0b-dbd9-48f1-ac1a-eb4d90e57118,60.683466,2018-08-04 05:32:31 +0200 CEST")
	fmt.Print(p)
	// app := fiber.New()
	// db := config.GetDb()

	// repo := repository.GetPromotionRepository(db)
	// promotionService := service.GetPromotionRepoService(repo)
	// scheduler := service.GetScheduledReader(promotionService)
	// promotionController := api.GetPromotionController(promotionService)

	// promotionController.GetRouts(app)
	// scheduler.ScheduleJob()

	// app.Listen(":3000")
}
