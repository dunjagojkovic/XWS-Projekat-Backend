package main

import (
	messageGW "common/proto/message_service"
	"context"
	"fmt"
	"log"
	"messageS/controller"
	"messageS/repository"
	"messageS/service"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type Server struct {
	config *Config
}

func NewServer(config *Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {

	uri := fmt.Sprintf("mongodb://%s:%s", server.config.MessageDBHost, server.config.MessageDBPort)
	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), options)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(uri)
	}
	store := repository.NewMessageStore(client)
	messageService := service.NewMessageService(store)

	messageController := controller.NewMessageController(messageService)

	server.startGrpcServer(messageController)
}

func (server *Server) startGrpcServer(jobController *controller.MessageController) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	messageGW.RegisterMessageServiceServer(grpcServer, jobController)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
