package controller

import (
	pb "common/proto/user_service"
	"common/tracer"
	"context"
	"userS/model"
	"userS/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	pb.UnimplementedUserServiceServer
	service      *service.UserService
	CustomLogger *CustomLogger
}

func NewUserController(service *service.UserService) *UserController {
	CustomLogger := NewCustomLogger()
	return &UserController{
		service:      service,
		CustomLogger: CustomLogger,
	}

}

func (uc *UserController) Registration(ctx context.Context, request *pb.RegistrationRequest) (*pb.RegistrationResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER Registration")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	user := mapNewUser(ctx, request.User)
	user, err := uc.service.RegisterUser(ctx, user)
	if err != nil {
		uc.CustomLogger.ErrorLogger.Error("User not created")
		return nil, err
	}
	response := &pb.RegistrationResponse{
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
	}
	uc.CustomLogger.SuccessLogger.Info("User registration successfull, created user with ID:" + response.Id)
	return response, nil

}

func (uc *UserController) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER Login")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	token, key, err := uc.service.Login(ctx, request.User.Username, request.User.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		Token: token,
		Key:   key,
	}, nil

}

func (uc *UserController) EditPassword(ctx context.Context, request *pb.EditPasswordRequest) (*pb.User, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER EditPassword")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	user, err := uc.service.EditPassword(ctx, request.Password.NewPassword, request.Password.OldPassword, request.Password.Username)
	if err != nil {
		return nil, err
	}
	userPb := mapEditedUser(ctx, user)
	uc.CustomLogger.SuccessLogger.Info("User with ID: " + user.Id.Hex() + "changed password successfully")

	return userPb, nil

}

func (uc *UserController) EditPrivacy(ctx context.Context, request *pb.EditPrivacyRequest) (*pb.User, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER EditPrivacy")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	user, err := uc.service.EditPrivacy(ctx, request.Privacy.IsPublic, request.Privacy.Username)
	if err != nil {
		return nil, err
	}
	userPb := mapEditedUser(ctx, user)

	uc.CustomLogger.SuccessLogger.Info("User with ID: " + user.Id.Hex() + " updated profile privacy successfully")
	return userPb, nil

}

func (uc *UserController) CurrentUser(ctx context.Context, request *pb.CurrentUserRequest) (*pb.User, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER CurrentUser")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	username := request.Username
	user, err := uc.service.CurrentUser(ctx, username)

	if err != nil {
		uc.CustomLogger.ErrorLogger.Error("User:" + username + " not found")
		return nil, err
	}
	userPb := mapEditedUser(ctx, &user)

	uc.CustomLogger.SuccessLogger.Info("Currently logged in user:" + username)
	return userPb, nil
}

func (uc *UserController) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.User, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER GetUser")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	id := request.Id
	objID, err := primitive.ObjectIDFromHex(id)
	user, err := uc.service.GetUser(ctx, objID)

	if err != nil {
		uc.CustomLogger.ErrorLogger.Error("User with ID:" + objID.Hex() + " not found")
		return nil, err
	}
	userPb := mapEditedUser(ctx, &user)
	uc.CustomLogger.SuccessLogger.Info("User by ID:" + objID.Hex() + " received successfully")
	return userPb, nil
}

func (uc *UserController) BlockUser(ctx context.Context, request *pb.BlockUserRequest) (*pb.GetUserRequest, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER BlockUser")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	blockUser := mapNewBlock(ctx, request)

	id, err := uc.service.BlockUser(ctx, blockUser)

	if err != nil {
		return nil, err
	}
	response := &pb.GetUserRequest{
		Id: id.Hex(),
	}

	uc.CustomLogger.SuccessLogger.Info("Blocking user successfully")
	return response, nil
}

