package controller

import (
	pb "common/proto/job_service"
	"common/tracer"
	"context"
	"jobS/model"
	"jobS/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobController struct {
	pb.UnimplementedJobServiceServer
	service *service.JobService
}

func NewJobController(service *service.JobService) *JobController {
	return &JobController{
		service: service,
	}

}

func (jc *JobController) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER GetAll")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	jobs, err := jc.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Offers: []*pb.JobOffer{},
	}
	for _, job := range jobs {
		current := mapJob(ctx, job)
		response.Offers = append(response.Offers, current)
	}
	return response, nil
}

func (jc *JobController) OwnerJobOffers(ctx context.Context, request *pb.OwnerJobOffersRequest) (*pb.GetAllResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER OwnerJobOffers")
	defer span.Finish()

	key := request.Key.OwnerKey
	ctx = tracer.ContextWithSpan(context.Background(), span)
	jobs, err := jc.service.GetOwnerJobOffers(ctx, key)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Offers: []*pb.JobOffer{},
	}
	for _, job := range jobs {
		current := mapJob(ctx, &job)
		response.Offers = append(response.Offers, current)
	}
	return response, nil
}

func (jc *JobController) CreateJobOffer(ctx context.Context, request *pb.CreateJobOfferRequest) (*pb.CreateJobOfferResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER CreateJobOffer")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	job := mapNewJob(ctx, request.Job)
	id, err := jc.service.CreateJobOffer(ctx, job)
	if err != nil {
		return nil, err
	}
	return &pb.CreateJobOfferResponse{
		Id: id.Hex(),
	}, nil

}

func (jc *JobController) AddKey(ctx context.Context, request *pb.AddKeyRequest) (*pb.GetAllRequest, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER AddKey")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	username := request.OfferKey.Username
	key := request.OfferKey.Key
	_, err := jc.service.InsertKey(ctx, username, key)
	if err != nil {
		return nil, err
	}
	return &pb.GetAllRequest{}, nil

}

func (jc *JobController) JobOfferSearch(ctx context.Context, request *pb.JobOfferSearchRequest) (*pb.GetAllResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER JobOfferSearch")
	defer span.Finish()

	position := request.Search.Position
	ctx = tracer.ContextWithSpan(context.Background(), span)
	jobs, err := jc.service.JobOfferSearch(ctx, position)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Offers: []*pb.JobOffer{},
	}
	for _, job := range jobs {
		current := mapJob(ctx, &job)
		response.Offers = append(response.Offers, current)
	}
	return response, nil

}

func mapNewJob(ctx context.Context, jobPb *pb.CreateJobOffer) *model.JobOffer {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER mapNewJob")
	defer span.Finish()

	job := &model.JobOffer{
		Id:              primitive.NewObjectID(),
		Position:        jobPb.Position,
		Description:     jobPb.Description,
		DailyActivities: jobPb.DailyActivities,
		Precondition:    jobPb.Precondition,
		User:            jobPb.User,
	}

	return job
}

func mapJob(ctx context.Context, job *model.JobOffer) *pb.JobOffer {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER mapJob")
	defer span.Finish()

	jobPb := &pb.JobOffer{
		Id:              job.Id.Hex(),
		Position:        job.Position,
		Description:     job.Description,
		DailyActivities: job.DailyActivities,
		Precondition:    job.Precondition,
		User:            job.User,
		Key:             job.Key,
	}

	return jobPb
}
