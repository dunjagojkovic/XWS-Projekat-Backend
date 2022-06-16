package controller

import (
	"context"
	"log"
	"os"
	"postS/model"
	"postS/service"
	"time"

	"google.golang.org/grpc/peer"

	pb "common/proto/post_service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func logEntry(logEntryType string, code string, ip string, user string) {
	f, err := os.OpenFile("logs//"+logEntryType+".log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	currentTime := time.Now()
	_, err2 := f.WriteString("[" + currentTime.String() + "] " + code + " | " + ip + " | " + user + " \n")

	if err2 != nil {
		log.Fatal(err2)
	}
}

type PostController struct {
	pb.UnimplementedPostServiceServer
	service *service.PostService
}

func NewPostController(service *service.PostService) *PostController {
	return &PostController{
		service: service,
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

func (pc *PostController) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	posts, err := pc.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, post := range posts {
		current := mapPost(post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
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

func (pc *PostController) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	post, err := pc.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	postPb := mapPost(&post)
	response := &pb.GetResponse{
		Post: postPb,
	}
	return response, nil
}

func (pc *PostController) GetUserPosts(ctx context.Context, request *pb.GetUserPostsRequest) (*pb.GetAllResponse, error) {

	username := request.Username
	posts, err := pc.service.GetUserPosts(username)

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

func (pc *PostController) CreatePost(ctx context.Context, request *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {

	post := mapNewPost(request.Post)
	id, err := pc.service.CreatePost(post)
	if err != nil {
		return nil, err
	}
	p, _ := peer.FromContext(ctx)
	logEntry("notification", "DATA_NP", p.Addr.String(), request.Post.User)
	return &pb.CreatePostResponse{
		Id: id.Hex(),
	}, nil

}

func (pc *PostController) GetPostComments(ctx context.Context, request *pb.GetRequest) (*pb.GetPostCommentsResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	comments, err := pc.service.GetPostComments(objectId)
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
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	likes, err := pc.service.GetPostLikes(objectId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetPostLikesResponse{
		Likes: likes,
	}
	return response, nil
}

func (pc *PostController) GetPostDislikes(ctx context.Context, request *pb.GetRequest) (*pb.GetPostLikesResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	dislikes, err := pc.service.GetPostDislikes(objectId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetPostLikesResponse{
		Likes: dislikes,
	}
	return response, nil
}

func (pc *PostController) CommentPost(ctx context.Context, request *pb.CommentPostRequest) (*pb.CreatePostResponse, error) {

	comment, _id := mapNewComment(request.Comment)
	id, err := pc.service.CommentPost(_id, comment)
	if err != nil {
		return nil, err
	}
	return &pb.CreatePostResponse{
		Id: id.Hex(),
	}, nil

}

func (pc *PostController) LikePost(ctx context.Context, request *pb.LikeDislikePostRequest) (*pb.CreatePostResponse, error) {

	username := request.LikeDislike.Username
	objectId, err := primitive.ObjectIDFromHex(request.LikeDislike.IdPost)
	id, err := pc.service.LikePost(objectId, username)
	if err != nil {
		return nil, err
	}
	return &pb.CreatePostResponse{
		Id: id.Hex(),
	}, nil

}

func (pc *PostController) DislikePost(ctx context.Context, request *pb.LikeDislikePostRequest) (*pb.CreatePostResponse, error) {

	username := request.LikeDislike.Username
	objectId, err := primitive.ObjectIDFromHex(request.LikeDislike.IdPost)
	id, err := pc.service.DislikePost(objectId, username)
	if err != nil {
		return nil, err
	}
	return &pb.CreatePostResponse{
		Id: id.Hex(),
	}, nil

}

func (pc *PostController) GetFollowingPosts(ctx context.Context, request *pb.GetFollowingPostsRequest) (*pb.GetAllResponse, error) {

	users := request.Following.Users
	posts, err := pc.service.GetFollowingPosts(users)
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
