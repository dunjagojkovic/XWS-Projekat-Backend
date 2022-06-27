package service

import (
	"userS/model"
	"userS/repository"
)

type UserService struct {
	store repository.UserStoreI
}

func NewUserService(store repository.UserStoreI) *UserService {
	return &UserService{
		store: store,
	}
}

func (service *UserService) RegisterUser(user *model.User) (*model.User, error) {
	return service.store.RegisterUser(user)
}
