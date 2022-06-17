package main

import (
	postGW "common/proto/post_service"
	"context"
	"fmt"
	"log"
	"net"
	"postS/controller"
	"postS/repository"
	"postS/service"

	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	config *Config
}

func NewServer(config *Config) *Server {
	return &Server{
		config: config,
	}
}

func accessibleRoles() map[string][]string {
	const servicePath = "/post.PostService/"
	return map[string][]string{
		//servicePath + "GetAll":            {"user"},
		//servicePath + "Get":               {"user"},
		//servicePath + "GetPostLikes":      {"user"},
		//servicePath + "GetPostDislikes":   {"user"},
		//servicePath + "GetUserPosts":      {"GetUserPosts"},
		servicePath + "CreatePost":        {"CreatePost"},
		servicePath + "CommentPost":       {"CommentPost"},
		servicePath + "LikePost":          {"LikePost"},
		servicePath + "DislikePost":       {"DislikePost"},
		servicePath + "GetFollowingPosts": {"GetFollowingPosts"},
	}
}

func permissionsOfRoles() map[string][]string {
	return map[string][]string{
		"user":  {"CreatePost", "CommentPost", "LikePost", "DislikePost", "GetFollowingPosts", "GetUserPosts"},
		"admin": {},
	}
}

func (server *Server) Start() {

	uri := fmt.Sprintf("mongodb://%s:%s", server.config.PostDBHost, server.config.PostDBPort)
	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), options)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(uri)
	}
	store := repository.NewPostStore(client)
	postService := service.NewPostService(store)

	postController := controller.NewPostController(postService)

	server.startGrpcServer(postController)
}

func (server *Server) startGrpcServer(postController *controller.PostController) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	interceptor := NewAuthInterceptor(accessibleRoles(), permissionsOfRoles())
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)
	postGW.RegisterPostServiceServer(grpcServer, postController)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
