package service

import "promotions/model"

type PromotionRepository interface {
	GetById(Id string) (model.Promotion, error)
	UpsertAll(promotions []model.Promotion)
}

type PromotionService struct {
	repository PromotionRepository
}

func GetPromotionRepoService(repo PromotionRepository) PromotionService {
	return PromotionService{
		repository: repo,
	}
}

func (s PromotionService) GetById(id string) (model.Promotion, error) {
	return s.repository.GetById(id)
}

func (s PromotionService) UpsertAll(promotions []model.Promotion) {
	s.repository.UpsertAll(promotions)
}
