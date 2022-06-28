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
	if user.Name != "" {
		update := bson.D{
			{"$set", bson.D{
				{"name", user.Name},
			},
			},
		}

		_, err := store.users.UpdateOne(context.TODO(), filter, update)

		if err != nil {
			return nil, err
		}
	}
	if user.Surname != "" {
		update := bson.D{
			{"$set", bson.D{
				{"surname", user.Surname},
			},
			},
		}

		_, err := store.users.UpdateOne(context.TODO(), filter, update)

		if err != nil {
			return nil, err
		}
	}
	if user.Email != "" {
		update := bson.D{
			{"$set", bson.D{
				{"email", user.Email},
			},
			},
		}

		_, err := store.users.UpdateOne(context.TODO(), filter, update)

		if err != nil {
			return nil, err
		}
	}
	if user.Username != "" {
		update := bson.D{
			{"$set", bson.D{
				{"username", user.Username},
			},
			},
		}

		_, err := store.users.UpdateOne(context.TODO(), filter, update)

		if err != nil {
			return nil, err
		}
	}
	if user.PhoneNumber != "" {
		update := bson.D{
			{"$set", bson.D{
				{"phone_number", user.PhoneNumber},
			},
			},
		}

		_, err := store.users.UpdateOne(context.TODO(), filter, update)

		if err != nil {
			return nil, err
		}
	}
	if user.BirthDate != "" {
		update := bson.D{
			{"$set", bson.D{
				{"birth_date", user.BirthDate},
			},
			},
		}

		_, err := store.users.UpdateOne(context.TODO(), filter, update)

		if err != nil {
			return nil, err
		}
	}
	if user.Biography != "" {
		update := bson.D{
			{"$set", bson.D{
				{"biography", user.Biography},
			},
			},
		}

		_, err := store.users.UpdateOne(context.TODO(), filter, update)

		if err != nil {
			return nil, err
		}
	}
	if user.Education != "" {
		update := bson.D{
			{"$set", bson.D{
				{"education", user.Education},
			},
			},
		}

		_, err := store.users.UpdateOne(context.TODO(), filter, update)

		if err != nil {
			return nil, err
		}
	}
	if user.Hobby != "" {
		update := bson.D{
			{"$set", bson.D{
				{"hobby", user.Hobby},
			},
			},
		}

		_, err := store.users.UpdateOne(context.TODO(), filter, update)

		if err != nil {
			return nil, err
		}
	}
	if user.Interest != "" {
		update := bson.D{
			{"$set", bson.D{
				{"interest", user.Interest},
			},
			},
		}

		_, err := store.users.UpdateOne(context.TODO(), filter, update)

		if err != nil {
			return nil, err
		}
	}
	if user.Gender != "" {
		update := bson.D{
			{"$set", bson.D{
				{"gender", user.Gender},
			},
			},
		}

		_, err := store.users.UpdateOne(context.TODO(), filter, update)

		if err != nil {
			return nil, err
		}
	}
	if work.Description != "" {
		update := bson.D{
			{"$push", bson.D{
				{"work_experiences", work},
			},
			},
		}

		_, err := store.users.UpdateOne(context.TODO(), filter, update)

		if err != nil {
			return nil, err
		}
	}

	findFilter := bson.D{{"username", user.Username}}
	var result model.User

	err1 := store.users.FindOne(context.TODO(), findFilter).Decode(&result)

	return &result, err1
}

func (store *UserStore) EditPassword(newPassword, oldPassword, username string) (*model.User, error) {

	filter := bson.D{{"username", username}}

	var user model.User

	err := store.users.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		return nil, err
	}

	if user.Password != oldPassword {
		return nil, nil
	}

	if newPassword != "" {
		update := bson.D{
			{"$set", bson.D{
				{"password", newPassword},
			},
			},
		}

		_, err := store.users.UpdateOne(context.TODO(), filter, update)

		if err != nil {
			return nil, err
		}
	}

	findFilter := bson.D{{"username", user.Username}}
	var result model.User

	err1 := store.users.FindOne(context.TODO(), findFilter).Decode(&result)

	return &result, err1
}
