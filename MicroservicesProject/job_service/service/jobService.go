package service

import (
	"common/tracer"
	"context"
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

func (service *JobService) GetAll(ctx context.Context) ([]*model.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetAll")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.GetAll(ctx)
}

func (service *JobService) CreateJobOffer(ctx context.Context, job *model.JobOffer) (primitive.ObjectID, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE CreateJobOffer")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.CreateJobOffer(ctx, job)
}

func (service *JobService) JobOfferSearch(ctx context.Context, position string) ([]model.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE JobOfferSearch")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.JobOfferSearch(ctx, position)
}

func (service *JobService) GetOwnerJobOffers(ctx context.Context, key string) ([]model.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetOwnerJobOffers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.GetOwnerJobOffers(ctx, key)
}

func (service *JobService) InsertKey(ctx context.Context, username string, key string) (primitive.ObjectID, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE InsertKey")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.InsertKey(ctx, username, key)
}
