package service

import (
	"common/tracer"
	"context"
	"followS/model"
	"followS/repository"
)

type FollowService struct {
	store repository.FollowStoreI
}

func NewFollowService(store repository.FollowStoreI) *FollowService {
	return &FollowService{
		store: store,
	}
}
func (service *FollowService) Follows(ctx context.Context, id string) ([]*model.User, error) {
	return service.store.Follows(ctx, id)
}
func (service *FollowService) Followers(ctx context.Context, id string) ([]*model.User, error) {
	return service.store.Followers(ctx, id)
}
func (service *FollowService) FollowRequests(ctx context.Context, id string) ([]*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetAll")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.FollowRequests(ctx, id)
}
func (service *FollowService) FollowerRequests(ctx context.Context, id string) ([]*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetAll")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.FollowerRequests(ctx, id)
}
func (service *FollowService) Relationship(ctx context.Context, followerId string, followedId string) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetAll")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.Relationship(ctx, followerId, followedId)
}
func (service *FollowService) Follow(ctx context.Context, followerId string, followedId string) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetAll")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.Follow(ctx, followerId, followedId)
}
func (service *FollowService) FollowRequest(ctx context.Context, followerId string, followedId string) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetAll")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.FollowRequest(ctx, followerId, followedId)
}
func (service *FollowService) AcceptFollow(ctx context.Context, followerId string, followedId string) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetAll")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.AcceptFollow(ctx, followerId, followedId)
}
func (service *FollowService) Unfollow(ctx context.Context, followerId string, followedId string) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetAll")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.Unfollow(ctx, followerId, followedId)
}
func (service *FollowService) FollowRequestRemove(ctx context.Context, followerId string, followedId string) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetAll")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.FollowRequestRemove(ctx, followerId, followedId)
}
func (service *FollowService) Recommended(ctx context.Context, id string) ([]*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetAll")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.Recommended(ctx, id)
}
