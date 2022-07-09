package repository

import (
	"common/tracer"
	"context"
	"jobS/model"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "jobs"
	COLLECTION = "job"
)

type JobStore struct {
	jobs *mongo.Collection
}

func NewJobStore(client *mongo.Client) JobStoreI {

	jobs := client.Database(DATABASE).Collection(COLLECTION)

	return &JobStore{
		jobs: jobs,
	}
}

func (store *JobStore) GetAll(ctx context.Context) ([]*model.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "REPOSITORY GetAll")
	defer span.Finish()

	filter := bson.D{{}}
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return store.filter(ctx, filter)

}

func (store *JobStore) CreateJobOffer(ctx context.Context, job *model.JobOffer) (primitive.ObjectID, error) {
	span := tracer.StartSpanFromContext(ctx, "REPOSITORY CreateJobOffer")
	defer span.Finish()

	result, err := store.jobs.InsertOne(context.TODO(), job)
	if err != nil {
		return primitive.NewObjectID(), err
	}
	job.Id = result.InsertedID.(primitive.ObjectID)

	return job.Id, nil
}

func (store *JobStore) filter(ctx context.Context, filter interface{}) ([]*model.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "REPOSITORY filter")
	defer span.Finish()

	cursor, err := store.jobs.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *JobStore) JobOfferSearch(ctx context.Context, position string) ([]model.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "REPOSITORY JobOfferSearch")
	defer span.Finish()

	cur, err := store.jobs.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	var jobs []model.JobOffer

	for cur.Next(context.TODO()) {
		var job model.JobOffer
		err := cur.Decode(&job)
		if err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	var foundJobs []model.JobOffer

	for _, job := range jobs {

		if strings.Contains(strings.ToLower(job.Position), strings.ToLower(position)) {
			foundJobs = append(foundJobs, job)
		}

	}

	return foundJobs, nil
}

func (store *JobStore) GetOwnerJobOffers(ctx context.Context, key string) ([]model.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "REPOSITORY GetOwnerJobOffers")
	defer span.Finish()

	cur, err := store.jobs.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	var jobs []model.JobOffer

	for cur.Next(context.TODO()) {
		var job model.JobOffer
		err := cur.Decode(&job)
		if err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	var ownerJobs []model.JobOffer

	for _, job := range jobs {
		if job.Key == key {
			ownerJobs = append(ownerJobs, job)
		}

	}

	return ownerJobs, nil
}

func (store *JobStore) InsertKey(ctx context.Context, username string, key string) (primitive.ObjectID, error) {
	span := tracer.StartSpanFromContext(ctx, "REPOSITORY InsertKey")
	defer span.Finish()

	filter := bson.D{{"user", username}}

	update := bson.D{
		{"$set", bson.D{
			{"key", key},
		}},
	}

	_, err := store.jobs.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return primitive.NewObjectID(), nil
}

func decode(cursor *mongo.Cursor) (jobs []*model.JobOffer, err error) {
	for cursor.Next(context.TODO()) {
		var job model.JobOffer
		err = cursor.Decode(&job)
		if err != nil {
			return
		}
		jobs = append(jobs, &job)
	}
	err = cursor.Err()
	return
}
