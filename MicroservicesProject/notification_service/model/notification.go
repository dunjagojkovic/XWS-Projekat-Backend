package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	Id     primitive.ObjectID `bson:"_id"`
	Text   string             `bson:"text"`
	Time   time.Time          `bson:"time"`
	UserId primitive.ObjectID `bson:"user_id"`
	Read   bool               `bson:"read"`
}
