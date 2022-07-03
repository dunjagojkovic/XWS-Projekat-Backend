package controller

import (
	pb "common/proto/follow_service"
	user "common/proto/user_service"
	"context"
	"fmt"
	"followS/service"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FollowController struct {
	pb.UnimplementedFollowServiceServer
	service            *service.FollowService
	userServiceAddress string
}

func NewFollowController(service *service.FollowService, userServiceEndpoint string) *FollowController {
	return &FollowController{
		service:            service,
		userServiceAddress: userServiceEndpoint,
	}
}

func (fc *FollowController) Follow(ctx context.Context, request *pb.FollowRequest) (*pb.FollowResponse, error) {

	fmt.Println("Follow")
	followerId := request.FollowerId
	followedId := request.FollowedId
	userClient := NewUsersClient(fc.userServiceAddress)
	userResponse, err := userClient.GetUser(context.TODO(), &user.GetUserRequest{Id: followedId})
	if err != nil {
		return nil, err
	}
	fmt.Println(userResponse.Name)
	if !userResponse.IsPublic {
		response, err := fc.service.FollowRequest(followerId, followedId)
		if err != nil {
			return nil, err
		}
		responsePb := &pb.FollowResponse{Response: response}
		return responsePb, nil
	}

	response, err := fc.service.Follow(followerId, followedId)
	if err != nil {
		return nil, err
	}
	responsePb := &pb.FollowResponse{Response: response}
	return responsePb, nil
}

func (fc *FollowController) Follows(ctx context.Context, request *pb.FollowsRequest) (*pb.FollowsResponse, error) {
	id := request.Id
	response, err := fc.service.Follows(id)
	if err != nil {
		return nil, err
	}
	responsePb := &pb.FollowsResponse{Follows: []*pb.Follower{}}
	for _, user := range response {
		responsePb.Follows = append(responsePb.Follows, &pb.Follower{Id: user.Id, Time: timestamppb.New(user.TimeOfFollow)})
	}
	return responsePb, nil
}

func (fc *FollowController) Followers(ctx context.Context, request *pb.FollowersRequest) (*pb.FollowersResponse, error) {
	id := request.Id
	fmt.Println(id)
	response, err := fc.service.Followers(id)
	if err != nil {
		return nil, err
	}
	responsePb := &pb.FollowersResponse{Followers: []*pb.Follower{}}
	for _, user := range response {
		fmt.Println(user.Id)
		responsePb.Followers = append(responsePb.Followers, &pb.Follower{Id: user.Id, Time: timestamppb.New(user.TimeOfFollow)})
	}
	return responsePb, nil
}

func (fc *FollowController) Relationships(ctx context.Context, request *pb.RelationshipsRequest) (*pb.RelationshipsResponse, error) {
	followerId := request.FollowerId
	followedId := request.FollowedId
	response, err := fc.service.Relationship(followerId, followedId)
	if err != nil {
		return nil, err
	}
	responsePb := &pb.RelationshipsResponse{}
	responsePb.Relationship = response
	return responsePb, nil
}

func (fc *FollowController) AcceptFollow(ctx context.Context, request *pb.AcceptFollowRequest) (*pb.AcceptFollowResponse, error) {

	followerId := request.FollowerId
	followedId := request.FollowedId
	response, err := fc.service.AcceptFollow(followerId, followedId)
	if err != nil {
		return nil, err
	}
	responsePb := &pb.AcceptFollowResponse{Response: response}
	return responsePb, nil
}

func (fc *FollowController) Unfollow(ctx context.Context, request *pb.UnfollowRequest) (*pb.UnfollowResponse, error) {

	followerId := request.FollowerId
	followedId := request.FollowedId
	response, err := fc.service.Unfollow(followerId, followedId)
	if err != nil {
		return nil, err
	}
	responsePb := &pb.UnfollowResponse{Response: response}
	return responsePb, nil
}

func (fc *FollowController) FollowRequestRemove(ctx context.Context, request *pb.FollowRequestRemoveRequest) (*pb.FollowRequestRemoveResponse, error) {

	followerId := request.FollowerId
	followedId := request.FollowedId
	response, err := fc.service.FollowRequestRemove(followerId, followedId)
	if err != nil {
		return nil, err
	}
	responsePb := &pb.FollowRequestRemoveResponse{Response: response}
	return responsePb, nil
}

func (fc *FollowController) FollowRequests(ctx context.Context, request *pb.FollowRequestsRequest) (*pb.FollowRequestsResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Println(md)
	id := request.Id
	response, err := fc.service.FollowRequests(id)
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
	fmt.Println("Radi")
	id := request.Id
	fmt.Println(id)
	response, err := fc.service.FollowerRequests(id)
	if err != nil {
		return nil, err
	}

	responsePb := &pb.FollowerRequestsResponse{FollowerRequests: []*pb.Follower{}}
	for _, user := range response {
		fmt.Println(user.Id)
		responsePb.FollowerRequests = append(responsePb.FollowerRequests, &pb.Follower{Id: user.Id, Time: timestamppb.New(user.TimeOfFollow)})
	}
	return responsePb, nil
}

func (fc *FollowController) GetRecommended(ctx context.Context, request *pb.Id) (*pb.ListId, error) {

	users, err := fc.service.Recommended(request.Id)

	if err != nil {
		return nil, err
	}
	result := &pb.ListId{
		ListId: []*pb.Id{},
	}
	for _, user := range users {
		result.ListId = append(result.ListId, &pb.Id{Id: user.Id})
	}
	return result, nil
}
