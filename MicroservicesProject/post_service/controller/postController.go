package controller

import (
	"context"
	"postS/model"
	"postS/service"
	"strconv"

	pb "common/proto/post_service"
	"common/tracer"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostController struct {
	pb.UnimplementedPostServiceServer
	service      *service.PostService
	CustomLogger *CustomLogger
}

func NewPostController(service *service.PostService) *PostController {
	CustomLogger := NewCustomLogger()
	return &PostController{
		service:      service,
		CustomLogger: CustomLogger,
	}

}

func mapNewPost(postPb *pb.CreatePost) *model.Post {

	post := &model.Post{
		Id:          primitive.NewObjectID(),
		User:        postPb.User,
		Image:       postPb.Image,
		Description: postPb.Description,
		Link:        postPb.Link,
		LikeList:    make([]string, 0),
		DislikeList: make([]string, 0),
		CommentList: make([]model.Comment, 0),
	}

	for _, commentPb := range postPb.CommentList {
		commID, _ := primitive.ObjectIDFromHex(commentPb.Id)
		comment := model.Comment{
			Id:      commID,
			User:    commentPb.User,
			Content: commentPb.Content,
		}
		post.CommentList = append(post.CommentList, comment)
	}

	return post
}

func mapNewComment(commentPb *pb.PostComment) (*model.Comment, primitive.ObjectID) {

	comment := &model.Comment{
		Id:      primitive.NewObjectID(),
		User:    commentPb.User,
		Content: commentPb.Content,
	}
	postID, _ := primitive.ObjectIDFromHex(commentPb.IdPost)
	return comment, postID
}

func mapPost(post *model.Post) *pb.Post {
	postPb := &pb.Post{
		Id:          post.Id.Hex(),
		User:        post.User,
		Image:       post.Image,
		Description: post.Description,
		Link:        post.Link,
		LikeList:    post.LikeList,
		DislikeList: post.DislikeList,
		CommentList: make([]*pb.Comment, 0),
	}

	for _, comment := range post.CommentList {

		commentPb := *&pb.Comment{
			Id:      comment.Id.Hex(),
			User:    comment.User,
			Content: comment.Content,
		}

		postPb.CommentList = append(postPb.CommentList, &commentPb)
	}

	return postPb
}

func mapComment(comment *model.Comment) *pb.Comment {
	commentPb := &pb.Comment{
		Id:      comment.Id.Hex(),
		User:    comment.User,
		Content: comment.Content,
	}

	return commentPb
}

func (pc *PostController) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "GetAll")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	posts, err := pc.service.GetAll(ctx)

	if err != nil {
		pc.CustomLogger.ErrorLogger.Error("Get all posts unsuccessful")
		return nil, err
	}
	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, post := range posts {
		current := mapPost(post)
		response.Posts = append(response.Posts, current)
	}
	pc.CustomLogger.SuccessLogger.Info("Found " + strconv.Itoa(len(posts)) + " posts")

	return response, nil
}

func (pc *PostController) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "Get")
	defer span.Finish()

	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		pc.CustomLogger.ErrorLogger.Error("ObjectId not created with ID:" + id)
		return nil, err
	}

	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "MongoGet")
	post, err := pc.service.Get(objectId)
	span1.Finish()

	if err != nil {
		pc.CustomLogger.ErrorLogger.Error("Post with ID: " + id + " not found")
		return nil, err
	}
	postPb := mapPost(&post)
	response := &pb.GetResponse{
		Post: postPb,
	}
	pc.CustomLogger.SuccessLogger.Info("Post with ID: " + id + " received successfully")
	return response, nil
}

func (pc *PostController) GetUserPosts(ctx context.Context, request *pb.GetUserPostsRequest) (*pb.GetAllResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserPosts")
	defer span.Finish()

	username := request.Username
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "MongoGetUserPosts")
	posts, err := pc.service.GetUserPosts(username)
	span1.Finish()

	if err != nil {
		pc.CustomLogger.ErrorLogger.Error("Get all by user: " + username)
		return nil, err
	}
	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, post := range posts {
		current := mapPost(&post)
		response.Posts = append(response.Posts, current)
	}
	pc.CustomLogger.SuccessLogger.Info("Found " + strconv.Itoa(len(posts)) + " posts created by user: " + username)
	return response, nil

}

