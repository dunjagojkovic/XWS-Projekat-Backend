package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	Id         primitive.ObjectID `bson:"_id"`
	FirstUser  primitive.ObjectID `bson:"first_user"`
	SecondUser primitive.ObjectID `bson:"second_user"`
	Messages   []Message          `bson:"messages"`
}
