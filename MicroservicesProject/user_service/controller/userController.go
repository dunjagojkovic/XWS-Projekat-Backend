package controller

import (
	pb "common/proto/user_service"
	"context"
	"fmt"
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

	fmt.Println(request.User.Username)
	token, key, err := uc.service.Login(request.User.Username, request.User.Password)
	if err != nil {
		return nil, err
	}
	fmt.Println(token)
	return &pb.LoginResponse{
		Token: token,
		Key:   key,
	}, nil

}

func mapUser(user *model.User) *pb.User {

	userPb := &pb.User{
		Id:              user.Id.Hex(),
		Name:            user.Name,
		Surname:         user.Surname,
		Email:           user.Email,
		Username:        user.Username,
		Password:        user.Password,
		PhoneNumber:     user.PhoneNumber,
		Gender:          user.Gender,
		IsPublic:        user.IsPublic,
		BirthDate:       user.BirthDate,
		Biography:       user.Biography,
		WorkExperiences: make([]*pb.WorkExperience, 0),
	}

	for _, workExperience := range user.WorkExperience {

		workPb := *&pb.WorkExperience{
			Id:          workExperience.Id.Hex(),
			Description: workExperience.Description,
		}

		userPb.WorkExperiences = append(userPb.WorkExperiences, &workPb)
	}

	return userPb
}

func mapEditedUser(user *model.User) *pb.User {

	userPb := &pb.User{
		Id:              user.Id.Hex(),
		Name:            user.Name,
		Surname:         user.Surname,
		Email:           user.Email,
		Username:        user.Username,
		Password:        user.Password,
		PhoneNumber:     user.PhoneNumber,
		Gender:          user.Gender,
		IsPublic:        user.IsPublic,
		BirthDate:       user.BirthDate,
		Biography:       user.Biography,
		WorkExperiences: make([]*pb.WorkExperience, 0),
		Education:       user.Education,
		Hobby:           user.Hobby,
		Interest:        user.Interest,
	}

	for _, workExperience := range user.WorkExperience {

		workPb := *&pb.WorkExperience{
			Id:          workExperience.Id.Hex(),
			Description: workExperience.Description,
		}

		userPb.WorkExperiences = append(userPb.WorkExperiences, &workPb)
	}

	return userPb
}

func (uc *UserController) CurrentUser(ctx context.Context, request *pb.CurrentUserRequest) (*pb.User, error) {

	username := request.Username
	user, err := uc.service.CurrentUser(username)

	if err != nil {
		return nil, err
	}
	userPb := mapEditedUser(&user)

	return userPb, nil
}

func (pc *UserController) GetUsers(ctx context.Context, request *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	users, err := pc.service.GetUsers()
	if err != nil {
		return nil, err
	}
	response := &pb.GetUsersResponse{
		Users: []*pb.User{},
	}
	for _, user := range users {
		current := mapUser(user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (pc *UserController) GetPublicUsers(ctx context.Context, request *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	users, err := pc.service.GetPublicUsers()
	if err != nil {
		return nil, err
	}
	response := &pb.GetUsersResponse{
		Users: []*pb.User{},
	}
	for _, user := range users {
		current := mapUser(user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (uc *UserController) EditUser(ctx context.Context, request *pb.EditUserRequest) (*pb.User, error) {

	fmt.Println(request.User.Username)
	user := mapEditUser(request.User)
	fmt.Println(user.Education)
	fmt.Println(request.User.Education)
	workExperience := mapWorkExperience(request.User.WorkExperience)
	editedUser, err := uc.service.EditUser(user, workExperience)

	if err != nil {
		return nil, err
	}
	userPb := mapEditedUser(editedUser)

	return userPb, nil
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
		IsPublic:       true,
		Biography:      "",
		BirthDate:      userPb.BirthDate,
		WorkExperience: make([]model.WorkExperience, 0),
		Education:      "",
		Hobby:          "",
		Interest:       "",
	}

	return user
}

func mapEditUser(userPb *pb.EditUser) *model.User {

	user := &model.User{
		Name:        userPb.Name,
		Surname:     userPb.Surname,
		Email:       userPb.Email,
		Username:    userPb.Username,
		PhoneNumber: userPb.PhoneNumber,
		Gender:      userPb.Gender,
		Biography:   userPb.Biography,
		BirthDate:   userPb.BirthDate,
		Education:   userPb.Education,
		Hobby:       userPb.Hobby,
		Interest:    userPb.Interest,
	}
	return user
}

func mapWorkExperience(workPb *pb.WorkExperience) *model.WorkExperience {

	work := &model.WorkExperience{
		Id:          primitive.NewObjectID(),
		Description: workPb.Description,
	}
	return work
}
