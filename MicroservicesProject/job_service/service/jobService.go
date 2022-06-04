package service

import (
	"jobS/model"
	"jobS/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobService struct {
	store repository.JobStoreI
}

func NewJobService(store repository.JobStoreI) *JobService {
	return &JobService{
		store: store,
	}
}

func (service *JobService) GetAll() ([]*model.JobOffer, error) {
	return service.store.GetAll()
}

func (service *JobService) CreateJobOffer(job *model.JobOffer) (primitive.ObjectID, error) {
	return service.store.CreateJobOffer(job)
}

func (service *JobService) JobOfferSearch(position string) ([]model.JobOffer, error) {
	return service.store.JobOfferSearch(position)
}

func (service *JobService) GetOwnerJobOffers(usernames []string) ([]model.JobOffer, error) {
	return service.store.GetOwnerJobOffers(usernames)
}
