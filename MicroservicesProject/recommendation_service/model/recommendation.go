package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type JobOffer struct {
	Id              primitive.ObjectID `bson:"_id"`
	Position        string             `bson:"position"`
	Description     string             `bson:"description"`
	DailyActivities string             `bson:"daily_activities"`
	Precondition    string             `bson:"precondition"`
	User            string             `bson:"user"`
}

type JobsId struct {
	Id string
}
