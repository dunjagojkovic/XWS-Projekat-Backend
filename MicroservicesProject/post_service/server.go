package main

import (
	postGW "common/proto/post_service"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"postS/controller"
	"postS/repository"
	"postS/service"

	"google.golang.org/grpc"

	"common/tracer"

	otgo "github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Name = "post_service"
)

type Server struct {
	config *Config
	Tracer otgo.Tracer
	Closer io.Closer
}

func NewServer(config *Config) *Server {
	tracer, closer := tracer.Init(Name)
	otgo.SetGlobalTracer(tracer)
	return &Server{
		config: config,
		Tracer: tracer,
		Closer: closer,
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
	grpcServer := grpc.NewServer()
	postGW.RegisterPostServiceServer(grpcServer, postController)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
