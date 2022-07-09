package repository

import (
	"context"
	"postS/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostStoreI interface {
	GetAll(ctx context.Context) ([]*model.Post, error)
	Get(context.Context, primitive.ObjectID) (model.Post, error)
	GetUserPosts(ctx context.Context, username string) ([]model.Post, error)
	CreatePost(context.Context, *model.Post) (primitive.ObjectID, error)
	GetPostComments(ctx context.Context, id primitive.ObjectID) ([]model.Comment, error)
	GetPostLikes(ctx context.Context, id primitive.ObjectID) ([]string, error)
	GetPostDislikes(ctx context.Context, id primitive.ObjectID) ([]string, error)
	CommentPost(ctx context.Context, id primitive.ObjectID, comment *model.Comment) (primitive.ObjectID, error)
	LikePost(ctx context.Context, id primitive.ObjectID, user string) (primitive.ObjectID, error)
	DislikePost(ctx context.Context, id primitive.ObjectID, user string) (primitive.ObjectID, error)
	GetFollowingPosts(ctx context.Context, users []string) ([]model.Post, error)
}
