package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Id       primitive.ObjectID `bson:"_id"`
	Text     string             `bson:"text"`
	Sender   primitive.ObjectID `bson:"sender"`
	Receiver primitive.ObjectID `bson:"receiver"`
	Time     time.Time          `bson:"time"`
	Status   string             `bson:"status"`
}
