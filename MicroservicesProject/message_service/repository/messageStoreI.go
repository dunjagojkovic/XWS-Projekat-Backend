package repository

import (
	"messageS/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageStoreI interface {
	GetAllById(id string) ([]*model.Message, error)
	CreateMessage(message *model.Message) (primitive.ObjectID, error)
	ChangeMessageStatus(status, id string) (primitive.ObjectID, error)
}
