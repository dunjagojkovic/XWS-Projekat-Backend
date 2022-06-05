package repository

import (
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

func (store *JobStore) GetAll() ([]*model.JobOffer, error) {
	filter := bson.D{{}}
	return store.filter(filter)

}

func (store *JobStore) CreateJobOffer(job *model.JobOffer) (primitive.ObjectID, error) {
	result, err := store.jobs.InsertOne(context.TODO(), job)
	if err != nil {
		return primitive.NewObjectID(), err
	}
	job.Id = result.InsertedID.(primitive.ObjectID)

	return job.Id, nil
}

func (store *JobStore) filter(filter interface{}) ([]*model.JobOffer, error) {
	cursor, err := store.jobs.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *JobStore) JobOfferSearch(position string) ([]model.JobOffer, error) {

	cur, err := store.jobs.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	var jobs []model.JobOffer

	// Iterate through the cursor
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

func (store *JobStore) GetOwnerJobOffers(usernames []string) ([]model.JobOffer, error) {

	cur, err := store.jobs.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	var jobs []model.JobOffer

	// Iterate through the cursor
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
		for _, username := range usernames {
			if job.User == username {
				ownerJobs = append(ownerJobs, job)
			}
		}
	}

	return ownerJobs, nil
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
