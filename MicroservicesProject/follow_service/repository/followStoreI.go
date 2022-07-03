package repository

import (
	"followS/model"
)

type FollowStoreI interface {
	Follows(id string) ([]*model.User, error)
	Followers(id string) ([]*model.User, error)
	FollowRequests(id string) ([]*model.User, error)
	FollowerRequests(id string) ([]*model.User, error)
	Relationship(followerId string, followedId string) (string, error)
	Follow(followerId string, followedId string) (string, error)
	FollowRequest(followerId string, followedId string) (string, error)
	AcceptFollow(followerId string, followedId string) (string, error)
	Unfollow(followerId string, followedId string) (string, error)
	FollowRequestRemove(followerId string, followedId string) (string, error)
	Recommended(id string) ([]*model.User, error)
}
