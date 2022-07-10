package controller

import (
	pb "common/proto/follow_service"
	user "common/proto/user_service"
	"common/tracer"
	"context"
	"fmt"
	"followS/service"
	"strconv"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FollowController struct {
	pb.UnimplementedFollowServiceServer
	service            *service.FollowService
	userServiceAddress string
	CustomLogger       *CustomLogger
}

func NewFollowController(service *service.FollowService, userServiceEndpoint string) *FollowController {
	CustomLogger := NewCustomLogger()
	return &FollowController{
		service:            service,
		userServiceAddress: userServiceEndpoint,
		CustomLogger:       CustomLogger,
	}
}

func (fc *FollowController) Follow(ctx context.Context, request *pb.FollowRequest) (*pb.FollowResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER Follow")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	followerId := request.FollowerId
	followedId := request.FollowedId
	userClient := NewUsersClient(fc.userServiceAddress)
	userResponse, err := userClient.GetUser(context.TODO(), &user.GetUserRequest{Id: followedId})
	if err != nil {
		return nil, err
	}
	if !userResponse.IsPublic {
		response, err := fc.service.FollowRequest(ctx, followerId, followedId)
		if err != nil {
			fc.CustomLogger.ErrorLogger.Error("Creating following request between user with ID: " + followerId + " and user with ID: " + followedId + " failed")
			return nil, err
		}
		responsePb := &pb.FollowResponse{Response: response}
		fc.CustomLogger.SuccessLogger.Info("Creating following request between user with ID: " + followerId + " and user with ID: " + followedId + " successful")

		return responsePb, nil
	}

	response, err := fc.service.Follow(ctx, followerId, followedId)
	if err != nil {
		fc.CustomLogger.ErrorLogger.Error("Creating connection between user with ID: " + followerId + " and user with ID: " + followedId + " failed")
		return nil, err
	}
	responsePb := &pb.FollowResponse{Response: response}
	fc.CustomLogger.SuccessLogger.Info("Creating connection between user with ID: " + followerId + " and user with ID: " + followedId + " successful")
	return responsePb, nil
}

func (fc *FollowController) Follows(ctx context.Context, request *pb.FollowsRequest) (*pb.FollowsResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER Follows")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	id := request.Id
	response, err := fc.service.Follows(ctx, id)
	if err != nil {
		fc.CustomLogger.ErrorLogger.Error("Follows for user with ID: " + id + " not found")
		return nil, err
	}
	responsePb := &pb.FollowsResponse{Follows: []*pb.Follower{}}
	for _, user := range response {
		responsePb.Follows = append(responsePb.Follows, &pb.Follower{Id: user.Id, Time: timestamppb.New(user.TimeOfFollow)})
	}

	fc.CustomLogger.SuccessLogger.Info("Found " + strconv.Itoa(len(responsePb.Follows)))
	return responsePb, nil
}

func (fc *FollowController) Followers(ctx context.Context, request *pb.FollowersRequest) (*pb.FollowersResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER GetAll")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	id := request.Id
	response, err := fc.service.Followers(ctx, id)
	if err != nil {
		return nil, err
	}
	responsePb := &pb.FollowersResponse{Followers: []*pb.Follower{}}
	for _, user := range response {
		responsePb.Followers = append(responsePb.Followers, &pb.Follower{Id: user.Id, Time: timestamppb.New(user.TimeOfFollow)})
	}
	return responsePb, nil
}

func (fc *FollowController) Relationships(ctx context.Context, request *pb.RelationshipsRequest) (*pb.RelationshipsResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER Relationships")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	followerId := request.FollowerId
	followedId := request.FollowedId
	response, err := fc.service.Relationship(ctx, followerId, followedId)
	if err != nil {
		return nil, err
	}
	responsePb := &pb.RelationshipsResponse{}
	responsePb.Relationship = response
	return responsePb, nil
}

func (fc *FollowController) AcceptFollow(ctx context.Context, request *pb.AcceptFollowRequest) (*pb.AcceptFollowResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER AcceptFollow")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	followerId := request.FollowerId
	followedId := request.FollowedId
	response, err := fc.service.AcceptFollow(ctx, followerId, followedId)
	if err != nil {
		fc.CustomLogger.ErrorLogger.Error("Connection between user with ID: " + followerId + " and user with ID: " + followedId + " not approved")
		return nil, err
	}
	responsePb := &pb.AcceptFollowResponse{Response: response}
	fc.CustomLogger.SuccessLogger.Info("Follow connection created between user with ID: " + followedId + " and user with ID: " + followerId)
	return responsePb, nil
}

func (fc *FollowController) Unfollow(ctx context.Context, request *pb.UnfollowRequest) (*pb.UnfollowResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER Unfollow")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	followerId := request.FollowerId
	followedId := request.FollowedId
	response, err := fc.service.Unfollow(ctx, followerId, followedId)
	if err != nil {
		fc.CustomLogger.ErrorLogger.Error("Connection between user with ID: " + followerId + " and user with ID: " + followedId + " not removed")
		return nil, err
	}
	responsePb := &pb.UnfollowResponse{Response: response}
	fc.CustomLogger.SuccessLogger.Info("User with ID: " + followerId + " succesfully unfollowed user with ID: " + followedId)
	return responsePb, nil
}

func (fc *FollowController) FollowRequestRemove(ctx context.Context, request *pb.FollowRequestRemoveRequest) (*pb.FollowRequestRemoveResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER FollowRequestRemove")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	followerId := request.FollowerId
	followedId := request.FollowedId
	response, err := fc.service.FollowRequestRemove(ctx, followerId, followedId)
	if err != nil {
		return nil, err
	}
	responsePb := &pb.FollowRequestRemoveResponse{Response: response}
	return responsePb, nil
}

func (fc *FollowController) FollowRequests(ctx context.Context, request *pb.FollowRequestsRequest) (*pb.FollowRequestsResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER FollowRequests")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Println(md)
	id := request.Id
	response, err := fc.service.FollowRequests(ctx, id)
	if err != nil {
		return nil, err
	}
	responsePb := &pb.FollowRequestsResponse{FollowRequests: []*pb.Follower{}}
	for _, user := range response {
		responsePb.FollowRequests = append(responsePb.FollowRequests, &pb.Follower{Id: user.Id, Time: timestamppb.New(user.TimeOfFollow)})
	}
	return responsePb, nil
}

func (fc *FollowController) FollowerRequests(ctx context.Context, request *pb.FollowerRequestsRequest) (*pb.FollowerRequestsResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER FollowerRequests")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	id := request.Id
	response, err := fc.service.FollowerRequests(ctx, id)
	if err != nil {
		fc.CustomLogger.ErrorLogger.Error("Requests for user with ID: " + id + " not found")
		return nil, err
	}

	responsePb := &pb.FollowerRequestsResponse{FollowerRequests: []*pb.Follower{}}
	for _, user := range response {
		responsePb.FollowerRequests = append(responsePb.FollowerRequests, &pb.Follower{Id: user.Id, Time: timestamppb.New(user.TimeOfFollow)})
	}
	fc.CustomLogger.SuccessLogger.Info("Found " + strconv.Itoa(len(responsePb.FollowerRequests)) + " requests")

	return responsePb, nil
}

func (fc *FollowController) GetRecommended(ctx context.Context, request *pb.Id) (*pb.ListId, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER GetRecommended")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	users, err := fc.service.Recommended(ctx, request.Id)

	if err != nil {
		fc.CustomLogger.ErrorLogger.Error("Recommended users not found")
		return nil, err
	}
	result := &pb.ListId{
		ListId: []*pb.Id{},
	}
	for _, user := range users {
		result.ListId = append(result.ListId, &pb.Id{Id: user.Id})
	}
	fc.CustomLogger.SuccessLogger.Info("Found " + strconv.Itoa(len(users)) + " recommended users")
	return result, nil
}
