package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Block struct {
	Id        primitive.ObjectID `bson:"_id"`
	BlockedId primitive.ObjectID `bson:"_blockedId"`
	BlockerId primitive.ObjectID `bson:"_blockerId"`
	Status    string             `bson:"_status"`
}
