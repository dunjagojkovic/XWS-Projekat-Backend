package repository

import (
	"context"

	"notificationS/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "notifications"
	COLLECTION = "notification"
)

type NotificationStore struct {
	notifications *mongo.Collection
}

func NewNotificationStore(client *mongo.Client) NotificationStoreI {

	notifications := client.Database(DATABASE).Collection(COLLECTION)

	return &NotificationStore{
		notifications: notifications,
	}
}

func (store *NotificationStore) GetByUserId(id primitive.ObjectID) ([]model.Notification, error) {
	filter := bson.D{{"user_id", id}, {"read", false}}

	cur, err := store.notifications.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var notifications []model.Notification

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var notification model.Notification
		err := cur.Decode(&notification)
		if err != nil {
			return nil, err
		}

		notifications = append(notifications, notification)
	}

	return notifications, nil
}

func (store *NotificationStore) ChangeStatus(id primitive.ObjectID, status bool) (primitive.ObjectID, error) {
	filter := bson.D{{"_id", id}}

	update := bson.D{
		{"$set", bson.D{
			{"read", status},
		}},
	}

	_, err := store.notifications.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return primitive.NilObjectID, err
	}
	return id, err
}

func (store *NotificationStore) CreateNotification(notification *model.Notification) (primitive.ObjectID, error) {
	result, err := store.notifications.InsertOne(context.TODO(), notification)
	if err != nil {
		return primitive.NewObjectID(), err
	}
	id := result.InsertedID.(primitive.ObjectID)

	return id, nil
}

func (store *NotificationStore) CreateNotifications(notifications []model.Notification) ([]primitive.ObjectID, error) {

	var ids []primitive.ObjectID
	for _, notification := range notifications {
		result, err := store.notifications.InsertOne(context.TODO(), notification)

		if err != nil {
			return nil, err
		}
		id := result.InsertedID.(primitive.ObjectID)
		ids = append(ids, id)
	}

	return ids, nil
}
