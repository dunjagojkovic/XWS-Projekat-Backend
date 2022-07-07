package repository

import (
	"common/tracer"
	"context"
	"fmt"
	"postS/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "posts"
	COLLECTION = "post"
)

type PostStore struct {
	posts *mongo.Collection
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

func NewPostStore(client *mongo.Client) PostStoreI {

	posts := client.Database(DATABASE).Collection(COLLECTION)

	return &PostStore{
		posts: posts,
	}
}

func (store *PostStore) GetAll(ctx context.Context) ([]*model.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAll")
	defer span.Finish()

	filter := bson.D{{}}
	return store.filter(filter)

}

func (store *PostStore) Get(id primitive.ObjectID) (model.Post, error) {
	filter := bson.D{{"_id", id}}
	var result model.Post

	err := store.posts.FindOne(context.TODO(), filter).Decode(&result)

	fmt.Println(err)

	return result, nil
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

func (store *PostStore) GetUserPosts(username string) ([]model.Post, error) {
	filter := bson.D{{"user", username}}

	cur, err := store.posts.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var posts []model.Post

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var post model.Post
		err := cur.Decode(&post)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (store *PostStore) CreatePost(post *model.Post) (primitive.ObjectID, error) {
	result, err := store.posts.InsertOne(context.TODO(), post)
	if err != nil {
		return primitive.NewObjectID(), err
	}
	post.Id = result.InsertedID.(primitive.ObjectID)

	return post.Id, nil
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

func (store *PostStore) CommentPost(id primitive.ObjectID, comment *model.Comment) (primitive.ObjectID, error) {

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

func (store *PostStore) LikePost(id primitive.ObjectID, user string) (primitive.ObjectID, error) {

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

func (store *PostStore) DislikePost(id primitive.ObjectID, user string) (primitive.ObjectID, error) {

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

func (store *PostStore) GetFollowingPosts(users []string) ([]model.Post, error) {

	cur, err := store.posts.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	var posts []model.Post

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var post model.Post
		err := cur.Decode(&post)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	var followingPosts []model.Post

	for _, post := range posts {
		for _, user := range users {
			if post.User == user {
				followingPosts = append(followingPosts, post)
			}
		}
	}

	return followingPosts, nil
}
