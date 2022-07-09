package repository

import (
	"context"
	"jobS/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobStoreI interface {
	GetAll(ctx context.Context) ([]*model.JobOffer, error)
	CreateJobOffer(ctx context.Context, job *model.JobOffer) (primitive.ObjectID, error)
	JobOfferSearch(ctx context.Context, position string) ([]model.JobOffer, error)
	GetOwnerJobOffers(ctx context.Context, key string) ([]model.JobOffer, error)
	InsertKey(ctx context.Context, username string, key string) (primitive.ObjectID, error)
}
