package service

import (
	"notificationS/model"
	"notificationS/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationService struct {
	store repository.NotificationStoreI
}

func NewNotificationService(store repository.NotificationStoreI) *NotificationService {
	return &NotificationService{
		store: store,
	}
}

func (service *NotificationService) GetByUserId(id primitive.ObjectID) ([]model.Notification, error) {
	return service.store.GetByUserId(id)
}

func (service *NotificationService) ChangeStatus(id primitive.ObjectID, status bool) (primitive.ObjectID, error) {
	return service.store.ChangeStatus(id, status)
}

func (service *NotificationService) CreateNotification(notification *model.Notification) (primitive.ObjectID, error) {
	return service.store.CreateNotification(notification)
}