func (pc *PostController) CreatePost(ctx context.Context, request *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePost")
	defer span.Finish()

	post := mapNewPost(request.Post)
	id, err := pc.service.CreatePost(post)

	if err != nil {
		pc.CustomLogger.ErrorLogger.Error("ObjectId not created with ID:" + post.Id.Hex())
		return nil, err
	}

	response := &pb.CreatePostResponse{
		Id: id.Hex(),
	}
	pc.CustomLogger.SuccessLogger.Info("Post with ID: " + post.Id.Hex() + " created succesfully by user with ID: " + post.User)
	return response, nil

}

func (pc *PostController) GetPostComments(ctx context.Context, request *pb.GetRequest) (*pb.GetPostCommentsResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostComments")
	defer span.Finish()

	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "MongoGetPostComments")
	comments, err := pc.service.GetPostComments(objectId)
	span1.Finish()

	if err != nil {
		return nil, err
	}
	response := &pb.GetPostCommentsResponse{
		Comments: []*pb.Comment{},
	}
	for _, comment := range comments {
		current := mapComment(&comment)
		response.Comments = append(response.Comments, current)
	}
	return response, nil
}

func (pc *PostController) GetPostLikes(ctx context.Context, request *pb.GetRequest) (*pb.GetPostLikesResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostLikes")
	defer span.Finish()

	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "MongoGetPostLikes")
	likes, err := pc.service.GetPostLikes(objectId)
	span1.Finish()

	if err != nil {
		return nil, err
	}
	response := &pb.GetPostLikesResponse{
		Likes: likes,
	}
	return response, nil
}

func (pc *PostController) GetPostDislikes(ctx context.Context, request *pb.GetRequest) (*pb.GetPostLikesResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostDislikes")
	defer span.Finish()

	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "MongoGetPostDislikes")
	dislikes, err := pc.service.GetPostDislikes(objectId)
	span1.Finish()

	if err != nil {
		return nil, err
	}
	response := &pb.GetPostLikesResponse{
		Likes: dislikes,
	}
	return response, nil
}

func (pc *PostController) CommentPost(ctx context.Context, request *pb.CommentPostRequest) (*pb.CreatePostResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CommentPost")
	defer span.Finish()

	comment, _id := mapNewComment(request.Comment)
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "MongoCommentPost")
	id, err := pc.service.CommentPost(_id, comment)
	span1.Finish()

	if err != nil {
		return nil, err
	}
	return &pb.CreatePostResponse{
		Id: id.Hex(),
	}, nil

}

func (pc *PostController) LikePost(ctx context.Context, request *pb.LikeDislikePostRequest) (*pb.CreatePostResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "LikePost")
	defer span.Finish()

	username := request.LikeDislike.Username
	objectId, err := primitive.ObjectIDFromHex(request.LikeDislike.IdPost)
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "MongoLikePost")
	id, err := pc.service.LikePost(objectId, username)
	span1.Finish()

	if err != nil {
		pc.CustomLogger.ErrorLogger.Error("Post with ID: " + objectId.Hex() + " was not succesfully liked by user: " + username)
		return nil, err
	}
	response := &pb.CreatePostResponse{
		Id: id.Hex(),
	}
	pc.CustomLogger.SuccessLogger.Info("Post with ID: " + objectId.Hex() + " liked by user with ID: " + username)
	return response, nil

}

func (pc *PostController) DislikePost(ctx context.Context, request *pb.LikeDislikePostRequest) (*pb.CreatePostResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "DislikePost")
	defer span.Finish()

	username := request.LikeDislike.Username
	objectId, err := primitive.ObjectIDFromHex(request.LikeDislike.IdPost)
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "MongoDislikePost")
	id, err := pc.service.DislikePost(objectId, username)
	span1.Finish()

	if err != nil {
		pc.CustomLogger.ErrorLogger.Error("Post with ID: " + objectId.Hex() + " was not disliked by user: " + username)
		return nil, err
	}
	response := &pb.CreatePostResponse{
		Id: id.Hex(),
	}
	pc.CustomLogger.SuccessLogger.Info("Post with ID: " + objectId.Hex() + " disliked by user with ID: " + username)
	return response, nil

}

func (pc *PostController) GetFollowingPosts(ctx context.Context, request *pb.GetFollowingPostsRequest) (*pb.GetAllResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetFollowingPosts")
	defer span.Finish()

	users := request.Following.Users
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "MongoGetFollowingPosts")
	posts, err := pc.service.GetFollowingPosts(users)
	span1.Finish()

	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, post := range posts {
		current := mapPost(&post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil

}
