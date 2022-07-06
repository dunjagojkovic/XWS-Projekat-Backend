package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"userS/controller"
	"userS/repository"
	"userS/service"

	userGW "common/proto/user_service"
	saga "common/saga/messaging"
	"common/saga/messaging/nats"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type Server struct {
	config *Config
}

const (
	QueueGroup = "user_service"
)

func NewServer(config *Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {

	uri := fmt.Sprintf("mongodb://%s:%s", server.config.UserDBHost, server.config.UserDBPort)
	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), options)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(uri)
	}
	store := repository.NewUserStore(client)
	userService := service.NewUserService(store)

	commandSubscriber := server.initSubscriber(server.config.CreateMessageCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.CreateMessageReplySubject)
	server.initCreateMessageHandler(userService, replyPublisher, commandSubscriber)

	userController := controller.NewUserController(userService)

	server.startGrpcServer(userController)
}

func (server *Server) startGrpcServer(userController *controller.UserController) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	userGW.RegisterUserServiceServer(grpcServer, userController)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initCreateMessageHandler(service *service.UserService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := controller.NewCreateMessageCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}
