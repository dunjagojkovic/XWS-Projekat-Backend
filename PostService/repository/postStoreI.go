package repository

import (
	"postservice/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostStoreI interface {
	GetAll() ([]*model.Post, error)
	Insert(post *model.Post) (primitive.ObjectID, error)
	InsertComment(id primitive.ObjectID, comment *model.Comment) (primitive.ObjectID, error)
	GetPostComments(id primitive.ObjectID) ([]model.Comment, error)
	InsertPostLike(id primitive.ObjectID, user string) (primitive.ObjectID, error)
	GetPostLikes(id primitive.ObjectID) ([]string, error)
	InsertPostDislike(id primitive.ObjectID, user string) (primitive.ObjectID, error)
	GetPostDislikes(id primitive.ObjectID) ([]string, error)
}
