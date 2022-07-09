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

func mapNewPost(ctx context.Context, postPb *pb.CreatePost) *model.Post {
	span := tracer.StartSpanFromContext(ctx, "mapNewPost")
	defer span.Finish()

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

func mapNewComment(ctx context.Context, commentPb *pb.PostComment) (*model.Comment, primitive.ObjectID) {
	span := tracer.StartSpanFromContext(ctx, "mapNewComment")
	defer span.Finish()

	comment := &model.Comment{
		Id:      primitive.NewObjectID(),
		User:    commentPb.User,
		Content: commentPb.Content,
	}
	postID, _ := primitive.ObjectIDFromHex(commentPb.IdPost)
	return comment, postID
}

func mapPost(ctx context.Context, post *model.Post) *pb.Post {
	span := tracer.StartSpanFromContext(ctx, "mapPost")
	defer span.Finish()

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

func mapComment(ctx context.Context, comment *model.Comment) *pb.Comment {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER mapComment")
	defer span.Finish()

	commentPb := &pb.Comment{
		Id:      comment.Id.Hex(),
		User:    comment.User,
		Content: comment.Content,
	}

	return commentPb
}

func (pc *PostController) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER GetAll")
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
		current := mapPost(ctx, post)
		response.Posts = append(response.Posts, current)
	}
	pc.CustomLogger.SuccessLogger.Info("Found " + strconv.Itoa(len(posts)) + " posts")

	return response, nil
}

func (pc *PostController) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER Get")
	defer span.Finish()

	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		pc.CustomLogger.ErrorLogger.Error("ObjectId not created with ID:" + id)
		return nil, err
	}
	ctx = tracer.ContextWithSpan(context.Background(), span)
	post, err := pc.service.Get(ctx, objectId)

	if err != nil {
		pc.CustomLogger.ErrorLogger.Error("Post with ID: " + id + " not found")
		return nil, err
	}
	postPb := mapPost(ctx, &post)
	response := &pb.GetResponse{
		Post: postPb,
	}
	pc.CustomLogger.SuccessLogger.Info("Post with ID: " + id + " received successfully")
	return response, nil
}

func (pc *PostController) GetUserPosts(ctx context.Context, request *pb.GetUserPostsRequest) (*pb.GetAllResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER GetUserPosts")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	username := request.Username
	posts, err := pc.service.GetUserPosts(ctx, username)

	if err != nil {
		pc.CustomLogger.ErrorLogger.Error("Get all by user: " + username + "unsuccessfully")
		return nil, err
	}
	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, post := range posts {
		current := mapPost(ctx, &post)
		response.Posts = append(response.Posts, current)
	}
	pc.CustomLogger.SuccessLogger.Info("Found " + strconv.Itoa(len(posts)) + " posts created by user: " + username)
	return response, nil

}

func (pc *PostController) CreatePost(ctx context.Context, request *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER CreatePost")
	span.SetOperationName("CONTROLLER CreatePost")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	post := mapNewPost(ctx, request.Post)
	id, err := pc.service.CreatePost(ctx, post)

	if err != nil {
		pc.CustomLogger.ErrorLogger.Error("ObjectId not created with ID:" + post.Id.Hex())
		return nil, err
	}

	response := &pb.CreatePostResponse{
		Id: id.Hex(),
	}
	pc.CustomLogger.SuccessLogger.Info("Post with ID: " + post.Id.Hex() + " created successfully by user with ID: " + post.User)
	return response, nil

}

func (pc *PostController) GetPostComments(ctx context.Context, request *pb.GetRequest) (*pb.GetPostCommentsResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER GetPostComments")
	defer span.Finish()

	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	ctx = tracer.ContextWithSpan(context.Background(), span)
	comments, err := pc.service.GetPostComments(ctx, objectId)

	if err != nil {
		return nil, err
	}
	response := &pb.GetPostCommentsResponse{
		Comments: []*pb.Comment{},
	}
	for _, comment := range comments {
		current := mapComment(ctx, &comment)
		response.Comments = append(response.Comments, current)
	}
	return response, nil
}

func (pc *PostController) GetPostLikes(ctx context.Context, request *pb.GetRequest) (*pb.GetPostLikesResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER GetPostLikes")
	defer span.Finish()

	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	ctx = tracer.ContextWithSpan(context.Background(), span)
	likes, err := pc.service.GetPostLikes(ctx, objectId)

	if err != nil {
		return nil, err
	}
	response := &pb.GetPostLikesResponse{
		Likes: likes,
	}
	return response, nil
}

func (pc *PostController) GetPostDislikes(ctx context.Context, request *pb.GetRequest) (*pb.GetPostLikesResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER GetPostDislikes")
	defer span.Finish()

	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	ctx = tracer.ContextWithSpan(context.Background(), span)
	dislikes, err := pc.service.GetPostDislikes(ctx, objectId)

	if err != nil {
		return nil, err
	}
	response := &pb.GetPostLikesResponse{
		Likes: dislikes,
	}
	return response, nil
}

func (pc *PostController) CommentPost(ctx context.Context, request *pb.CommentPostRequest) (*pb.CreatePostResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER CommentPost")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	comment, _id := mapNewComment(ctx, request.Comment)
	id, err := pc.service.CommentPost(ctx, _id, comment)

	if err != nil {
		return nil, err
	}
	return &pb.CreatePostResponse{
		Id: id.Hex(),
	}, nil

}

func (pc *PostController) LikePost(ctx context.Context, request *pb.LikeDislikePostRequest) (*pb.CreatePostResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER LikePost")
	defer span.Finish()

	username := request.LikeDislike.Username
	objectId, err := primitive.ObjectIDFromHex(request.LikeDislike.IdPost)
	ctx = tracer.ContextWithSpan(context.Background(), span)
	id, err := pc.service.LikePost(ctx, objectId, username)

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
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER DislikePost")
	defer span.Finish()

	username := request.LikeDislike.Username
	objectId, err := primitive.ObjectIDFromHex(request.LikeDislike.IdPost)
	ctx = tracer.ContextWithSpan(context.Background(), span)
	id, err := pc.service.DislikePost(ctx, objectId, username)

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
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER GetFollowingPosts")
	defer span.Finish()

	users := request.Following.Users
	ctx = tracer.ContextWithSpan(context.Background(), span)
	posts, err := pc.service.GetFollowingPosts(ctx, users)

	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, post := range posts {
		current := mapPost(ctx, &post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil

}
