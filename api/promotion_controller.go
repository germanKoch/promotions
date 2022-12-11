package api

import (
	"promotions/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Promotion struct {
	Id             uint
	ExternalId     string
	Price          float32
	ExpirationDate time.Time
}

type PromotionService interface {
	GetById(Id uint) model.Promotion
}

type PromotionController struct {
	service PromotionService
}

func mapToResource(promotion model.Promotion) Promotion {
	return Promotion{
		Id:             promotion.Id,
		ExternalId:     promotion.ExternalId,
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
	id, err := c.ParamsInt("id")
	user := cont.service.GetById(uint(id))
	resource := mapToResource(user)
	return c.Status(200).JSON(resource)
}

func (cont PromotionController) GetRouters(app *fiber.App) {
	app.Get("/api/users/:id", cont.GetUser)
}
