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

func (store *UserStore) CurrentUser(username string) (model.User, error) {
	filter := bson.D{{"username", username}}
	var result model.User

	err := store.users.FindOne(context.TODO(), filter).Decode(&result)

	fmt.Println(err)

	return result, nil
}

func (store *UserStore) GetUsers() ([]*model.User, error) {
	filter := bson.D{{}}
	return store.filter(filter)

}

func (store *UserStore) GetPublicUsers() ([]*model.User, error) {
	filter := bson.D{{"is_public", true}}
	return store.filter(filter)

}

func (store *UserStore) filter(filter interface{}) ([]*model.User, error) {
	cursor, err := store.users.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func decode(cursor *mongo.Cursor) (users []*model.User, err error) {
	for cursor.Next(context.TODO()) {
		var user model.User
		err = cursor.Decode(&user)
		if err != nil {
			return
		}
		users = append(users, &user)
	}
	err = cursor.Err()
	return
}

func (store *UserStore) EditUser(user *model.User, work *model.WorkExperience) (*model.User, error) {

	filter := bson.D{{"username", user.Username}}

	update := bson.D{
		{"$set", bson.D{
			{"name", user.Name},
			{"surname", user.Surname},
			{"email", user.Email},
			{"phone_number", user.PhoneNumber},
			{"username", user.Username},
			{"biography", user.Biography},
			{"birth_date", user.BirthDate},
			{"education", user.Education},
			{"hobby", user.Hobby},
			{"interest", user.Interest},
			{"gender", user.Gender},
		},
		}, {"$push", bson.D{
			{"work_experiences", work},
		}},
	}

	_, err := store.users.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return nil, err
	}

	findFilter := bson.D{{"username", user.Username}}
	var result model.User

	err1 := store.users.FindOne(context.TODO(), findFilter).Decode(&result)

	return &result, err1
}
