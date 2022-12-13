package api

import (
	"promotions/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Promotion struct {
	Id             string    `json:"id"`
	Price          float64   `json:"price"`
	ExpirationDate time.Time `json:"expirationDate"`
}

type PromotionService interface {
	GetById(Id string) model.Promotion
}

type PromotionController struct {
	service PromotionService
}

func mapToResource(promotion model.Promotion) Promotion {
	return Promotion{
		Id:             promotion.Id,
		Price:          promotion.Price,
		ExpirationDate: promotion.ExpirationDate,
	}
}

func GetPromotionController(service PromotionService) PromotionController {
	return PromotionController{
		service: service,
	}
}

func (cont PromotionController) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := cont.service.GetById(id)
	resource := mapToResource(user)
	return c.Status(200).JSON(resource)
}

func (cont PromotionController) GetRouts(app *fiber.App) {
	app.Get("/promotions/:id", cont.GetUser)
}
