package repository

import (
	"context"
	"userS/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStoreI interface {
	RegisterUser(ctx context.Context, user *model.User) (*model.User, error)
	Login(ctx context.Context, username, password string) (bool, error)
	CurrentUser(ctx context.Context, username string) (model.User, error)
	GetUser(ctx context.Context, id primitive.ObjectID) (model.User, error)
	GetUsers(ctx context.Context) ([]*model.User, error)
	GetPublicUsers(ctx context.Context) ([]*model.User, error)
	GetUsersById(context.Context, []string) ([]*model.User, error)
	EditUser(context.Context, *model.User, *model.WorkExperience) (*model.User, error)
	EditPassword(context.Context, string, string, string) (*model.User, error)
	EditPrivacy(context.Context, bool, string) (*model.User, error)
	BlockUser(ctx context.Context, block *model.Block) (primitive.ObjectID, error)
	Unblock(ctx context.Context, block *model.Block) (primitive.ObjectID, error)
	CheckBlocking(ctx context.Context, first, second string) bool
}
