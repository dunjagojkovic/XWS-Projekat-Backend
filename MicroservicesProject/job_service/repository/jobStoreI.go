package repository

import (
	"jobS/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobStoreI interface {
	GetAll() ([]*model.JobOffer, error)
	CreateJobOffer(job *model.JobOffer) (primitive.ObjectID, error)
	JobOfferSearch(position string) ([]model.JobOffer, error)
	GetOwnerJobOffers(key string) ([]model.JobOffer, error)
	InsertKey(username string, key string) (primitive.ObjectID, error)
	GetById(ids []string) ([]model.JobOffer, error)
}
