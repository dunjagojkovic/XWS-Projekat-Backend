package repository

import (
	"context"
	"fmt"
	"messageS/model"
	"sort"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "messages"
	COLLECTION = "message"
)

type MessageStore struct {
	messages *mongo.Collection
}

func NewMessageStore(client *mongo.Client) MessageStoreI {

	messages := client.Database(DATABASE).Collection(COLLECTION)

	return &MessageStore{
		messages: messages,
	}
}

func (store *MessageStore) GetAllById(id string) ([]*model.Message, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	filterSender := bson.D{{"sender", objID}}
	senderMessages, err := store.filter(filterSender)
	if err != nil {
		return nil, err
	}
	filterReceiver := bson.D{{"receiver", objID}}
	receiverMessages, err := store.filter(filterReceiver)
	if err != nil {
		return nil, err
	}

	var result []*model.Message

	for _, rM := range receiverMessages {
		fmt.Println(rM.Id)

		result = append(result, rM)
	}
	for _, sM := range senderMessages {
		fmt.Println(sM.Id)
		result = append(result, sM)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Time.Before(result[j].Time)
	})

	return result, nil

}

func (store *MessageStore) CreateMessage(message *model.Message) (primitive.ObjectID, error) {
	result, err := store.messages.InsertOne(context.TODO(), message)
	if err != nil {
		return primitive.NewObjectID(), err
	}
	message.Id = result.InsertedID.(primitive.ObjectID)

	return message.Id, nil
}

func (store *MessageStore) filter(filter interface{}) ([]*model.Message, error) {
	cursor, err := store.messages.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *MessageStore) ChangeMessageStatus(status, id string) (primitive.ObjectID, error) {

	objID, err := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objID}}

	update := bson.D{
		{"$set", bson.D{
			{"status", status},
		}},
	}

	store.messages.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return objID, nil
}

func decode(cursor *mongo.Cursor) (messages []*model.Message, err error) {
	for cursor.Next(context.TODO()) {
		var message model.Message
		err = cursor.Decode(&message)
		if err != nil {
			return
		}
		messages = append(messages, &message)
	}
	err = cursor.Err()
	return
}
