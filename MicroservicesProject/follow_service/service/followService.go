package service

import (
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
func (service *FollowService) Follows(id string) ([]*model.User, error) {
	return service.store.Follows(id)
}
func (service *FollowService) Followers(id string) ([]*model.User, error) {
	return service.store.Followers(id)
}
func (service *FollowService) FollowRequests(id string) ([]*model.User, error) {
	return service.store.FollowRequests(id)
}
func (service *FollowService) FollowerRequests(id string) ([]*model.User, error) {
	return service.store.FollowerRequests(id)
}
func (service *FollowService) Relationship(followerId string, followedId string) (string, error) {
	return service.store.Relationship(followerId, followedId)
}
func (service *FollowService) Follow(followerId string, followedId string) (string, error) {
	return service.store.Follow(followerId, followedId)
}
func (service *FollowService) FollowRequest(followerId string, followedId string) (string, error) {
	return service.store.FollowRequest(followerId, followedId)
}
func (service *FollowService) AcceptFollow(followerId string, followedId string) (string, error) {
	return service.store.AcceptFollow(followerId, followedId)
}
func (service *FollowService) Unfollow(followerId string, followedId string) (string, error) {
	return service.store.Unfollow(followerId, followedId)
}
func (service *FollowService) FollowRequestRemove(followerId string, followedId string) (string, error) {
	return service.store.FollowRequestRemove(followerId, followedId)
}
func (service *FollowService) Recommended(id string) ([]*model.User, error) {
	return service.store.Recommended(id)
}
