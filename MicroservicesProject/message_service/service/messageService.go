package service

import (
	"messageS/model"
	"messageS/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageService struct {
	store repository.MessageStoreI
}

func NewMessageService(store repository.MessageStoreI) *MessageService {
	return &MessageService{
		store: store,
	}
}

func (service *MessageService) GetAllById(id string) ([]*model.Message, error) {
	return service.store.GetAllById(id)
}

func (service *MessageService) CreateMessage(message *model.Message) (primitive.ObjectID, error) {
	return service.store.CreateMessage(message)
}

func (service *MessageService) ChangeMessageStatus(status, id string) (primitive.ObjectID, error) {
	return service.store.ChangeMessageStatus(status, id)
}
