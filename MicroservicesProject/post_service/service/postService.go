package service

import (
	"common/tracer"
	"context"
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

func (service *PostService) GetAll(ctx context.Context) ([]*model.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetAll")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.GetAll(ctx)
}

func (service *PostService) Get(ctx context.Context, id primitive.ObjectID) (model.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE Get")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.Get(ctx, id)
}

func (service *PostService) GetUserPosts(ctx context.Context, username string) ([]model.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetUserPosts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.GetUserPosts(ctx, username)
}

func (service *PostService) CreatePost(ctx context.Context, post *model.Post) (primitive.ObjectID, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE CreatePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.CreatePost(ctx, post)
}

func (service *PostService) GetPostComments(ctx context.Context, id primitive.ObjectID) ([]model.Comment, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetPostComments")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.GetPostComments(ctx, id)
}

func (service *PostService) GetPostLikes(ctx context.Context, id primitive.ObjectID) ([]string, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetPostLikes")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.GetPostLikes(ctx, id)
}

func (service *PostService) GetPostDislikes(ctx context.Context, id primitive.ObjectID) ([]string, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetPostDislikes")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.GetPostDislikes(ctx, id)
}

func (service *PostService) CommentPost(ctx context.Context, id primitive.ObjectID, comment *model.Comment) (primitive.ObjectID, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE CommentPost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.CommentPost(ctx, id, comment)
}

func (service *PostService) LikePost(ctx context.Context, id primitive.ObjectID, user string) (primitive.ObjectID, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE LikePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.LikePost(ctx, id, user)
}

func (service *PostService) DislikePost(ctx context.Context, id primitive.ObjectID, user string) (primitive.ObjectID, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE DislikePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.DislikePost(ctx, id, user)
}

func (service *PostService) GetFollowingPosts(ctx context.Context, users []string) ([]model.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetFollowingPosts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.GetFollowingPosts(ctx, users)
}
