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

func (uc *UserController) EditPassword(ctx context.Context, request *pb.EditPasswordRequest) (*pb.User, error) {

	user, err := uc.service.EditPassword(request.Password.NewPassword, request.Password.OldPassword, request.Password.Username)
	if err != nil {
		return nil, err
	}
	userPb := mapEditedUser(user)
	return userPb, nil

}

func (uc *UserController) EditPrivacy(ctx context.Context, request *pb.EditPrivacyRequest) (*pb.User, error) {

	user, err := uc.service.EditPrivacy(request.Privacy.IsPublic, request.Privacy.Username)
	if err != nil {
		return nil, err
	}
	userPb := mapEditedUser(user)
	return userPb, nil

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

func (uc *UserController) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.User, error) {

	id := request.Id
	objID, err := primitive.ObjectIDFromHex(id)
	user, err := uc.service.GetUser(objID)

	if err != nil {
		return nil, err
	}
	userPb := mapEditedUser(&user)

	return userPb, nil
}

func (uc *UserController) BlockUser(ctx context.Context, request *pb.BlockUserRequest) (*pb.GetUserRequest, error) {

	blockUser := mapNewBlock(request)

	id, err := uc.service.BlockUser(blockUser)

	if err != nil {
		return nil, err
	}
	response := &pb.GetUserRequest{
		Id: id.Hex(),
	}

	return response, nil
}

func (uc *UserController) Unblock(ctx context.Context, request *pb.BlockUserRequest) (*pb.GetUserRequest, error) {

	blockUser := mapNewBlock(request)

	id, err := uc.service.Unblock(blockUser)

	if err != nil {
		return nil, err
	}
	response := &pb.GetUserRequest{
		Id: id.Hex(),
	}

	return response, nil
}

func (uc *UserController) GetUsers(ctx context.Context, request *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	users, err := uc.service.GetUsers()
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

func (uc *UserController) GetUsersById(ctx context.Context, request *pb.GetUsersByIdRequest) (*pb.GetUsersResponse, error) {
	usersById := request.UserById
	var users []string
	for _, user := range usersById {
		fmt.Println(user)
		users = append(users, user.Id)
	}
	result, err := uc.service.GetUsersById(users)
	if err != nil {
		return nil, err
	}
	response := &pb.GetUsersResponse{
		Users: []*pb.User{},
	}
	for _, user := range result {
		current := mapUser(user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (uc *UserController) GetPublicUsers(ctx context.Context, request *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	users, err := uc.service.GetPublicUsers()
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
	fmt.Println(request.User.WorkExperience.Description)
	workExperience := mapWorkExperience(request.User.WorkExperience)
	editedUser, err := uc.service.EditUser(user, workExperience)

	if err != nil {
		return nil, err
	}
	userPb := mapEditedUser(editedUser)

	return userPb, nil
}

func (uc *UserController) FilterUsers(ctx context.Context, request *pb.FilterUsersRequest) (*pb.FilterUsersResponse, error) {

	searchTerm := request.SearchTerm
	users, err := uc.service.FilterUsers(searchTerm)
	if err != nil {
		return nil, err
	}
	response := &pb.FilterUsersResponse{
		Users: []*pb.User{},
	}
	for _, user := range users {
		current := mapUser(user)
		response.Users = append(response.Users, current)
	}
	return response, nil

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
		BlockedUsers:   make([]model.Block, 0),
	}

	return user
}

func mapNewBlock(blockPb *pb.BlockUserRequest) *model.Block {

	blocked, _ := primitive.ObjectIDFromHex(blockPb.BlockedId)
	blocker, _ := primitive.ObjectIDFromHex(blockPb.BlockerId)
	blockUser := &model.Block{
		Id:        primitive.NewObjectID(),
		BlockedId: blocked,
		BlockerId: blocker,
		Status:    blockPb.Status,
	}

	return blockUser
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
