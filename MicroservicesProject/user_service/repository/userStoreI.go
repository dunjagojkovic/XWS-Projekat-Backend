package repository

import (
	"userS/model"
)

type UserStoreI interface {
	RegisterUser(user *model.User) (*model.User, error)
}
