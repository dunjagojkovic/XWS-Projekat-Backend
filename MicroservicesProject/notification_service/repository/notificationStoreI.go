package repository

import (
	"notificationS/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationStoreI interface {
	GetByUserId(id primitive.ObjectID) ([]model.Notification, error)
	ChangeStatus(id primitive.ObjectID, status bool) (primitive.ObjectID, error)
	CreateNotification(notification *model.Notification) (primitive.ObjectID, error)
	CreateNotifications(notifications []model.Notification) ([]primitive.ObjectID, error)
}
