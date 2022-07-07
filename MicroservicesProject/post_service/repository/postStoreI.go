package repository

import (
	"context"
	"postS/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostStoreI interface {
	GetAll(ctx context.Context) ([]*model.Post, error)
	Get(primitive.ObjectID) (model.Post, error)
	GetUserPosts(username string) ([]model.Post, error)
	CreatePost(*model.Post) (primitive.ObjectID, error)
	GetPostComments(id primitive.ObjectID) ([]model.Comment, error)
	GetPostLikes(id primitive.ObjectID) ([]string, error)
	GetPostDislikes(id primitive.ObjectID) ([]string, error)
	CommentPost(id primitive.ObjectID, comment *model.Comment) (primitive.ObjectID, error)
	LikePost(id primitive.ObjectID, user string) (primitive.ObjectID, error)
	DislikePost(id primitive.ObjectID, user string) (primitive.ObjectID, error)
	GetFollowingPosts(users []string) ([]model.Post, error)
}
