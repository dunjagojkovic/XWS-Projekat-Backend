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
	span := tracer.StartSpanFromContext(ctx, "REPOSITORY GetAll")
	defer span.Finish()

	filter := bson.D{{}}
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return store.filter(ctx, filter)

}

func (store *PostStore) Get(ctx context.Context, id primitive.ObjectID) (model.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "REPOSITORY Get")
	defer span.Finish()

	filter := bson.D{{"_id", id}}
	var result model.Post

	err := store.posts.FindOne(context.TODO(), filter).Decode(&result)
	fmt.Println(err)

	return result, nil
}

func (store *PostStore) filter(ctx context.Context, filter interface{}) ([]*model.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "REPOSITORY filter")
	defer span.Finish()

	cursor, err := store.posts.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return nil, err
	}
	return decode(ctx, cursor)
}

func decode(ctx context.Context, cursor *mongo.Cursor) (posts []*model.Post, err error) {
	span := tracer.StartSpanFromContext(ctx, "REPOSITORY decode")
	defer span.Finish()

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

func (store *PostStore) GetUserPosts(ctx context.Context, username string) ([]model.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "REPOSITORY GetUserPosts")
	defer span.Finish()

	filter := bson.D{{"user", username}}

	cur, err := store.posts.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var posts []model.Post

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

func (store *PostStore) CreatePost(ctx context.Context, post *model.Post) (primitive.ObjectID, error) {
	span := tracer.StartSpanFromContext(ctx, "REPOSITORY CreatePost")
	defer span.Finish()

	result, err := store.posts.InsertOne(context.TODO(), post)
	if err != nil {
		return primitive.NewObjectID(), err
	}
	post.Id = result.InsertedID.(primitive.ObjectID)

	return post.Id, nil
}

func (store *PostStore) GetPostComments(ctx context.Context, id primitive.ObjectID) ([]model.Comment, error) {
	span := tracer.StartSpanFromContext(ctx, "REPOSITORY GetPostComments")
	defer span.Finish()

	filter := bson.D{{"_id", id}}
	var result model.Post

	err := store.posts.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}
	return result.CommentList, nil
}

func (store *PostStore) GetPostLikes(ctx context.Context, id primitive.ObjectID) ([]string, error) {
	span := tracer.StartSpanFromContext(ctx, "REPOSITORY GetPostLikes")
	defer span.Finish()

	filter := bson.D{{"_id", id}}
	var result model.Post

	err := store.posts.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}
	return result.LikeList, nil
}

func (store *PostStore) GetPostDislikes(ctx context.Context, id primitive.ObjectID) ([]string, error) {
	span := tracer.StartSpanFromContext(ctx, "REPOSITORY GetPostDislikes")
	defer span.Finish()

	filter := bson.D{{"_id", id}}
	var result model.Post

	err := store.posts.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}
	return result.DislikeList, nil
}

func (store *PostStore) CommentPost(ctx context.Context, id primitive.ObjectID, comment *model.Comment) (primitive.ObjectID, error) {
	span := tracer.StartSpanFromContext(ctx, "REPOSITORY CommentPost")
	defer span.Finish()

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

func (store *PostStore) LikePost(ctx context.Context, id primitive.ObjectID, user string) (primitive.ObjectID, error) {
	span := tracer.StartSpanFromContext(ctx, "REPOSITORY LikePost")
	defer span.Finish()

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

func (store *PostStore) DislikePost(ctx context.Context, id primitive.ObjectID, user string) (primitive.ObjectID, error) {
	span := tracer.StartSpanFromContext(ctx, "REPOSITORY DislikePost")
	defer span.Finish()

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

func (store *PostStore) GetFollowingPosts(ctx context.Context, users []string) ([]model.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "REPOSITORY GetFollowingPosts")
	defer span.Finish()

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
