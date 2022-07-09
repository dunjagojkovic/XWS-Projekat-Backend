package repository

import "recommendationS/model"

type RecommendationStoreI interface {
	JobRecommendations(id string, experiences []*model.WorkExperience, skills []string, jobOffers []*model.JobOffer) ([]*model.JobsId, error)
}
