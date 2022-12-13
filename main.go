package main

import (
	"fmt"
	"promotions/service/storage"
	// "promotions/api"
	// "promotions/config"
	// "promotions/repository"
	// "promotions/service"
	// "github.com/gofiber/fiber/v2"
)

func main() {
	stor := storage.GetLocalStorage("C:\\Users\\germi\\Desktop\\test_data")
	stor.Walk(func(file storage.FileData) {
		liner := file.Content()
		for i := 0; liner.HasNext(); i++ {
			s := liner.NextLine()
			fmt.Println(s)
		}
	})
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
