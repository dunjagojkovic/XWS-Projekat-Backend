package repository

import (
	"context"
	"followS/model"
)

type FollowStoreI interface {
	Follows(ctx context.Context, id string) ([]*model.User, error)
	Followers(ctx context.Context, id string) ([]*model.User, error)
	FollowRequests(ctx context.Context, id string) ([]*model.User, error)
	FollowerRequests(ctx context.Context, id string) ([]*model.User, error)
	Relationship(ctx context.Context, followerId string, followedId string) (string, error)
	Follow(ctx context.Context, followerId string, followedId string) (string, error)
	FollowRequest(ctx context.Context, followerId string, followedId string) (string, error)
	AcceptFollow(ctx context.Context, followerId string, followedId string) (string, error)
	Unfollow(ctx context.Context, followerId string, followedId string) (string, error)
	FollowRequestRemove(ctx context.Context, followerId string, followedId string) (string, error)
	Recommended(ctx context.Context, id string) ([]*model.User, error)
}
