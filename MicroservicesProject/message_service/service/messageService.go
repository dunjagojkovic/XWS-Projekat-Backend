package service

import (
	"common/tracer"
	"context"
	"fmt"
	"messageS/model"
	"messageS/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageService struct {
	store        repository.MessageStoreI
	orchestrator *CreateMessageOrchestrator
}

func NewMessageService(store repository.MessageStoreI, orchestrator *CreateMessageOrchestrator) *MessageService {
	return &MessageService{
		store:        store,
		orchestrator: orchestrator,
	}
}

func (service *MessageService) GetMessages(ctx context.Context, id string) ([]model.Message, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetMessages")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.GetMessages(ctx, id)
}

func (service *MessageService) CreateMessage(ctx context.Context, message *model.Message) (primitive.ObjectID, primitive.ObjectID, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE CreateMessage")
	defer span.Finish()

	message.Status = "Pending Sent"
	ctx = tracer.ContextWithSpan(context.Background(), span)
	id, idChat, err := service.store.CreateMessage(ctx, message)

	if err != nil {
		return primitive.NilObjectID, primitive.NilObjectID, err
	}
	err = service.orchestrator.Start(message, idChat.Hex())
	if err != nil {
		res, _ := service.store.ChangeMessageStatus(ctx, "Cancelled", id.Hex(), idChat.Hex())
		fmt.Println(res)
		return primitive.NilObjectID, primitive.NilObjectID, err
	}
	return id, idChat, nil

}

func (service *MessageService) GetChats(ctx context.Context, user string) ([]*model.Chat, []string, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetChats")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.GetChats(ctx, user)
}

func (service *MessageService) ChangeMessageStatus(ctx context.Context, status, id, chatId string) (primitive.ObjectID, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE ChangeMessageStatus")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.ChangeMessageStatus(ctx, status, id, chatId)
}
