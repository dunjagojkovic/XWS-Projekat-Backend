package controller

import (
	pb "common/proto/user_service"
	"context"
	"userS/model"
	"userS/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	pb.UnimplementedUserServiceServer
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{
		service: service,
	}

}

func (uc *UserController) Registration(ctx context.Context, request *pb.RegistrationRequest) (*pb.RegistrationResponse, error) {

	user := mapNewUser(request.User)
	user, err := uc.service.RegisterUser(user)
	if err != nil {
		return nil, err
	}
	return &pb.RegistrationResponse{
		Id:          user.Id.Hex(),
		Name:        user.Name,
		Surname:     user.Surname,
		Email:       user.Email,
		Username:    user.Username,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Gender:      user.Gender,
		IsPublic:    user.IsPublic,
		BirthDate:   user.BirthDate,
	}, nil

}

func (uc *UserController) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {

	token, key, err := uc.service.Login(request.User.Username, request.User.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		Token: token,
		Key:   key,
	}, nil

}

func mapNewUser(userPb *pb.RegisterUser) *model.User {

	user := &model.User{
		Id:             primitive.NewObjectID(),
		Name:           userPb.Name,
		Surname:        userPb.Surname,
		Email:          userPb.Email,
		Username:       userPb.Username,
		Password:       userPb.Password,
		PhoneNumber:    userPb.PhoneNumber,
		Gender:         userPb.Gender,
		IsPublic:       userPb.IsPublic,
		Biography:      "",
		BirthDate:      userPb.BirthDate,
		WorkExperience: make([]model.WorkExperience, 0),
		Education:      "",
		Hobby:          "",
		Interest:       "",
	}

	return user
}
