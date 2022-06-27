package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id             primitive.ObjectID `bson:"_id"`
	Name           string             `bson:"name"`
	Surname        string             `bson:"surname"`
	Email          string             `bson:"email"`
	Username       string             `bson:"username"`
	Password       string             `bson:"password"`
	PhoneNumber    string             `bson:"phone_number"`
	Gender         string             `bson:"gender"`
	IsPublic       bool               `bson:"is_public"`
	Biography      string             `bson:"biography"`
	BirthDate      string             `bson:"birth_date"`
	WorkExperience []WorkExperience   `bson:"work_experiences"`
	Education      string             `bson:"education"`
	Hobby          string             `bson:"hobby"`
	Interest       string             `bson:"interest"`
}
