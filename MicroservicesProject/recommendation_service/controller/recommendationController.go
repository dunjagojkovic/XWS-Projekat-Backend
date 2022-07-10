package controller

import (
	pb "common/proto/recommendation_service"
	"common/tracer"
	"context"
	"fmt"
	"recommendationS/model"
	"recommendationS/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RecommendationController struct {
	pb.UnimplementedRecommendationServiceServer
	service *service.RecommendationService
}

func NewRecommendationController(service *service.RecommendationService) *RecommendationController {

	return &RecommendationController{
		service: service,
	}
}

func (handler *RecommendationController) JobRecommendations(ctx context.Context, request *pb.JobRecommendationsRequest) (*pb.JobRecommendationsResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER GetAll")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	var jobs []*model.JobOffer

	for _, job := range request.JobOffers {
		domainJob := mapJobOffer(ctx, job)
		jobs = append(jobs, domainJob)
	}

	var experiences []*model.WorkExperience

	for _, work := range request.Experiences {
		domainWork := mapWork(ctx, work)
		experiences = append(experiences, domainWork)
	}

	recommendations, err := handler.service.JobRecommendations(ctx, request.Id, experiences, request.Skills, jobs)
	if err != nil {
		return nil, err
	}
	response := &pb.JobRecommendationsResponse{}
	for _, rec := range recommendations {
		response.Ids = append(response.Ids, rec.Id)

		fmt.Println(rec)
	}
	return response, nil
}

func mapJobOffer(ctx context.Context, jobPb *pb.JobOffer) *model.JobOffer {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER mapJobOffer")
	defer span.Finish()

	id, _ := primitive.ObjectIDFromHex(jobPb.Id)

	job := &model.JobOffer{
		Id:              id,
		Position:        jobPb.Position,
		Description:     jobPb.Description,
		DailyActivities: jobPb.DailyActivities,
		Precondition:    jobPb.Precondition,
		User:            jobPb.User,
	}

	return job
}

func mapWork(ctx context.Context, workPb *pb.WorkExperience) *model.WorkExperience {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER GetAll")
	defer span.Finish()

	id, _ := primitive.ObjectIDFromHex(workPb.Id)

	work := &model.WorkExperience{
		Id:          id,
		Description: workPb.Description,
	}

	return work
}
