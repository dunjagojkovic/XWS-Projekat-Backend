package repository

import (
	"userS/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStoreI interface {
	RegisterUser(user *model.User) (*model.User, error)
	Login(username, password string) (bool, error)
	CurrentUser(username string) (model.User, error)
	GetUser(id primitive.ObjectID) (model.User, error)
	GetUsers() ([]*model.User, error)
	GetPublicUsers() ([]*model.User, error)
	GetUsersById([]string) ([]*model.User, error)
	EditUser(*model.User, *model.WorkExperience) (*model.User, error)
	EditPassword(string, string, string) (*model.User, error)
	EditPrivacy(bool, string) (*model.User, error)
	BlockUser(block *model.Block) (primitive.ObjectID, error)
	Unblock(block *model.Block) (primitive.ObjectID, error)
}
