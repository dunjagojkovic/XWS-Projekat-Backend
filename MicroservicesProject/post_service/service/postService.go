package service

import (
	"postS/model"
	"postS/repository"

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

func (service *PostService) Get(id primitive.ObjectID) (model.Post, error) {
	return service.store.Get(id)
}

func (service *PostService) GetUserPosts(username string) ([]model.Post, error) {
	return service.store.GetUserPosts(username)
}

func (service *PostService) CreatePost(post *model.Post) (primitive.ObjectID, error) {
	return service.store.CreatePost(post)
}

func (service *PostService) GetPostComments(id primitive.ObjectID) ([]model.Comment, error) {
	return service.store.GetPostComments(id)
}

func (service *PostService) GetPostLikes(id primitive.ObjectID) ([]string, error) {
	return service.store.GetPostLikes(id)
}

func (service *PostService) GetPostDislikes(id primitive.ObjectID) ([]string, error) {
	return service.store.GetPostDislikes(id)
}

func (service *PostService) CommentPost(id primitive.ObjectID, comment *model.Comment) (primitive.ObjectID, error) {
	return service.store.CommentPost(id, comment)
}

func (service *PostService) LikePost(id primitive.ObjectID, user string) (primitive.ObjectID, error) {
	return service.store.LikePost(id, user)
}

func (service *PostService) DislikePost(id primitive.ObjectID, user string) (primitive.ObjectID, error) {
	return service.store.DislikePost(id, user)
}

func (service *PostService) GetFollowingPosts(users []string) ([]model.Post, error) {
	return service.store.GetFollowingPosts(users)
}
