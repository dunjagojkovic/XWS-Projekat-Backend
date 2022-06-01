package controller

import (
	pb "common/proto/job_service"
	"context"
	"fmt"
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
	fmt.Println("usla")
	jobs, err := jc.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Offers: []*pb.JobOffer{},
	}
	for _, job := range jobs {
		current := mapJob(job)
		response.Offers = append(response.Offers, current)
	}
	return response, nil
}

func (jc *JobController) CreateJobOffer(ctx context.Context, request *pb.CreateJobOfferRequest) (*pb.CreateJobOfferResponse, error) {

	job := mapNewJob(request.Job)
	id, err := jc.service.CreateJobOffer(job)
	if err != nil {
		return nil, err
	}
	return &pb.CreateJobOfferResponse{
		Id: id.Hex(),
	}, nil

}

func (jc *JobController) JobOfferSearch(ctx context.Context, request *pb.JobOfferSearchRequest) (*pb.GetAllResponse, error) {

	position := request.Search.Position
	jobs, err := jc.service.JobOfferSearch(position)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Offers: []*pb.JobOffer{},
	}
	for _, job := range jobs {
		current := mapJob(&job)
		response.Offers = append(response.Offers, current)
	}
	return response, nil

}

func mapNewJob(jobPb *pb.CreateJobOffer) *model.JobOffer {

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

func mapJob(job *model.JobOffer) *pb.JobOffer {
	jobPb := &pb.JobOffer{
		Id:              job.Id.Hex(),
		Position:        job.Position,
		Description:     job.Description,
		DailyActivities: job.DailyActivities,
		Precondition:    job.Precondition,
		User:            job.User,
	}

	return jobPb
}