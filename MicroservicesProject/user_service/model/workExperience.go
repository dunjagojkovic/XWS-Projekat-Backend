package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type WorkExperience struct {
	Id          primitive.ObjectID `bson:"_id"`
	Description string             `bson:"description"`
}
