package repository

import (
	"context"
	"fmt"
	"userS/model"

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
