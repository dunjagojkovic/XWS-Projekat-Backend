package repository

import (
	"context"
	"fmt"
	"postservice/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE   = "posts"
	COLLECTION = "post"
)

type PostStore struct {
	posts *mongo.Collection
}

func NewPostStore() PostStoreI {
	host := "localhost"
	port := "27017"
	uri := fmt.Sprintf("mongodb://%s:%s/", host, port)
	options := options.Client().ApplyURI(uri)
	client, _ := mongo.Connect(context.TODO(), options)
	posts := client.Database(DATABASE).Collection(COLLECTION)

	return &PostStore{
		posts: posts,
	}
}

func (store *PostStore) GetAll() ([]*model.Post, error) {
	filter := bson.D{{}}
	return store.filter(filter)

}

func (store *PostStore) Insert(post *model.Post) (primitive.ObjectID, error) {
	result, err := store.posts.InsertOne(context.TODO(), post)
	if err != nil {
		return primitive.NewObjectID(), err
	}
	post.Id = result.InsertedID.(primitive.ObjectID)

	return post.Id, nil
}

func (store *PostStore) InsertComment(id primitive.ObjectID, comment *model.Comment) (primitive.ObjectID, error) {

	filter := bson.D{{"_id", id}}

	update := bson.D{
		{"$push", bson.D{
			{"comment_list", comment},
		}},
	}

	_, err := store.posts.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return primitive.NewObjectID(), err
	}
	return comment.Id, err
}

func (store *PostStore) GetPostComments(id primitive.ObjectID) ([]model.Comment, error) {
	filter := bson.D{{"_id", id}}
	var result model.Post

	err := store.posts.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}
	return result.CommentList, nil
}

func (store *PostStore) InsertPostLike(id primitive.ObjectID, user string) (primitive.ObjectID, error) {

	filter := bson.D{{"_id", id}}

	update := bson.D{
		{"$push", bson.D{
			{"like_list", user},
		}},
	}

	_, err := store.posts.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return primitive.NewObjectID(), err
	}
	return id, err

}

func (store *PostStore) InsertPostDislike(id primitive.ObjectID, user string) (primitive.ObjectID, error) {

	filter := bson.D{{"_id", id}}

	update := bson.D{
		{"$push", bson.D{
			{"dislike_list", user},
		}},
	}

	_, err := store.posts.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return primitive.NewObjectID(), err
	}
	return id, err

}

func (store *PostStore) GetPostLikes(id primitive.ObjectID) ([]string, error) {
	filter := bson.D{{"_id", id}}
	var result model.Post

	err := store.posts.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}
	return result.LikeList, nil
}

func (store *PostStore) GetPostDislikes(id primitive.ObjectID) ([]string, error) {
	filter := bson.D{{"_id", id}}
	var result model.Post

	err := store.posts.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}
	return result.DislikeList, nil
}

func (store *PostStore) filter(filter interface{}) ([]*model.Post, error) {
	cursor, err := store.posts.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func decode(cursor *mongo.Cursor) (posts []*model.Post, err error) {
	for cursor.Next(context.TODO()) {
		var post model.Post
		err = cursor.Decode(&post)
		if err != nil {
			return
		}
		posts = append(posts, &post)
	}
	err = cursor.Err()
	return
}
