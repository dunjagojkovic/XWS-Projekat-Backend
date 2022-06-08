package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	Id          primitive.ObjectID `bson:"_id"`
	User        string             `bson:"user"`
	Image       string             `bson:"image"`
	Description string             `bson:"description"`
	Link        string             `bson:"link"`
	LikeList    []string           `bson:"like_list"`
	DislikeList []string           `bson:"dislike_list"`
	CommentList []Comment          `bson:"comment_list"`
}
