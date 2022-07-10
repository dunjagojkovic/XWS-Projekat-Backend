package service

import (
	"common/tracer"
	"context"
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

func (service *RecommendationService) JobRecommendations(ctx context.Context, id string, experiences []*model.WorkExperience, skills []string, jobOffers []*model.JobOffer) ([]*model.JobsId, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE JobRecommendations")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var recommendations []*model.JobsId

	recommendations, err := service.store.JobRecommendations(ctx, id, experiences, skills, jobOffers)
	if err != nil {
		return nil, nil
	}
	//for _, r := range recommendations {
	//	recommendations = append(recommendations, r)
	//}
	return recommendations, nil
}
