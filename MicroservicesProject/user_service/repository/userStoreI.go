package repository

import (
	"userS/model"
)

type UserStoreI interface {
	RegisterUser(user *model.User) (*model.User, error)
	Login(username, password string) (bool, error)
	CurrentUser(username string) (model.User, error)
	GetUsers() ([]*model.User, error)
	GetPublicUsers() ([]*model.User, error)
	EditUser(*model.User, *model.WorkExperience) (*model.User, error)
	EditPassword(string, string, string) (*model.User, error)
	EditPrivacy(bool, string) (*model.User, error)

}
