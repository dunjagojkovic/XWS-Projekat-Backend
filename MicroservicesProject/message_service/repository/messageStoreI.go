package repository

import (
	"messageS/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageStoreI interface {
	GetMessages(id string) ([]model.Message, error)
	GetChats(user string) ([]*model.Chat, error)
	CreateMessage(message *model.Message) (primitive.ObjectID, error)
	ChangeMessageStatus(status, id, chatId string) (primitive.ObjectID, error)
}
