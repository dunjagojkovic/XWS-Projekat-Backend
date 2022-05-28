package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	Id      primitive.ObjectID `bson:"_id"`
	User    string             `bson:"user"`
	Content string             `bson:"content"`
}
