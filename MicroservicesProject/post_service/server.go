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

	otgrpc "github.com/opentracing-contrib/go-grpc"

	"google.golang.org/grpc"

	"common/tracer"
	"postS/config"

	otgo "github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Name = "post_service"
)

type Server struct {
	config       *config.Config
	Tracer       otgo.Tracer
	Closer       io.Closer
	CustomLogger *controller.CustomLogger
}

func NewServer(config *config.Config) *Server {
	CustomLogger := controller.NewCustomLogger()
	tracer, closer := tracer.Init(Name)
	otgo.SetGlobalTracer(tracer)
	return &Server{
		config:       config,
		CustomLogger: CustomLogger,
		Tracer:       tracer,
		Closer:       closer,
	}
}

func (server *Server) Start() {

	uri := fmt.Sprintf("mongodb://%s:%s", server.config.PostDBHost, server.config.PostDBPort)
	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), options)

	if err != nil {
		fmt.Println(err)
		server.CustomLogger.ErrorLogger.WithFields(logrus.Fields{"Post db host": server.config.PostDBHost, "Post db port": server.config.PostDBPort}).Error("MongoDB initialization for post service failed")
	} else {
		fmt.Println(uri)
		server.CustomLogger.SuccessLogger.Info("MongoDB initialization for post service succesfull")

	}
	store := repository.NewPostStore(client)
	postService := service.NewPostService(store)
	postController := controller.NewPostController(postService)

	server.CustomLogger.SuccessLogger.Info("Starting gRPC server for post service")
	server.startGrpcServer(postController)
}

func (server *Server) startGrpcServer(postController *controller.PostController) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		server.CustomLogger.ErrorLogger.Error("Failed to listen in post service: ", listener)
		log.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(server.Tracer)),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(server.Tracer)),
	}

	grpcServer := grpc.NewServer(opts...)
	postGW.RegisterPostServiceServer(grpcServer, postController)
	if err := grpcServer.Serve(listener); err != nil {
		server.CustomLogger.ErrorLogger.Error("Failed to serve gRPC in post service: ", listener)
		log.Fatalf("failed to serve: %s", err)
	}
}
