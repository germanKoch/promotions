package api

import "promotions/model"

type PromotionService interface {
	GetById(Id uint) model.Promotion
}

type PromotionController struct {
	service PromotionService
}

func PromotionControllerV1(service PromotionService) PromotionController {
	return PromotionController{
		service: service,
	}
}
