package repository

import (
	"context"
	"fmt"
	"userS/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "users"
	COLLECTION = "user"
)

type UserStore struct {
	users *mongo.Collection
}

func NewUserStore(client *mongo.Client) UserStoreI {

	users := client.Database(DATABASE).Collection(COLLECTION)

	return &UserStore{
		users: users,
	}
}

func (store *UserStore) RegisterUser(user *model.User) (*model.User, error) {
	result, err := store.users.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	fmt.Println(result)
	return user, nil
}

func (store *UserStore) Login(username, password string) (bool, error) {
	filter := bson.D{{"username", username}}

	cur, err := store.users.Find(context.TODO(), filter)
	if err != nil {
		return false, err
	}

	var user model.User
	for cur.Next(context.TODO()) {

		err := cur.Decode(&user)
		if err != nil {
			return false, err
		}

		if user.Password == password {
			return true, nil
		}
	}

	return false, nil
}
