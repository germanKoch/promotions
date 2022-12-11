package service

import "promotions/model"

type PromotionRepository interface {
	GetById(Id uint) model.Promotion
	Save(model.Promotion)
}

type PromotionService struct {
	repository PromotionRepository
}

func PromotionRepoService(repo PromotionRepository) PromotionService {
	return PromotionService{
		repository: repo,
	}
}

func (s PromotionService) GetById(id uint) model.Promotion {
	return s.repository.GetById(id)
}

func (s PromotionService) save(promotion model.Promotion) {
	s.repository.Save(promotion)
}
