package service

import (
	"postservice/model"
	"postservice/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostService struct {
	store repository.PostStoreI
}

func NewPostService(store repository.PostStoreI) *PostService {
	return &PostService{
		store: store,
	}
}

func (service *PostService) GetAll() ([]*model.Post, error) {
	return service.store.GetAll()
}

func (service *PostService) Insert(post *model.Post) (primitive.ObjectID, error) {
	return service.store.Insert(post)
}

func (service *PostService) InsertComment(id primitive.ObjectID, comment *model.Comment) (primitive.ObjectID, error) {
	return service.store.InsertComment(id, comment)
}

func (service *PostService) GetPostComments(id primitive.ObjectID) ([]model.Comment, error) {
	return service.store.GetPostComments(id)
}

func (service *PostService) InsertPostLike(id primitive.ObjectID, user string) (primitive.ObjectID, error) {
	return service.store.InsertPostLike(id, user)
}

func (service *PostService) GetPostLikes(id primitive.ObjectID) ([]string, error) {
	return service.store.GetPostLikes(id)
}

func (service *PostService) InsertPostDislike(id primitive.ObjectID, user string) (primitive.ObjectID, error) {
	return service.store.InsertPostDislike(id, user)
}

func (service *PostService) GetPostDislikes(id primitive.ObjectID) ([]string, error) {
	return service.store.GetPostDislikes(id)
}
