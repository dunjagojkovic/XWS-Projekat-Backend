package repository

import (
	"context"
	"recommendationS/model"
)

type RecommendationStoreI interface {
	JobRecommendations(ctx context.Context, id string, experiences []*model.WorkExperience, skills []string, jobOffers []*model.JobOffer) ([]*model.JobsId, error)
}