func (uc *UserController) Unblock(ctx context.Context, request *pb.BlockUserRequest) (*pb.GetUserRequest, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER Unblock")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	blockUser := mapNewBlock(ctx, request)

	id, err := uc.service.Unblock(ctx, blockUser)

	if err != nil {
		return nil, err
	}
	response := &pb.GetUserRequest{
		Id: id.Hex(),
	}

	return response, nil
}

func (uc *UserController) GetUsers(ctx context.Context, request *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER GetUsers")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	users, err := uc.service.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	response := &pb.GetUsersResponse{
		Users: []*pb.User{},
	}
	for _, user := range users {
		current := mapUser(ctx, user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (uc *UserController) GetUsersById(ctx context.Context, request *pb.GetUsersByIdRequest) (*pb.GetUsersResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER GetUsersById")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	usersById := request.UserById
	var users []string
	for _, user := range usersById {
		users = append(users, user.Id)
	}
	result, err := uc.service.GetUsersById(ctx, users)
	if err != nil {
		return nil, err
	}
	response := &pb.GetUsersResponse{
		Users: []*pb.User{},
	}
	for _, user := range result {
		current := mapUser(ctx, user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (uc *UserController) GetUserUsernamesById(ctx context.Context, request *pb.GetUserUsernamesByIdRequest) (*pb.GetUsersResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER GetUsersUsernamesById")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	result, err := uc.service.GetUsersById(ctx, request.UserById)
	if err != nil {
		return nil, err
	}
	response := &pb.GetUsersResponse{
		Users: []*pb.User{},
	}
	for _, user := range result {
		current := mapUser(ctx, user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (uc *UserController) GetPublicUsers(ctx context.Context, request *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER GetPublicUsers")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	users, err := uc.service.GetPublicUsers(ctx)
	if err != nil {
		return nil, err
	}
	response := &pb.GetUsersResponse{
		Users: []*pb.User{},
	}
	for _, user := range users {
		current := mapUser(ctx, user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (uc *UserController) EditUser(ctx context.Context, request *pb.EditUserRequest) (*pb.User, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER EditUser")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	user := mapEditUser(ctx, request.User)
	workExperience := mapWorkExperience(ctx, request.User.WorkExperience)
	editedUser, err := uc.service.EditUser(ctx, user, workExperience)

	if err != nil {
		return nil, err
	}
	userPb := mapEditedUser(ctx, editedUser)

	uc.CustomLogger.SuccessLogger.Info("User with ID: " + user.Id.Hex() + "updated successfully")
	return userPb, nil
}

func (uc *UserController) FilterUsers(ctx context.Context, request *pb.FilterUsersRequest) (*pb.FilterUsersResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER FilterUsers")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	searchTerm := request.SearchTerm
	users, err := uc.service.FilterUsers(ctx, searchTerm)
	if err != nil {
		return nil, err
	}
	response := &pb.FilterUsersResponse{
		Users: []*pb.User{},
	}
	for _, user := range users {
		current := mapUser(ctx, user)
		response.Users = append(response.Users, current)
	}
	return response, nil

}

func mapUser(ctx context.Context, user *model.User) *pb.User {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER mapUser")
	defer span.Finish()

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

func mapEditedUser(ctx context.Context, user *model.User) *pb.User {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER mapEditedUser")
	defer span.Finish()

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

func mapNewUser(ctx context.Context, userPb *pb.RegisterUser) *model.User {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER mapNewUser")
	defer span.Finish()

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

func mapNewBlock(ctx context.Context, blockPb *pb.BlockUserRequest) *model.Block {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER mapNewBlock")
	defer span.Finish()

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

func mapEditUser(ctx context.Context, userPb *pb.EditUser) *model.User {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER mapEditUser")
	defer span.Finish()

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

func mapWorkExperience(ctx context.Context, workPb *pb.WorkExperience) *model.WorkExperience {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER mapWorkExperience")
	defer span.Finish()

	work := &model.WorkExperience{
		Id:          primitive.NewObjectID(),
		Description: workPb.Description,
	}
	return work
}
