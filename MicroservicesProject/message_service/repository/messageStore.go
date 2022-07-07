package repository

import (
	"context"
	"messageS/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "chats"
	COLLECTION = "chat"
)

type MessageStore struct {
	chats *mongo.Collection
}

func NewMessageStore(client *mongo.Client) MessageStoreI {

	chats := client.Database(DATABASE).Collection(COLLECTION)

	return &MessageStore{
		chats: chats,
	}
}

func (store *MessageStore) GetMessages(id string) ([]model.Message, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objID}}

	var result *model.Chat

	store.chats.FindOne(context.TODO(), filter).Decode(&result)

	return result.Messages, nil

}

func (store *MessageStore) GetChats(user string) ([]*model.Chat, []string, error) {
	u, err := primitive.ObjectIDFromHex(user)
	filter1 := bson.D{{"first_user", u}}
	cursor1, err := store.chats.Find(context.TODO(), filter1)
	var list []string
	defer cursor1.Close(context.TODO())

	if err != nil {
		return nil, nil, err
	}
	chats1, err := decode(cursor1)
	if err != nil {
		return nil, nil, err

	}
	filter2 := bson.D{{"second_user", u}}
	cursor2, err := store.chats.Find(context.TODO(), filter2)
	defer cursor2.Close(context.TODO())

	if err != nil {
		return nil, nil, err
	}
	chats2, err := decode(cursor2)
	if err != nil {
		return nil, nil, err

	}
	var result []*model.Chat
	for _, chat1 := range chats1 {
		result = append(result, chat1)
		if chat1.SecondUser != u {
			list = append(list, chat1.SecondUser.Hex())
		}
	}
	for _, chat2 := range chats2 {
		result = append(result, chat2)
		if chat2.FirstUser != u {
			list = append(list, chat2.FirstUser.Hex())
		}
	}
	return result, list, nil
}

func (store *MessageStore) CreateMessage(message *model.Message) (primitive.ObjectID, primitive.ObjectID, error) {

	filter := bson.D{{"first_user", message.Sender}, {"second_user", message.Receiver}}
	var result *model.Chat

	store.chats.FindOne(context.TODO(), filter).Decode(&result)

	if result == nil {
		filter1 := bson.D{{"first_user", message.Receiver}, {"second_user", message.Sender}}
		var result1 *model.Chat

		store.chats.FindOne(context.TODO(), filter1).Decode(&result1)

		if result1 == nil {
			chat := &model.Chat{
				Id:         primitive.NewObjectID(),
				FirstUser:  message.Sender,
				SecondUser: message.Receiver,
				Messages:   make([]model.Message, 0),
			}
			_, err := store.chats.InsertOne(context.TODO(), chat)
			if err != nil {
				return primitive.NilObjectID, primitive.NilObjectID, err
			}

			update := bson.D{
				{"$push", bson.D{
					{"messages", message},
				}},
			}

			store.chats.UpdateOne(context.TODO(), filter, update)

			if err != nil {
				return primitive.NewObjectID(), primitive.NilObjectID, err
			}

			return message.Id, chat.Id, nil
		}
		update1 := bson.D{
			{"$push", bson.D{
				{"messages", message},
			}},
		}

		store.chats.UpdateOne(context.TODO(), filter1, update1)

		return message.Id, result1.Id, nil
	}

	update := bson.D{
		{"$push", bson.D{
			{"messages", message},
		}},
	}

	store.chats.UpdateOne(context.TODO(), filter, update)

	return message.Id, result.Id, nil
}

func (store *MessageStore) filter(filter interface{}) ([]*model.Chat, error) {
	cursor, err := store.chats.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *MessageStore) ChangeMessageStatus(status, id, chatId string) (primitive.ObjectID, error) {

	objID, _ := primitive.ObjectIDFromHex(chatId)
	messID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objID}}
	var chat *model.Chat

	store.chats.FindOne(context.TODO(), filter).Decode(&chat)

	var list []model.Message

	for _, message := range chat.Messages {
		if message.Id == messID {
			message.Status = status

			list = append(list, message)

		} else {
			list = append(list, message)
		}
	}
	chat.Messages = list

	store.chats.FindOneAndReplace(context.TODO(), filter, chat)
	return messID, nil

}

func decode(cursor *mongo.Cursor) (chats []*model.Chat, err error) {
	for cursor.Next(context.TODO()) {
		var chat model.Chat
		err = cursor.Decode(&chat)
		if err != nil {
			return
		}
		chats = append(chats, &chat)
	}
	err = cursor.Err()
	return
}
