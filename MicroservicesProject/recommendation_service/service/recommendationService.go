package service

import (
	"recommendationS/model"
	"recommendationS/repository"
)

type RecommendationService struct {
	store repository.RecommendationStoreI
}

func NewRecommendationService(store repository.RecommendationStoreI) *RecommendationService {
	return &RecommendationService{
		store: store,
	}
}

func (service *RecommendationService) JobRecommendations(id string, experiences []*model.WorkExperience, skills []string, jobOffers []*model.JobOffer) ([]*model.JobsId, error) {

	var recommendations []*model.JobsId

	recommendations, err := service.store.JobRecommendations(id, experiences, skills, jobOffers)
	if err != nil {
		return nil, nil
	}
	//for _, r := range recommendations {
	//	recommendations = append(recommendations, r)
	//}
	return recommendations, nil
}
