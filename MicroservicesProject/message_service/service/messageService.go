package service

import (
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

func (service *MessageService) GetMessages(id string) ([]model.Message, error) {
	return service.store.GetMessages(id)
}

func (service *MessageService) CreateMessage(message *model.Message) (primitive.ObjectID, primitive.ObjectID, error) {
	message.Status = "Pending Sent"
	id, idChat, err := service.store.CreateMessage(message)

	if err != nil {
		return primitive.NilObjectID, primitive.NilObjectID, err
	}
	err = service.orchestrator.Start(message, idChat.Hex())
	if err != nil {
		res, _ := service.store.ChangeMessageStatus("Canceled", id.Hex(), idChat.Hex())
		fmt.Println(res)
		return primitive.NilObjectID, primitive.NilObjectID, err
	}
	return id, idChat, nil

}

func (service *MessageService) GetChats(user string) ([]*model.Chat, []string, error) {
	return service.store.GetChats(user)
}

func (service *MessageService) ChangeMessageStatus(status, id, chatId string) (primitive.ObjectID, error) {
	return service.store.ChangeMessageStatus(status, id, chatId)
}
