package repository

import (
	"context"
	"messageS/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageStoreI interface {
	GetMessages(ctx context.Context, id string) ([]model.Message, error)
	GetChats(ctx context.Context, user string) ([]*model.Chat, []string, error)
	CreateMessage(ctx context.Context, message *model.Message) (primitive.ObjectID, primitive.ObjectID, error)
	ChangeMessageStatus(ctx context.Context, status, id, chatId string) (primitive.ObjectID, error)
}
